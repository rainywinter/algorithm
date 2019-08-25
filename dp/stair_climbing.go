package dp

//爬楼梯问题描述：
/*
How to Solve the Staircase Problem

There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.

Given N, write a function that returns the number of unique ways you can climb the staircase.
The order of the steps matters.

For example, if N is 4, then there are 5 unique ways:

1, 1, 1, 1
2, 1, 1
1, 2, 1
1, 1, 2
2, 2
What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from
a set of positive integers X? For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.
*/

var StairClimbingCache = map[int]int{}

//-----------------

// 方法一： f(n) = f(n-1) + f(n-2)
// 时间复杂度 O(n^2)
func StairClimbingRecursion(seq int) int {
	if seq <= 0 {
		return 0
	}
	switch seq {
	case 1:
		return 1
	case 2: // 两个一步或者一个两步，共计两种方式
		return 2
	}
	return StairClimbingRecursion(seq-1) + StairClimbingRecursion(seq-2)
}

// 方法二： 优化一，将已求出的缓存
// 时间空间复杂度 O(n)
func StairClimbingRecursionWithCache(seq int) int {
	if seq <= 0 {
		return 0
	}
	switch seq {
	case 1:
		return 1
	case 2:
		return 2
	}
	if n, ok := StairClimbingCache[seq]; ok {
		return n
	}
	n := StairClimbingRecursionWithCache(seq-1) + StairClimbingRecursionWithCache(seq-2)
	StairClimbingCache[seq] = n
	return n
}

// 方法三： 既然找到状态方式（递推公式），则可以从头开始计算
// f(n)=f(n-1) + f(n-2),f(1)=1,f(2)=2,则可以顺序求出 f(3). f(4)....
// 斐波那契数列
// 时间复杂度O(n)
func StairClimbingForwardPropagate(seq int) int {
	if seq <= 0 {
		return 0
	}
	switch seq {
	case 1:
		return 1
	case 2:
		return 2
	}

	a, b := 1, 2
	var total int
	for i := 3; i <= seq; i++ {
		total = a + b
		// 顺序往后赋值
		a = b
		b = total
	}
	return total
}
