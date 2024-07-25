package algorithm

/*
 冒泡排序:
 先冒一个大的泡，再冒第二个小的泡
 每相邻的两个数进行比较，较大的往后
1, 67, 5, 4, 12, 45, 94, 415, 5

*/

func AIBubbleSort(data []int) []int {
	for i := cap(data); i > 0; i-- {
		var change bool
		for j := 0; j < i-1; j++ {
			a := data[j]
			b := data[j+1]
			if b >= a {
				continue
			}
			change = true
			a = a + b
			b = a - b
			a = a - b
			data[j] = a
			data[j+1] = b
		}
		if !change {
			break
		}
	}
	return data
}
