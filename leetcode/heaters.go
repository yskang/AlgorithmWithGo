package leetcode

import (
	"fmt"
	"math"
	"sort"
)

func FindRadius() {
	fmt.Println(findRadius([]int{282475249,622650073,984943658,144108930,470211272,101027544,457850878,458777923}, []int{823564440,115438165,784484492,74243042,114807987,137522503,441282327,16531729,823378840,143542612}))
	//fmt.Println(findRadius([]int{1,2,3,4,5,6,7,8,9,10}, []int{2,5,8,9}))
}

func findRadius(houses []int, heaters []int) int {
	heaterMap := make(map[int]int)
	sort.Ints(heaters)

	for _, house := range(houses) {
		if house < heaters[0] {
			fmt.Println("house is left of heaters")
			putHouseToHeater(house, heaters[0], heaterMap)
		} else if house > heaters[len(heaters) - 1] {
			fmt.Println("house is right of heaters")
			putHouseToHeater(house, heaters[len(heaters) - 1], heaterMap)
		} else {
			for i := 0 ; i < len(heaters) - 1 ; i++ {
				if heaters[i] < house && house < heaters[i+1] {
					distanceA := house - heaters[i]
					distanceB := heaters[i+1] - house
					if distanceA < distanceB {
						putHouseToHeater(house, heaters[i], heaterMap)
					} else {
						putHouseToHeater(house, heaters[i+1], heaterMap)
					}
				}
			}
		}
	}

	fmt.Println(heaterMap)

	maxRange := 0

	for _, distance := range(heaterMap) {
		if maxRange < distance {
			maxRange = distance
		}
	}

	return maxRange
}

func putHouseToHeater(house int, heater int, heaterMap map[int]int) {
	distance := int(math.Abs(float64(heater - house)))
	if maximumRange, isExist := heaterMap[heater]; isExist {
		if maximumRange < distance {
			heaterMap[heater] = distance
		}
	} else {
		heaterMap[heater] = distance
	}
}

