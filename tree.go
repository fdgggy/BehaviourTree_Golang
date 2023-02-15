package gobehaviortree

import (
	"fmt"
)

//Tree tree describe
type Tree struct {
	root Root
}

//NewTree new tree
func NewTree() *Tree {
	return &Tree{}
}

//SetRoot setroot
func (t *Tree) SetRoot(node Root) {
	t.root = node
	t.root.SetTree(t)
}

//Run run
func (t *Tree) Run() {
	t.excuser()
}

func (t *Tree) excuser() {
	fmt.Println("tree excuser")
	if t.root.IsInit() == false {
		t.root.OnInstall()
	}

	t.root.OnEnter()
}

//OnChildrenFinish chidren finished
func (t *Tree) OnChildrenFinish(result Result, childrenIndex int, owner string) {
	fmt.Printf("node:%s run over!\n", owner)
}
