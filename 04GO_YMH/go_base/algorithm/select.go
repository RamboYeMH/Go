package algorithm

/*
 选择排序
 选择最小的一个记录其下标，在第几次循环的时候为其替换位置
*/

func SelectSort(source []int) []int {
	for i, v := range source {
		var temp = i
		for j := i + 1; j < len(source); j++ {
			if source[j] < v {
				i = j
				v = source[j]
			}
		}
		if temp == i {
			continue
		}
		v = v + source[temp]
		source[temp] = v - source[temp]
		source[i] = v - source[temp]
	}
	return source
}
