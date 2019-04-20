package balance

import (
	"fmt"
	"testing"
)

func TestShuffle_GetIndex(t *testing.T) {
	var cnt2 = map[int]int{}
	s := Shuffle{}
	var idx = 0
	for i := 0; i < 10000000; i++ {
		idx = s.GetIndex("a", 7)
		cnt2[idx]++
	}
	fmt.Println(cnt2)
}
