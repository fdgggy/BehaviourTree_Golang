package gobehaviortree

import (
	"fmt"
)

//Result result enums
type Result int

const (
	ResultNil Result = iota
	ResultFailed
	ResultSuccess
	ResultRunning
)

//BaseNode basenode
type BaseNode interface {
	IsInit() bool
	SetParent(parent BaseNode)
	SetIndex(index int)
	ToString()
	WhoAmI() string
	SetTree(t *Tree)

	OnInstall()
	OnUninstall()
	OnEnter()
	OnExit()
	ExecNode(child BaseNode)
	SendParentResult(exitNode BaseNode, result Result)
	OnChildrenFinish(result Result, childIndex int, owner string)
}

//AINode node describe
type AINode struct {
	ChildCount  int
	IsAlInit    bool
	Parent      BaseNode
	nodeList    []BaseNode
	curNode     BaseNode
	IdxInParent int
	Owner       string
}

//ExecNode excuse a node
func (a *AINode) ExecNode(child BaseNode) {
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
func (a *AINode) SendParentResult(exitNode BaseNode, result Result) {
	exitNode.OnExit()

	if a.Parent != nil {
		a.Parent.OnChildrenFinish(result, a.IdxInParent, a.Owner)
	} else {
		fmt.Printf("node:%s, a.Parent is nil\n", a.Owner)
	}
}

//OnChildrenFinish don't delete
func (a *AINode) OnChildrenFinish(result Result, childIndex int, owner string) {
}

//SetTree don't delete
func (a *AINode) SetTree(t *Tree) {
}
