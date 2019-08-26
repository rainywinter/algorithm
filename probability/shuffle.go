package probability

import (
	"math/rand"
	"time"
)

/* 洗牌算法 */

// 简单不放回抽样，每个样本都是等概论的
// 保证概率相等 所以是随机算法
func Shuffle(card []int) {
	rand.Seed(time.Now().Unix())
	n := len(card)
	for i := 0; i < n; i++ {
		index := rand.Intn(n-i) + i
		card[i], card[index] = card[index], card[i]
	}
}
