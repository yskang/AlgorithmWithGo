package leetcode

func FriendCircles(M [][]int) int {
	return findCircleNum(M)
}

func findCircleNum(M [][]int) int {
	members := make([]int, len(M))
	count := 0

	for member, hasCheck := range members {
		if hasCheck != 1 {
			count += 1
			members[member] = 1
			dfs(member, M, members)
		}
	}

	return count
}

func dfs(member int, M [][]int, members []int) {
	for i, isFriend := range M[member] {
		if isFriend == 1 && members[i] != 1 {
			members[i] = 1
			dfs(i, M, members)
		}
	}
}
