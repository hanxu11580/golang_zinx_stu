package utils

import (
	"math/rand"
)

// 包含min/max
// 注意：以时间戳为随机种子 有秒级别和纳秒级别、还有其他的
// time.Now().UnixNano() 纳秒
// time.Now().Unix() 秒
func RangeRandom(min, max int, seed int64) int {
	rand.Seed(seed)
	randNum := rand.Intn(max-min+1) + min
	return randNum
}
