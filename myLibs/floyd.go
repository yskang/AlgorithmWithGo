package myLibs

import "fmt"


//1 procedure floyd-warshall(두 정점을 잇는 모서리의 가중치 테이블 W)
//2     D : 두 정점을 잇는 경로의 최소 비용 테이블. 모든 성분을 ∞로 초기화
//3     M : 두 정점을 지나가는 최소 비용 경로가 거쳐야 할 마지막 경유지 테이블. 모든 성분을 null로 초기화
//4     for i from 1 to |V|
//5         for j from 1 to |V|
//6             D[i][j] := W[i][j]
//7     for v from 1 to |V|
//8         D[v][v] := 0
//9     for k from 1 to |V|
//10        for i from 1 to |V|
//11            for j from 1 to |V|
//12                if D[i][j] > D[i][k] + D[k][j]
//13                    D[i][j] := D[i][k] + D[k][j]
//14                    M[i][j] := k
//15    return D

func Floyd(n int)  {
	costTable := make([][]int, 0)
	viaTable := make([][]int, 0)

	for i := 0 ; i < n ; i++ {
		costTable = append(costTable, make([]int, n))
		viaTable = append(viaTable, make([]int, n))
	}

	for k := 0 ; k < n ; k++ {
		for i := 0 ; i < n ; i++ {
			for j := 0 ; j < n ; j++ {
				if costTable[i][j] > costTable[i][k] + costTable[k][j] {
					costTable[i][j] = costTable[i][k] + costTable[k][j]
					viaTable[i][j] = k
				}
			}
		}
	}

	fmt.Println(costTable)
	fmt.Println(viaTable)
}