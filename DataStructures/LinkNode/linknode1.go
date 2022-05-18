package LinkNode

import "fmt"

type LinkNode1 struct {
	Data     int64
	NextNode *LinkNode1
}

func NewLinkNode1() {
	node := new(LinkNode1)
	node.Data = 1

	node1 := new(LinkNode1)
	node1.Data = 2
	node.NextNode = node1 // node1 链接到 node 节点上

	node2 := new(LinkNode1)
	node2.Data = 3
	node1.NextNode = node2 // node2 链接到 node1 节点上

	// 顺序打印。把原链表头结点赋值到新的NowNode上
	// 这样仍然保留了原链表头结点node不变
	nowNode := node
	for nowNode != nil {
		fmt.Println(nowNode.Data)
		// 获取下一个节点。链表向下滑动
		nowNode = nowNode.NextNode
	}
}
