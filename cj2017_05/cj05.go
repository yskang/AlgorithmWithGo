package cj2017_05

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"sort"
	"math"
)

func CJ_2017_05() {
	args := os.Args[1:]
	inputFileName := args[0]
	outputFileName := strings.Replace(inputFileName, ".in", ".out", 1)

	results := make([]string, 0)
	inputs := readFile(inputFileName)

	T, err := strconv.Atoi(inputs.Pop())
	checkErr(err)

	for t := 0 ; t < T ; t++ {

		line := strings.Split(inputs.Pop(), " ")
		n, _ := strconv.Atoi(line[0]) // the size of island (max 10,000,000)
		k, _ := strconv.Atoi(line[1]) // the number of the watch center (max 2n)
		m, _ := strconv.Atoi(line[2]) // the number of the tsunami (max n)

		fmt.Println("size of this island is ", n)

		watches := make([]Watch, 0)
		watchMap := make(map[int]Watch)

		for i := 0 ; i < k ; i++ {
			position := strings.Split(inputs.Pop(), " ")
			x, _ := strconv.Atoi(position[0])
			y, _ := strconv.Atoi(position[1])
			watches = append(watches, Watch{i+1, x, y, false, 0})
			watchMap[i+1] = watches[len(watches)-1]
		}

		watches2 := make([]Watch, len(watches))
		copy(watches2, watches)

		outerWatches, outerMap := wrapIsland(watches)

		shelterMap := findShelters(outerWatches, outerMap, watchMap)

		//fmt.Println("shelter map", shelterMap)

		sum := 0
		prevAlert := make([]int, 2)
		for i := 0 ; i < m ; i++ {
			reports := strings.Split(inputs.Pop(), " ")
			a, _ := strconv.Atoi(reports[0])
			b, _ := strconv.Atoi(reports[1])

			//fmt.Println("Alert!! ", a, ",", b, "are sunken.. TT")

			if isValidAlert(a, b, watches2, prevAlert, outerMap) {
				//go to shelter
				shelterNumber := 0
				//fmt.Println("go to shelter")
				if a < b {
					//fmt.Println(shelterMap[fmt.Sprintf("%d,%d", a, b)])
					shelterNumber = shelterMap[fmt.Sprintf("%d,%d", a, b)]
				} else {
					//fmt.Println(shelterMap[fmt.Sprintf("%d,%d", b, a)])
					shelterNumber = shelterMap[fmt.Sprintf("%d,%d", b, a)]
				}
				sum += shelterNumber % 1000000007
			}
		}

		results = append(results, strconv.Itoa(sum % 1000000007))
	}

	result := strings.Join(results, "\n")

	writeResultFile(outputFileName, []byte(result))
}

func findShelters(outerWatches []int, outerMap map[int][]int, watchMap map[int]Watch) map[string]int{
	shelterMap := make(map[string]int)
	//fmt.Println(watchMap)
	base := outerWatches[0]
	for i := 1 ; i < len(outerWatches) ; i++ {
		//fmt.Println(base, outerWatches[i])
		if base < outerWatches[i] {
			shelterMap[fmt.Sprintf("%d,%d", base, outerWatches[i])] = getShelter(base, outerWatches[i], watchMap, outerWatches, i)
		} else {
			shelterMap[fmt.Sprintf("%d,%d", outerWatches[i], base)] = getShelter(base, outerWatches[i], watchMap, outerWatches, i)
		}
		base = outerWatches[i]
	}
	//fmt.Println(base, outerWatches[0])
	if base < outerWatches[0] {
		shelterMap[fmt.Sprintf("%d,%d", base, outerWatches[0])] = getShelter(base, outerWatches[0], watchMap, outerWatches, 0)
	} else {
		shelterMap[fmt.Sprintf("%d,%d", outerWatches[0], base)] = getShelter(base, outerWatches[0], watchMap, outerWatches, 0)
	}
	return shelterMap
}

