package main

import (
	"fmt"
	"os"
)
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type ball struct {
	direction   []string
	count       int
	x           int
	y           int
	information *info
	origin      position
}

type info struct {
	golfMap []string
	width int
	height int
}

func (i *info) setMap(m []string) {
	copy(i.golfMap, m)
}

func (i *info) isOnWater(x int, y int) bool {
	if byte(i.golfMap[y][x]) == 'X' {
		return true
	}
	return false
}

func (b *ball) getNext() []ball {
	balls := make([]ball, 0)
	if b.count == 0 {
		return nil
	}

	if b.direction[len(b.direction)-1] == "first" {
		if b.y-b.count >= 0 && b.y-b.count < b.information.height && !b.information.isOnWater(b.x, b.y-b.count){
			balls = append(balls, ball{append(b.direction, "up"),b.count-1, b.x, b.y-b.count, b.information, b.origin})
		}
		if b.y+b.count < b.information.height && !b.information.isOnWater(b.x, b.y+b.count){
			balls = append(balls, ball{append(b.direction, "down"),b.count-1, b.x, b.y+b.count, b.information, b.origin})
		}
		if b.x-b.count >= 0 && b.x-b.count < b.information.width && !b.information.isOnWater(b.x-b.count, b.y){
			balls = append(balls, ball{append(b.direction, "left"),b.count-1, b.x-b.count, b.y, b.information, b.origin})
		}
		if b.x+b.count < b.information.width && !b.information.isOnWater(b.x+b.count, b.y) {
			balls = append(balls, ball{append(b.direction, "right"),b.count-1, b.x+b.count, b.y, b.information, b.origin})
		}
	} else {
		if b.direction[len(b.direction)-1] != "up" && b.direction[len(b.direction)-1] != "down" && b.y-b.count >= 0  && b.y-b.count < b.information.height && !b.information.isOnWater(b.x, b.y-b.count){
			balls = append(balls, ball{append(b.direction, "up"),b.count-1, b.x, b.y-b.count, b.information, b.origin})
		}
		if b.direction[len(b.direction)-1] != "down" && b.direction[len(b.direction)-1] != "up" && b.y+b.count < b.information.height && !b.information.isOnWater(b.x, b.y+b.count){
			balls = append(balls, ball{append(b.direction, "down"),b.count-1, b.x, b.y+b.count, b.information, b.origin})
		}
		if b.direction[len(b.direction)-1] != "left" && b.direction[len(b.direction)-1] != "right" && b.x-b.count >= 0 && b.x-b.count < b.information.width && !b.information.isOnWater(b.x-b.count, b.y){
			balls = append(balls, ball{append(b.direction, "left"),b.count-1, b.x-b.count, b.y, b.information, b.origin})
		}
		if b.direction[len(b.direction)-1] != "right" && b.direction[len(b.direction)-1] != "left" && b.x+b.count < b.information.width && !b.information.isOnWater(b.x+b.count, b.y){
			balls = append(balls, ball{append(b.direction, "right"),b.count-1, b.x+b.count, b.y, b.information, b.origin})
		}
	}

	return balls
}

type hole struct {
	x int
	y int
}

