package baseai

import (
	"fmt"
)

//Rootnode only one rootnode
type Rootnode struct {
	AiNode
	owner *Tree
}

//NewRootNode rootnode
func NewRootNode() *Rootnode {
	return &Rootnode{}
}

//AddNode addnode
func (root *Rootnode) AddNode(arg BaseNode) {
	fmt.Println("Rootnode AddNode")
	arg.SetIndex(0)
	arg.SetParent(root)
	root.nodeList = append(root.nodeList, arg)
	root.ChildCount = 1
}

//OnInstall install
func (root *Rootnode) OnInstall() {
	fmt.Println("Rootnode OnInstall")
	root.IsAlInit = true
}

//OnEnter enter
func (root *Rootnode) OnEnter() {
	if len(root.nodeList) != 1 {
		fmt.Println("Rootnode's child is not one!")
	} else {
		fmt.Println("Rootnode OnEnter")
		root.curNode = root.nodeList[0]
		root.ExcuserNode(root.curNode)
	}
}

//OnChildrenFinish children finish
func (root *Rootnode) OnChildrenFinish(result ResultType, childrenIndex int, owner string) {
	fmt.Printf("%s root childNode return %d\n", owner, result)
	root.owner.OnChildrenFinish(result, childrenIndex, owner)
}

//OnExit exit
func (root *Rootnode) OnExit() {
	root.curNode = nil
}

//OnUnstall uninstall
func (root *Rootnode) OnUnstall() {
	if len(root.nodeList) != 1 {
		fmt.Println("Rootnode's child is not one!")
	} else {
		root.nodeList[0].OnUnstall()
	}

	root.nodeList = root.nodeList[:0]
}

//WhoAmI who
func (root *Rootnode) WhoAmI() (am string) {
	return root.Owner
}

//SetIndex setindex
func (root *Rootnode) SetIndex(index int) {
	root.IdxInParent = index
}

//SetParent setparent
func (root *Rootnode) SetParent(parent BaseNode) {
	root.Parent = parent
}

//SetTree settree
func (root *Rootnode) SetTree(t *Tree) {
	root.owner = t
}

//IsInit isinit?
func (root *Rootnode) IsInit() (init bool) {
	return root.IsAlInit
}

//Tostring tostring
func (root *Rootnode) Tostring() {
	root.nodeList[0].Tostring()
}
