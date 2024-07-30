package algorithm

/**
插入排序
*/


func InsertSort(source []int) []int {
	for i := 0; i < len(source); i++ {
		var temp = source[i]
		for j := i; j > 0; j-- {
			// 如果0-i已经是有序的了
			// 需要对这新来的i进行排序
			if source[j] > temp {
				break
			}
			if source[j-1] <

		}
	}
	return source
}
