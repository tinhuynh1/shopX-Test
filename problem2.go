package main

import "fmt"

func main() {

	var image = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateImage(image)
	fmt.Println(image)

}

func rotateImage(m [][]int) {
	n := len(m)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := m[i][j]
			m[i][j] = m[j][i]
			m[j][i] = temp

		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			temp := m[i][j]
			m[i][j] = m[i][n-1-j]
			m[i][n-1-j] = temp

		}
	}
}
