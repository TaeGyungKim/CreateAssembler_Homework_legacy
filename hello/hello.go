/*C는 한정된 용량

n개의 물건
무게 Wi 1 <= i <= n
가치 vi 1 <= i <= n
배낭에 담을 수 있는 물건의 최대 가치
	배낭에 담은 물건의 무게의 합은 C를 초과하지 말아야함
*/

package main

import "fmt"

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Knapsack(C int, w []int, v []int, n int) [5][8]int {
	var F [5][8]int //물건과 개수

	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			if w[i] > j {
				F[i][j] = F[i-1][j]
			} else {
				F[i][j] = Max(v[i]+F[i-1][j-w[i]], F[i-1][j])
			}
		}
	}
	return F
}

func main() {
	W := []int{0, 3, 1, 2, 4}     // 물건의 무게
	V := []int{0, 25, 15, 20, 30} //물건의 가치
	C := 7                        //배낭의 용량
	n := 4                        // 물건 개수

	var F [5][8]int
	F = Knapsack(C, W, V, n) //배낭

	for i, v := range F {
		fmt.Printf("물건 i : %d, 값: %d\n", i, v)
	}
}
