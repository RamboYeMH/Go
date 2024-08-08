package algorithm

/**
希尔排序：希尔排序是插入排序的一种更高效的改进
该算法是冲破O(n2)的第一批算法之一
https://www.cnblogs.com/chengxiao/p/6104371.html
*/

func ShellSort(source []int) {
	// gap是组的意思
	for gap := cap(source) / 2; gap > 0; gap = gap / 2 {
		// 从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < cap(source); i++ {
			j := i
			for source[j] < source[j-gap] && j > gap {
				swap(source, j, j-gap)
				// todo ? 20 * 15
				j -= gap
			}
		}
	}
}

func swap(source []int, a int, b int) {
	source[a] = source[a] + source[b]
	source[b] = source[a] - source[b]
	source[a] = source[a] - source[b]
}
