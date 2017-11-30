package baseai

import (
	"fmt"
)

//ResultType result enums
type ResultType uint16

const (
	B_FALSE ResultType = iota
	B_TRUE
	B_RUNING
)

//BaseNode basenode
type BaseNode interface {
	IsInit() bool
	SetParent(parent BaseNode)
	SetIndex(index int)
	Tostring()
	WhoAmI() string
	SetTree(t *Tree)

	OnInstall()
	OnUnstall()
	OnEnter()
	OnExit()
	ExcuserNode(child BaseNode)
	SendParentResult(exitNode BaseNode, result ResultType)
	OnChildrenFinish(result ResultType, childIndex int, owner string)
}

//AiNode node describe
type AiNode struct {
	ChildCount  int
	IsAlInit    bool
	Parent      BaseNode
	nodeList    []BaseNode
	curNode     BaseNode
	IdxInParent int
	Owner       string
}

//ExcuserNode excuse a node
func (a *AiNode) ExcuserNode(child BaseNode) {
	if child == nil {
		fmt.Println("the child node is nil!")
		return
	}
	fmt.Printf("excuse %s node\n", child.WhoAmI())

	if child.IsInit() == false {
		child.OnInstall()
	}

	child.OnEnter()
}

//SendParentResult send to him parent
func (a *AiNode) SendParentResult(exitNode BaseNode, result ResultType) {
	exitNode.OnExit()

	if a.Parent != nil {
		a.Parent.OnChildrenFinish(result, a.IdxInParent, a.Owner)
	} else {
		fmt.Printf("node:%s, a.Parent is nil\n", a.Owner)
	}
}

//OnChildrenFinish don't delete
func (a *AiNode) OnChildrenFinish(result ResultType, childIndex int, owner string) {
}

//SetTree don't delete
func (a *AiNode) SetTree(t *Tree) {
}
