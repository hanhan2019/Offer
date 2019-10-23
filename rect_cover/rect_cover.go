//我们可以用2*1的小矩形横着或者竖着去覆盖更大的矩形。请问用n个2*1的小矩形无重叠地覆盖一个2*n的大矩形，总共有多少种方法？
//分析：
// 有以下几种情形：
// 1⃣️target <= 0 大矩形为<= 2*0,直接return 1；
// 2⃣️target = 1大矩形为2*1，只有一种摆放方法，return1；
// 3⃣️target = 2 大矩形为2*2，有两种摆放方法，return2；
// 4⃣️target = n 分为两步考虑：
// 第一次摆放一块 2*1 的小矩阵，则摆放方法总共为f(target - 1)
// 1 0 0 0 0 0 0 0 0 。。。。。。
// 1 0 0 0 0 0 0 0 0 。。。。。。

// 第一次摆放一块1*2的小矩阵，则摆放方法总共为f(target-2)
// 因为，摆放了一块1*2的小矩阵（用√√表示），对应下方的1*2（用××表示）摆放方法就确定了，所以为f(targte-2)
// 1 1 0 0 0 0 0 0 0 。。。。。。
// x x 0 0 0 0 0 0 0 。。。。。。
//所以就是斐波拉切数列！！！
package main

func main() {

}

func rect_cover_dp(n int) int {
	if n < 2 {
		return n
	}
	pre, cur := 0, 1
	for i := 0; i < n-1; i++ {
		count := pre + cur
		pre = cur
		cur = count
	}
	return cur
}
