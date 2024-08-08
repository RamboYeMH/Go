package algorithm

/**
归并排序：是利用归并的思想实现排序方法，该算法采用经典的分治策略
分：将问题分成一些小的问题然后递归求解
治：将分的阶段得到的各个答案“修补”在一起
即分而治之
*/
// 步骤
// 1.计算数组中点mid.递归划分左子数组（区间[left,mid]）和右子数组(区间[mid+1,right])
// 2.递归执行不走1.,直至子数组长度为1时终止。

func mergeSort(nums []int, left int, mid int, right int) {
	// 左子树区间为[left, mid]. 右子树区间为[mid+1,right]
	// 创建一个临时切片tmp,用户存放合并后的结构
	temp := make([]int, left-right+1)
	// 初始化左子数组和右子数组的起始索引
	i := left
	j := mid + 1
	k := 0

}
