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
	direction []string
	count int
	x int
	y int
	infomation info
	origin position
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
		if b.y-b.count >= 0 && b.y-b.count < b.infomation.height && !b.infomation.isOnWater(b.x, b.y-b.count){
			balls = append(balls, ball{append(b.direction, "up"),b.count-1, b.x, b.y-b.count, b.infomation, b.origin})
		}
		if b.y+b.count < b.infomation.height && !b.infomation.isOnWater(b.x, b.y+b.count){
			balls = append(balls, ball{append(b.direction, "down"),b.count-1, b.x, b.y+b.count, b.infomation, b.origin})
		}
		if b.x-b.count >= 0 && b.x-b.count < b.infomation.width && !b.infomation.isOnWater(b.x-b.count, b.y){
			balls = append(balls, ball{append(b.direction, "left"),b.count-1, b.x-b.count, b.y, b.infomation, b.origin})
		}
		if b.x+b.count < b.infomation.width && !b.infomation.isOnWater(b.x+b.count, b.y) {
			balls = append(balls, ball{append(b.direction, "right"),b.count-1, b.x+b.count, b.y, b.infomation, b.origin})
		}
	} else {
		if b.direction[len(b.direction)-1] != "down" && b.y-b.count >= 0  && b.y-b.count < b.infomation.height && !b.infomation.isOnWater(b.x, b.y-b.count){
			balls = append(balls, ball{append(b.direction, "up"),b.count-1, b.x, b.y-b.count, b.infomation, b.origin})
		}
		if b.direction[len(b.direction)-1] != "up" && b.y+b.count < b.infomation.height && !b.infomation.isOnWater(b.x, b.y+b.count){
			balls = append(balls, ball{append(b.direction, "down"),b.count-1, b.x, b.y+b.count, b.infomation, b.origin})
		}
		if b.direction[len(b.direction)-1] != "right" && b.x-b.count >= 0 && b.x-b.count < b.infomation.width && !b.infomation.isOnWater(b.x-b.count, b.y){
			balls = append(balls, ball{append(b.direction, "left"),b.count-1, b.x-b.count, b.y, b.infomation, b.origin})
		}
		if b.direction[len(b.direction)-1] != "left" && b.x+b.count < b.infomation.width && !b.infomation.isOnWater(b.x+b.count, b.y){
			balls = append(balls, ball{append(b.direction, "right"),b.count-1, b.x+b.count, b.y, b.infomation, b.origin})
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
			if byte(c) >= '0' && byte(c) <= '9'  {
				balls = append(balls, ball{[]string{"first"},int(byte(c) - '0'), x, i, gameInfo, position{x, i}})
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

	routeMap := make(map[position][][]string)
	ballInfos := make([]ballInfo, 0)
	for _, ball := range balls {
		fmt.Fprintln(os.Stderr, "for ball:", ball)
		ballInfos = append(ballInfos, ballInfo{ball.origin, ball.count})
		hitBall(field, ball, routeMap)
	}

	fmt.Fprintln(os.Stderr, "route map:", routeMap)

	output := make([]string, height)

	drawRoute(output, routeMap, ballInfos, 0)

	fmt.Fprintln(os.Stderr, output)
}

type ballInfo struct {
	pos position
	num int
}

func drawRoute(output []string, routeMap map[position][][]string, ballInfos []ballInfo, depth int) {
	if depth >= len(ballInfos) - 1 {
		return
	}

	fmt.Fprintln(os.Stderr, "depth", depth)
	for _, route := range routeMap[ballInfos[depth+1].pos] {
		if success, output := checkRoute(route, output, ballInfos[depth+1]); success {
			drawRoute(output, routeMap, ballInfos, depth+1)

		}
	}
}

func checkRoute(route []string, out []string, info ballInfo) (bool, []string) {
	startX := info.pos.x
	startY := info.pos.y
	distance := info.num

	for i, direction := range route[1:] {
		distance = distance - i
		if direction == "up" {
			for y := 1 ; y <= distance ; y++ {
				if byte(out[startY - y][startX]) == '.' {
					//byte(out[startY - y][startX]) = '^'
					out[startY - y] = out[startY - y][:startX] + "^" +out[startY - y][startX+1:]
				} else {
					return false, nil
				}
			}
		} else if direction == "down" {
			for y := 1 ; y <= distance ; y++ {
				if byte(out[startY + y][startX]) == '.' {
					//byte(out[startY + y][startX]) = 'v'
					out[startY - y] = out[startY + y][:startX] + "v" +out[startY + y][startX+1:]
				} else {
					return false, nil
				}
			}
		} else if direction == "left" {
			for x := 1 ; x <= distance ; x++ {
				if byte(out[startY][startX-x]) == '.' {
					//byte(out[startY][startX-x]) = '<'
					out[startY] = out[startY][:startX-x] + "<" +out[startY][startX-x+1:]
				} else {
					return false, nil
				}
			}
		} else if direction == "right" {
			for x := 1 ; x <= distance ; x++ {
				if byte(out[startY][startX+x]) == '.' {
					//byte(out[startY][startX+x]) = '>'
					out[startY] = out[startY][:startX+x] + ">" +out[startY][startX+x+1:]
				} else {
					return false, nil
				}
			}
		}
	}

	return true, out
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