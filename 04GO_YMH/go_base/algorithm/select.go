package algorithm

/*
 选择排序
 选择最小的一个记录其下标，在第几次循环的时候为其替换位置
 选择排序的缺点就是无论如何都要走O(n^2) 不能像冒泡排序一样可以在一次O(n)就能检测出游戏的数据
*/

func SelectSort(source []int) []int {
	for i, v := range source {
		var selectIndex = i
		for j := i + 1; j < len(source); j++ {
			if source[j] < v {
				i = j
				v = source[j]
			}
		}
		if selectIndex == i {
			continue
		}
		v = v + source[selectIndex]
		source[selectIndex] = v - source[selectIndex]
		source[i] = v - source[selectIndex]
	}
	return source
}
