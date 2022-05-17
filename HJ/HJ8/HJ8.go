package main

import (
	"fmt"
	"sort"
)

func main() {
	hash := make(map[int]int)
	var key, value, num int

	// 接受输入
	fmt.Scan(&num)
	for i := num; i != 0; i-- {
		fmt.Scan(&key, &value)
		hash[key] += value
	}

	// 因为go 的hash是无序的
	// 需要将hash中的数据放到数组中排序
	nums := make([][]int, len(hash))
	numsInx := 0
	for k, v := range hash {
		nums[numsInx] = []int{k, v}
		numsInx++
	}

	// 调用标准库排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	// 输出
	for _, v := range nums {
		fmt.Printf("%d %d\n", v[0], v[1])
	}
}
func HJ8() {
	n := 0
	k, v := 0, 0
	H := make(map[int]int)

	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		fmt.Scanf("%d %d", &k, &v)
		if sum, ok := H[k]; ok {
			H[k] = sum + v
		} else {
			H[k] = v
		}
	}
	// 将keys放入数组并排序
	var s []int
	for k := range H {
		s = append(s, k)
	}
	sort.Ints(s)
	for _, k := range s {
		fmt.Printf("%d %d\n", k, H[k])
	}
}
