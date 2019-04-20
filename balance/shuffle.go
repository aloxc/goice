package balance

import (
	"math/rand"
)

//通过洗牌算法对服务器请求做负载均衡
type Shuffle struct {
}

var (
	shuffleMap = map[string][]int{}
)

func (this *Shuffle) GetIndex(name string, count int) int {
	arr, ok := shuffleMap[name]
	if !ok {
		arr = []int{}
		for i := 0; i < count; i++ {
			arr = append(arr, i)
		}
		shuffleMap[name] = arr
	}
	for i := count; i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		arr[lastIdx], arr[idx] = arr[idx], arr[lastIdx]
	}
	return arr[0]
}
