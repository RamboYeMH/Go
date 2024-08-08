package algorithm

/**
插入排序又称直接插入排序算法
因为第一个数是有序的所以，除开第一个数以外的数都是无序的，当它到来的时间
需要寻找他能到哪一个下标,相比于插入排序其最快的排序为O(n)
*/

func InsertSort(source []int) []int {
	for i, v := range source {
		var temp int = v
		// 从已经排序的序列最右边开始比较，找到比其小的数
		var j = i
		for j > 0 && temp < source[j-1] {
			source[j] = source[j-1]
			j--
		}
		// 存在比其小的数，插入
		if j != i {
			source[j] = temp
		}

	}
	return source
}
