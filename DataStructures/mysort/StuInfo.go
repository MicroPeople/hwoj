package main

//数据集合排序
//
//我们在对数据集合进行排序之前，需要实现sort.Interface接口的三个方法，来看一下其定义：
//type Interface interface {
//	// 获取元素集合个数元素
//	Len() int
//	// 如果 i 索引的数据小于 j 所有的数据，返回 true，不会调用
//	Less(i, j int) bool
//	// 交换 i 和 j 索引得两个元素的位置
//	Swap(i, j int)
//}
//
//我们在实现了这三个方法后，就可以调用Sort()方法来进行排序了。其定义也很简单，只需要一个参数，也就是待排序的数据集合。
//func Sort(data interface)
//
//该包中还提供了一个IsSorted()函数，用来判断数据集合是否已经拍好顺序。
//func IsSorted(data interface) {
//	n := data.Len()
//	for i := n - 1; i > 0; i-- {
//		if data.Less(i, i-1) {
//			return false
//		}
//	}
//	return true
//}

import (
	"fmt"
	"sort"
)

type StuInfo struct {
	name  string
	score int
}

type StuInfos []StuInfo

func (s StuInfos) Len() int {
	return len(s)
}

func (s StuInfos) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuInfos{
		{"Alice", 95},
		{"Bob", 91},
		{"Cindy", 96},
		{"David", 90},
	}
	fmt.Println("=====Original=====")
	for _, stu := range stus {
		fmt.Println(stu.name, ": ", stu.score)
	}
	fmt.Println()

	sort.Sort(stus)

	fmt.Println("=====After Sort=====")
	for _, stu := range stus {
		fmt.Println(stu.name, ": ", stu.score)
	}

	fmt.Println("Has been sorted?", sort.IsSorted(stus))
}

/** The result is:
David :  90
Bob :  91
Alice :  95
Cindy :  96
Has been sorted? true
*/
//但其实还有一种方法，也很简单，那么就是Reverse()函数，这个函数我们在学习数据结构或者其他语言的时候会经常对反转函数这么命名，那么有了这个函数我们就不需要再修改Less()函数了。
//func Reverse(data interface) Interface
//
//在这里需要注意的是，Reverse()函数返回的是一个接口。
//那么这样一来，我们就可以在刚才那个例子中使用Reverse()来实现成绩升序排序：
//sort.Sort(sort.Reverse(stus))
//for _, stu := range stus {
//fmt.Println(stu.name, ": ", stu.score)
//}
//
//最后还有一个方法：Search()，该方法使用“二分查找” 算法来搜索某指定切片[0:n],定义如下：
//func Search(n int, f func(int) bool) int
//
//Search()函数的一个常用方式是搜索元素x是否已经在升序排好的切片s中：
//x := 11
//s := []int{3, 6, 8, 11, 45}
//pos := sort.Search(len(s), func(i int) bool {return s[i] >= x })
//if pos < len(s) && s[pos] == x {
//fmt.Println("The position of ", x, "is: ", pos)
//} else {
//fmt.Println(x, " is not in slice.")
//}
//
