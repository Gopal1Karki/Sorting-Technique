package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort!")

	fmt.Printf("\n")

	array := []int{10, 56, 32, 96, 8}

	BubbleSort(array)
	fmt.Printf("\n")
	SelectionSort(array)
	fmt.Printf("\n")
	InsertionSort(array)
	fmt.Printf("\n")
}

func BubbleSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	fmt.Println("The Sorted array using bubble sort is ", array)
}

func SelectionSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		p := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[p] {
				p = j
			}
		}
	}
	fmt.Println("The Sorted array using Selection sort is ", array)
}

func InsertionSort(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		temp := array[i]
		j := i - 1
		for j >= 0 && array[j] > temp {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = temp
	}
	fmt.Println("The Sorted array using Insertion sort is ", array)
}