func cBuffer(buffer []int, i int) int {
	if i < len(buffer) - 1 {
		return buffer[i]
	} else if i < 0 {
		return buffer[len(buffer)+i]
	}
	return buffer[i % (len(buffer))]
}

func getShelter(a int, b int, watches map[int]Watch, outWatches []int, key int) int {
	maxDistance := float64(0)
	maxShelter := make([]int, 0)

	if len(outWatches) > 500 {
		for i := len(outWatches) / 2 - 50 + key ; i < len(outWatches) / 2 + 50 + key ; i++ {
			distance := calcDistance(watches[a].x, watches[a].y, watches[b].x, watches[b].y, watches[cBuffer(outWatches, i)].x, watches[cBuffer(outWatches, i)].y)
			if maxDistance < distance {
				maxDistance = distance
				maxShelter = make([]int, 0)
				maxShelter = append(maxShelter, cBuffer(outWatches, i))
			} else if maxDistance == distance {
				maxShelter = append(maxShelter, cBuffer(outWatches, i))
			}
		}
	} else {
		for i := 0 ; i < len(outWatches) ; i++ {
			distance := calcDistance(watches[a].x, watches[a].y, watches[b].x, watches[b].y, watches[outWatches[i]].x, watches[outWatches[i]].y)
			if maxDistance < distance {
				maxDistance = distance
				maxShelter = make([]int, 0)
				maxShelter = append(maxShelter, outWatches[i])
			} else if maxDistance == distance {
				maxShelter = append(maxShelter, outWatches[i])
			}
		}
	}
	return chooseShelter(maxShelter, watches)
}

func to(key int, length int) int {
	to := key + length / 2 + 50
	if to > length {
		to = length
	}
	return length
}

func chooseShelter(shelterCandidates []int, watches map[int]Watch) int {
	if len(shelterCandidates) == 1 {
		return shelterCandidates[0]
	}

	minimumX := 9999999999
	minimumY := 9999999999
	minimumXShelter := 0
	minimumYShelter := 0
	for i := 0 ; i < len(shelterCandidates) ; i++ {
		if minimumX > watches[shelterCandidates[i]].x {
			minimumX = watches[shelterCandidates[i]].x
			minimumXShelter = watches[shelterCandidates[i]].number
		}
		if minimumY > watches[shelterCandidates[i]].y {
			minimumY = watches[shelterCandidates[i]].y
			minimumYShelter = watches[shelterCandidates[i]].number
		}
	}

	x := watches[shelterCandidates[0]].x
	for i := 0 ; i < len(shelterCandidates) ; i++ {
		if watches[shelterCandidates[1]].x != x {
			return minimumXShelter
		}
	}
	return minimumYShelter
}

func calcDistance(x1 int, y1 int, x2 int, y2 int, x int, y int) float64 {
	if x1 == x2 {
		return math.Abs(float64(x)-float64(x1))
	}
	a := float64(y1 - y2) / float64(x1 - x2)
	return math.Abs(float64(a) * float64(x - x1) - float64(y) + float64(y1)) / math.Sqrt(a * a + 1)
}

func isValidAlert(a int, b int, watches []Watch, prevAlert []int, outerMap map[int][]int) bool {
	if watches[a-1].prevSunken == watches[b-1].prevSunken && (prevAlert[0] != 0 && prevAlert[1] != 0) {
		//fmt.Println("false alarm! ")
		return false
	}

	//fmt.Println("outer map: ", outerMap)
	if val, isExist := outerMap[a]; isExist {
		if val[0] != b && val[1] != b {
			//fmt.Println("oops! false alarm!")
			return false
		}
	} else {
		//fmt.Println("oops! false alarm!")
		return false
	}
	//fmt.Println("True!! run!")

	//fmt.Println("++++", prevAlert[0])
	watches[prevAlert[0]].prevSunken = false
	//fmt.Println("++++", prevAlert[1])
	watches[prevAlert[1]].prevSunken = false

	watches[a-1].prevSunken = true
	watches[b-1].prevSunken = true


	prevAlert[0] = a - 1
	prevAlert[1] = b - 1

	return true
}

