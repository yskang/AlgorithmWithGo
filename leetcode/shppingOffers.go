package leetcode

func ShoppingOffers(price []int, special [][]int, needs []int) int {
	return shoppingOffers(price, special, needs)
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	return shopping(price, special, needs)
}

func shopping(price []int, special [][]int, needs []int) int {
	cost := singleCost(price, needs)

	for _, offer := range special {
		offerCost := offer[len(offer)-1]
		newNeeds := append([]int{}, needs...)
		isValidOffer := true
		for i := 0 ; i < len(offer) - 1 ; i++ {
			temp := needs[i] - offer[i]
			if temp < 0 {
				isValidOffer = false
				break
			}
			newNeeds[i] = temp
		}
		if isValidOffer {
			cost = shoppingMin(cost, offerCost + shopping(price, special, newNeeds))
		}
	}
	return cost
}

func singleCost(prices []int, needs []int) int {
	totalCost := 0
	for i := range prices {
		totalCost += prices[i] * needs[i]
	}
	return totalCost
}

func shoppingMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}