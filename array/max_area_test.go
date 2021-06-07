package array

import (
	"testing"
)

//给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，
//选择两堵墙，和 x 轴构成的容器可以容纳最多的水。


func TestMaxAreaFunc(t *testing.T) {
	arr := []int{5, 3, 9}
	maxArea := maxArea(arr)
	t.Logf("Max Area=%d", maxArea)
}