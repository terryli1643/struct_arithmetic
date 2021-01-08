package test_lru

import (
	"fmt"
	"testing"

	"github.com/terryli1643/struct_arithmetic/arithmetic/lru"
)

func TestLRU(t *testing.T) {
	myLRU := lru.NewMyList(10)
	for i := 0; i < 10; i++ {
		myLRU.Put(i)
	}

	for i := 0; i < 5; i++ {
		myLRU.Get(i)
	}

	for i := 10; i < 15; i++ {
		myLRU.Put(i)
	}

	for i := 0; i < myLRU.Len(); i++ {
		fmt.Print(myLRU.Value)
		myLRU = myLRU.Next()
	}
}
