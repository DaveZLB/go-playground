package array

//在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。


//巧用map降低时间复杂度到o(n)
func twoSum(arr []int, target int) []int {
	m := make(map[int]int, len(arr))
	for i, _ := range arr {
		t1 := target - arr[i]
		if _, ok := m[t1]; ok {
			return []int{m[t1], i}
		}
		m[arr[i]] = i
	}

	return []int{-1, -1}
}
