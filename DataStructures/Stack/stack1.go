package Stack

import (
	"fmt"
	"sync"
)

// unsafe stack
func main() {
	// 用切片制作一个栈
	var stack []int
	// 元素1 入栈
	stack = append(stack, 1, 5, 7, 2)
	// 栈取出最近添加的数据。例如[1,5,7,2] ，len = 4
	x := stack[len(stack)-1] // 2
	// 切掉最近添加的数据，上一步和这一步模仿栈的pop。
	stack = stack[:len(stack)-1] // [1,5,7]
	fmt.Printf("%d", x)
}

// 数组栈，后进先出
type Mystack struct {
	array []string   // 底层切片
	size  int        // 栈的元素数量
	lock  sync.Mutex // 为了并发安全使用的锁
}

func (stack *Mystack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 放入切片中，后进的元素放在数组最后面
	stack.array = append(stack.array, v)

	// 栈中元素数量+1
	stack.size = stack.size + 1
}

//1、如果切片偏移量向前移动 stack.array[0 : stack.size-1]，表明最后的元素已经不属于该数组了，数组变相的缩容了。此时，切片被缩容的部分并不会被回收，仍然占用着空间，所以空间复杂度较高，但操作的时间复杂度为：O(1)。

//2、如果我们创建新的数组 newArray，然后把老数组的元素复制到新数组，就不会占用多余的空间，但移动次数过多，时间复杂度为：O(n)。

// Pop get the top element of stack
func (stack *Mystack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素
	v := stack.array[stack.size-1]

	// 切片收缩，但可能占用空间越来越大
	//stack.array = stack.array[0 : stack.size-1]

	// 创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	// 栈中元素数量-1
	stack.size = stack.size - 1
	return v
}

// Peek 获取栈顶元素
func (stack *Mystack) Peek() string {
	// 栈中元素已空
	if stack.size == 0 {
		panic("empty")
	}

	// 栈顶元素值
	v := stack.array[stack.size-1]
	return v
}

// Size 栈大小
func (stack *Mystack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *Mystack) IsEmpty() bool {
	return stack.size == 0
}
