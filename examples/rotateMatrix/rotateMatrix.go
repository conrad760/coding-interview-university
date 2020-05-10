package main

import "fmt"

// main should rotate a matrix clockwise and print it out
func main() {
	//fill an matrix to test
	size := 5
	m := randMatrix(size)
	//rotate the array
	fmt.Println("filled:")
	printMatrix(m)
	clockwise(m)
	fmt.Println("rotated clockwise:")
	printMatrix(m)

}

func clockwise(m [][]int) {
	size := len(m[0])
	for layer := 0; layer < size/2; layer++ {
		start := layer
		last := size - layer - 1
		for i := start; i < last; i++ {
			offset := i - start
			// temp top
			top := m[start][i]
			// start is now bottom left
			m[start][i] = m[last-offset][start]
			// bottom left is now bottom right
			m[last-offset][start] = m[last][last-offset]
			// bottom right is now top right
			m[last][last-offset] = m[i][last]
			// top right is now top left
			m[i][last] = top
		}
	}
}
func printMatrix(m [][]int) {

	for i := 0; i < len(m[0]); i++ {
		fmt.Println(m[i])
	}

}
func randMatrix(n int) [][]int {
	// https://gobyexample.com/arrays
	m := make([][]int, n)
	// fmt.Println("emp:", m)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = j
		}
	}
	return m
}
