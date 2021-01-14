package sort

func InsertionSort(a []int) {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i - 1; j > 0; j-- {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			} else {
				break
			}
		}
	}
}
