package array

//给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，
//选择两堵墙，和 x 轴构成的容器可以容纳最多的水。

//使用对撞指针思路
func maxArea(arr []int) int {
	start := 0
	end := len(arr) - 1
	max := 0
	for ; start < end; {
		width := end - start
		high := 0
		if arr[start] < arr[end] {
			high = arr[start]
			start++
		} else {
			high = arr[end]
			end--
		}

		area := width * high
		if max < area {
			max = area
		}
	}
	return max
}
