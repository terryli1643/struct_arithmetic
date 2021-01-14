package sort

/*
冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。
如果不满足就让它俩互换。一次冒泡会让至少一个元素移动到它应该在的位置，
重复 n 次，就完成了 n 个数据的排序工作。

最好时间复杂度为:O(n)
最坏时间复杂度为：O(n²)
平均时间复杂度为：O(n²)
*/
func BubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		swich := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swich = true
			}
		}
		if !swich {
			return
		}
	}
}
