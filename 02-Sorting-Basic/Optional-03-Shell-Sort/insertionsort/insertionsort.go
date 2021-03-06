package insertionsort

import (
	"github.com/donng/Play-with-Data-Algorithms/02-Sorting-Basic/Optional-03-Shell-Sort/helper"
)

func Sort(arr []int, length int) {
	for i := 1; i < length; i++ {
		// 寻找元素 arr[i] 合适的插入位置
		temp := arr[i]
		// j 保存元素 temp 应该插入的位置
		var j int
		for j = i; j > 0 && temp < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 测试InsertionSort
func main() {
	N := 20000
	arr := helper.GenerateRandomArray(N, 0, 100000)
	helper.TestSort(Sort, arr)
}
