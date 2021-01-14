package sort

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