func main() {
	var width, height int
	fmt.Scan(&width, &height)

	fmt.Fprintln(os.Stderr, "map width, height : ", width, height)

	field := make([]string, height)
	balls := make([]ball, 0)
	holes := make([]hole, 0)

	gameInfo := info{make([]string, height), width, height}

	for i := 0; i < height; i++ {
		var row string
		fmt.Scan(&row)

		field[i] = row

		for x, c := range row {
			if byte(c) >= '0' && byte(c) <= '9' {
				balls = append(balls, ball{[]string{"first"}, int(byte(c) - '0'), x, i, &gameInfo, position{x, i}})
			}
			if byte(c) == 'H' {
				holes = append(holes, hole{x, i})
			}
		}
	}

	gameInfo.setMap(field)

	fmt.Fprintln(os.Stderr, "Field:")
	for _, row := range field {
		fmt.Fprintln(os.Stderr, row)
	}

	routeVectorsMap := make(map[position][][]string)
	ballInfos := make([]ballInfo, 0)
	for _, ball := range balls {
		fmt.Fprintln(os.Stderr, "for ball:", ball)
		ballInfos = append(ballInfos, ballInfo{ball.origin, ball.count})
		hitBall(field, ball, routeVectorsMap)
	}

	fmt.Fprintln(os.Stderr, "route vector map:", routeVectorsMap)

	routeMap := make(map[position][]route)
	for p := range routeVectorsMap {
		for _, r := range routeVectorsMap[p] {
			routeMap[p] = append(routeMap[p], vector2route(r, p, width, height, int(byte(field[p.y][p.x])-'0')))
		}
	}

	// print route map
	for _, ball := range balls {
		fmt.Fprintln(os.Stderr, "route of ball ", ball.origin)
		for _, r := range routeMap[ball.origin] {
			for _, row := range r.cells {
				line := ""
				for _, c := range row {
					if c == "" {
						c = " "
					}
					line += c
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
		fmt.Fprintln(os.Stderr, "")
	}

	mergeRoutes := make([][]route, height+1)
	mergeRoutes[0] = make([]route, 1)
	mergeRoutes[0][0] = makeRoute(width, height)
	for i := 1 ; i <= height ; i++ {
		mergeRoutes[i] = make([]route, 0)
	}

	drawMergeRoute(routeMap, balls, 0, &mergeRoutes)

	fmt.Fprintln(os.Stderr, "final route:", mergeRoutes)
	fmt.Fprintln(os.Stderr, "route size:", len(mergeRoutes[len(balls)]))

	for _, row := range mergeRoutes[len(balls)][0].print() {
		fmt.Println(row)
	}


}

func drawMergeRoute(routeMap map[position][]route, balls []ball, i int, mergeRoutes *[][]route) {
	if len(balls) <= i {
		return
	}

	for _, r := range routeMap[balls[i].origin] {
		for _, mr := range (*mergeRoutes)[i] {
			if !isCross(mr, r) {
				(*mergeRoutes)[i+1] = append((*mergeRoutes)[i+1], merge(mr, r))
			}
		}
	}

	drawMergeRoute(routeMap, balls, i + 1, mergeRoutes)
}

func merge(r route, rm route) route {
	merged := makeRoute(len(r.cells[0]), len(r.cells))

	for y := range r.cells {
		for x := range r.cells[y] {
			merged.cells[y][x] = r.cells[y][x] + rm.cells[y][x]
		}
	}

	//fmt.Fprintln(os.Stderr, "merged route:", merged)
	return merged
}


func isCross(r route, rm route) bool {
	for y := range r.cells {
		for x := range r.cells[y] {
			if r.cells[y][x] != "" && rm.cells[y][x] != "" {
				return true
			}
		}
	}

	return false
}


func vector2route(vector []string, p position, width int, height int, num int) route {
	r := makeRoute(width, height)
	pos := p
	l := num
	for i, v := range vector[1:] {
		pos = r.setDirection(pos, v, l-i)
	}
	r.cells[pos.y][pos.x] = "H"
	return r
}

type route struct {
	cells [][]string
}

func makeRoute(width int, height int) route {
	r := make([][]string, height)
	for i := range r {
		r[i] = make([]string, width)
	}
	return route{r}
}

func (c *route) setDirection(p position, direction string, length int) position {
	x, y := p.x, p.y
	for i := 0 ; i < length ; i++ {
		switch direction {
		case "up":
			c.cells[p.y-i][p.x] = "^"
			y -= 1
			break
		case "down":
			c.cells[p.y+i][p.x] = "v"
			y += 1
			break
		case "left":
			c.cells[p.y][p.x-i] = "<"
			x -= 1
			break
		case "right":
			c.cells[p.y][p.x+i] = ">"
			x += 1
			break
		}
	}
	return position{x, y}
}

func (r *route) print() []string {
	h := len(r.cells)

	result := make([]string, h)
	for y := range r.cells {
		row := ""
		for x := range r.cells[0] {
			if r.cells[y][x] == "" || r.cells[y][x] == "H" {
				row += "."
			} else {
				row += r.cells[y][x]
			}
		}
		result[y] = row
	}
	return result
}

type ballInfo struct {
	pos position
	num int
}

func hitBall(field []string, b ball, routeMap map[position][][]string) {
	if byte(field[b.y][b.x]) == 'H' {
		fmt.Fprintln(os.Stderr, "Hole!! append", b.direction)
		routeMap[b.origin] = append(routeMap[b.origin], b.direction)

		return
	}

	if b.getNext() == nil {
		return
	}

	for _, nextBall := range b.getNext() {
		hitBall(field, nextBall, routeMap)
	}
}


type position struct {
	x int
	y int
}