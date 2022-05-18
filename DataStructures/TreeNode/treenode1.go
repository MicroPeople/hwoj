package TreeNode

import "fmt"

//二叉树：每个节点最多只有两个儿子节点的树。
//
//满二叉树：叶子节点与叶子节点之间的高度差为 0 的二叉树，即整棵树是满的，树呈满三角形结构。在国外的定义，非叶子节点儿子都是满的树就是满二叉树。我们以国内为准。
//
//完全二叉树：完全二叉树是由满二叉树而引出来的，设二叉树的深度为 k，除第 k 层外，其他各层的节点数都达到最大值，且第 k 层所有的节点都连续集中在最左边。
//
//二叉树结构定义
//
// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}

//树的遍历
//
//1、先序遍历：先访问根节点，再访问左子树，最后访问右子树。
//
//2、后序遍历：先访问左子树，再访问右子树，最后访问根节点。
//
//3、中序遍历：先访问左子树，再访问根节点，最后访问右子树。
//
//4、层次遍历：每一层从左到右访问每一个节点。

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印左子树
	PreOrder(tree.Left)
	// 再打印右字树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印右字树
	MidOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印左子树
	MidOrder(tree.Left)
	// 再打印右字树
	MidOrder(tree.Right)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
}

//按层遍历：

func Level(head *TreeNode) {

	if head == nil {
		return
	}

	// 用切片模仿队列
	var queue []*TreeNode
	queue = append(queue, head)

	for len(queue) != 0 {
		// 队头弹出，再把队头切掉，模仿队列的poll操作
		cur := queue[0]
		queue = queue[1:]

		fmt.Printf("%d", (*cur).Data)

		// 当前节点有左孩子，加入左孩子进队列
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		// 当前节点有右孩子，加入右孩子进队列
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}

}
