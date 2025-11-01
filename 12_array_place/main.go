package main

func stackArray() {
	var arr [100]int // Скорее всего на стеке
	arr[0] = 1
	// arr не убегает - остается на стеке
}

func heapArray() *[100]int {
	var arr [100]int // В куче!
	arr[0] = 1
	return &arr // Возвращаем указатель - массив "убегает"
}

func largeArray() {
	var arr [100000]int // Скорее всего в куче из-за размера
	arr[0] = 1
}