func wrapIsland(watches []Watch) ([]int, map[int][]int) {

	p0 := getLowAndLeftPoint(watches)

	getAngles(watches, p0)

	sort.Slice(watches, func(i, j int) bool {
		if watches[i].angle == watches[j].angle {
			if watches[i].number == p0.number {
				return true
			} else if watches[j].number == p0.number {
				return false
			}
		}
		return watches[i].angle < watches[j].angle
	})

	outers := NewStack()

	outers.Push(watches[0])
	outers.Push(watches[1])

	next := 2

	for next < len(watches) {
		for outers.Size() >= 2 {
			first := outers.Top()
			outers.Pop()
			second := outers.Top()

			if ccw(second, first, watches[next]) > 0 {
				outers.Push(first)
				break
			}
		}
		outers.Push(watches[next])
		next += 1
	}

	outerWaches := make([]int, 0)
	outersMap := make(map[int][]int)
	prevNumber := p0.number
	for !outers.IsEmpty() {
		outer := outers.Pop()
		outerWaches = append(outerWaches, outer.number)
		outersMap[outer.number] = append(outersMap[outer.number], prevNumber)
		outersMap[prevNumber] = append(outersMap[prevNumber], outer.number)
		prevNumber = outer.number
	}

	return outerWaches, outersMap
}

func ccw(p1 Watch, p2 Watch, p3 Watch) int{
	return (p2.x - p1.x) * (p3.y - p1.y) - (p2.y - p1.y) * (p3.x - p1.x)
}

func getAngles(watches []Watch, p0 Watch) {
	for i := 0 ; i < len(watches) ; i++ {
		cx := watches[i].x - p0.x
		cy := watches[i].y - p0.y

		angle := math.Atan2(float64(cy), float64(cx))
		angle = angle * 180 / math.Pi
		watches[i].angle = math.Abs(angle)
	}
}

func getLowAndLeftPoint(watches []Watch) Watch {
	lowestY := 999999999
	lowests := make([]Watch, 0)

	for _, watch := range watches {
		if watch.y < lowestY {
			lowestY = watch.y
			lowests = make([]Watch, 0)
			lowests = append(lowests, watch)
		} else if watch.y == lowestY {
			lowests = append(lowests, watch)
		}
	}

	sort.Slice(lowests, func(i, j int) bool {return lowests[i].x < lowests[j].x})
	return lowests[0]
}


type Watch struct {
	number     int
	x          int
	y          int
	prevSunken bool
	angle      float64
}

func readFile(path string) Queue {
	t1 := time.Now()
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(strings.Replace(string(dat), "\r", "", -1), "\n")
	size := len(lines)
	fmt.Println("read file time : ", time.Now().Sub(t1))
	return Queue{lines, size, 0, size-1, size-1}
}

func writeResultFile(path string, data []byte) {
	ioutil.WriteFile(path, data, os.ModePerm)
}

type Queue struct {
	nodes []string
	size  int
	head  int
	tail  int
	count int
}

func (q *Queue) Pop() string {
	if q.count == 0 {
		return "Empty"
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type myStack struct {
	stack []Watch
}

func NewStack() *myStack {
	return &myStack{make([]Watch, 0)}
}

func (s *myStack) Push(v Watch) {
	s.stack = append(s.stack, v)
}

func (s *myStack) Pop() Watch {
	topVal := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return topVal
}

func (s *myStack) Top() Watch {
	topVal := s.stack[len(s.stack)-1]
	return topVal
}

func (s *myStack) Size() int {
	return len(s.stack)
}

func (s *myStack) IsEmpty() bool {
	if len(s.stack) > 0 {
		return  false
	}
	return true
}

func (s *myStack) Clear() {
	for {
		if s.IsEmpty() == true {
			return
		}
		s.Pop()
	}
}