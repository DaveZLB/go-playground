package array

import (
	"testing"
)

//在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

func TestTwoSum(t *testing.T) {
	arr := []int{1, 2, 4, 8, 7}
	target := 15
	indexArr := twoSum(arr, target)
	t.Logf("%+v\n", indexArr)
}