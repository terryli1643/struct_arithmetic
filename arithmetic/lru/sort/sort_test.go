import (
	"net/http"
	"testing"

	"github.com/terryli1643/struct_arithmetic/arithmetic/lru/sort"
)

func TestBubbleSort(t *testing.T) {
	a := []int{2, 5, 6, 7, 32, 6, 23, 4, 7, 8, 9, 4, 3, 5, 3}
	sort.BubbleSort(a)
	for _, v := range a {
		t.Log(v)
	}
}

func TestInsertionSort(t *testing.T) {
	a := []int{2, 5, 6, 7, 32, 6, 23, 4, 7, 8, 9, 4, 3, 5, 3}
	sort.InsertionSort(a)
	for _, v := range a {
		t.Log(v)
	}
}

func TestSelectionSort(t *testing.T) {
	a := []int{2, 5, 6, 7, 32, 6, 23, 4, 7, 8, 9, 4, 3, 5, 3}
	sort.SelectionSort(a)
	for _, v := range a {
		t.Log(v)
	}
}
