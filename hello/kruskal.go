/*
//D: 가중치 배열 (1차원)
//r, b), r은 이미 선택된 정점, b는 선택이 안된 정점
//S: 선택이 안된 정점
Prim 알고리즘2
D[v0] = 0 //임의로 한 정점 v0를 시작점으로 선택
for(v0를 제외한 정점 vi) D[vi] = w[v0, vi]
S= 모든 정점
T = T ∪ {v0}
S = S - {v0}
while(T의 정점 수 ＜ n) {
 vmin = MIN(D[vj]), vi∈ T , vj ∈ S, (vi,vj) ≠ ∞
 T = T ∪ {vmin}, T = T∪ {(vi, vmin)}
 S = S - {vmin}
 for(S에 속하는 모든 정점, vj)
  if(W[vmin, vj] ＜ D[vj]) D[vj] = W[vmin, vj]
}

Kruskal 알고리즘
Edgelist = 가중치의 오름차순으로 간선들을 정렬
T = 공집합
ecount = 0, k=0
while(ecount < n-1) {
 (v,w) = Edgelist[k]
 if(vㄷT and w ㄷT) { //사이클인지 확인
 T = T∪ {(v,w)}
 T = T∪{v}
 T = T∪{w}
 ecount += 1
 k+=1
}


*/

package main

import "fmt"

func quickSort(data []int, left int, right int) {
	var i, j, key, temp int
	if left < right {
		i = left
		j = right + 1
		key = data[left]
	}
	for i < j {
		for data[i] < key {
			i++
		}
		for data[j] > key {
			j--
		}
		if i < j {
			temp = data[i]
			data[i] = data[j]
			data[j] = temp
		}
	}
	temp = data[left]
	data[left] = data[j]
	data[j] = temp
	quickSort(data, left, j-1)
	quickSort(data, j+1, right)

}

func kruskal() {

}

func main() {
	fmt.Println("Hello")
}
