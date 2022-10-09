/*
그래프 색칠하기
m-색칠하기 문제 : 무방향 그래프에서 최대 m(>1)개의 색들을 사용하여 인접한 정점들을 서로 다른 색으로 칠하라
*/

package main

import "fmt"

type Input struct {
	n      int //정점 수
	m      int // 색의 수
	vcolor []int
}

func valid(G [][]int, i int, vcolor []int) bool {

	for j := 1; j < i; j++ {
		if G[i-1][j-1] == 1 && vcolor[i] == vcolor[j] {
			return false //인접한 정점끼리 같은 색이 있으면 false
		}
	}
	return true
}

//재귀호출 사용 (or stack)
func m_coloring(G [][]int, i int, vcolor []int, n int, m int, count *int) {

	if valid(G, i, vcolor) { //i번 정점 칠할 수 있는지 확인.(가지치기)
		if i == n { //모든 정점이 칠해졌는지 확인. (종료조건)
			for i, v := range vcolor {
				if v != 0 {

					fmt.Printf("%d번째 색칠 노드: %d, 색: %d\n", *count, i, v)
					*count++
					*count = *count % 6
					if *count == 0 {
						*count++
					}

				}
			}
		} else {
			for c := 1; c <= m; c++ {
				vcolor[i+1] = c
				m_coloring(G, i+1, vcolor, n, m, count)
			}
		}
	}
}

func main() {
	n := 5 // 정점들의 수
	m := 3 //색들의 수
	//그래프 G = (V, E)는 인접 행렬 G로 표현
	G := [][]int{{0, 1, 0, 1, 1}, //G[i-1][j-1] = 1,(i,j) ∈ E 간선이 있는 경우
		{1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 1},
		{1, 0, 0, 1, 0}}

	count := 1

	vcolor := []int{0, 0, 0, 0, 0, 0}
	//vcolor[i], i:1~n 정점 i에 칠해진 색

	m_coloring(G, 0, vcolor, n, m, &count)
	//구현 문제 때문에 n, m을 넣었고(전역변수로 구현하지 않기 위해)
	// count는 모든 경우 출력하기 때문에 구분을 위해 넣었음 (없어도 구현에 문제x)

}
