package gobehaviortree

import (
	"fmt"
)

//Root only one Root
type Root struct {
	AINode
	owner *Tree
}

//NewRoot Root
func NewRoot() *Root {
	return &Root{}
}

//AddNode addnode
func (r *Root) AddNode(arg BaseNode) {
	fmt.Println("Root AddNode")
	arg.SetIndex(0)
	arg.SetParent(r)
	r.nodeList = append(r.nodeList, arg)
	r.ChildCount = 1
}

//OnInstall install
func (r *Root) OnInstall() {
	fmt.Println("Root OnInstall")
	r.IsAlInit = true
}

//OnEnter enter
func (r *Root) OnEnter() {
	if len(r.nodeList) != 1 {
		fmt.Println("Root's child is not one!")
	} else {
		fmt.Println("Root OnEnter")
		r.curNode = r.nodeList[0]
		r.ExecNode(r.curNode)
	}
}

//OnChildrenFinish children finish
func (r *Root) OnChildrenFinish(result Result, childrenIndex int, owner string) {
	fmt.Printf("%s root childNode return %d\n", owner, result)
	r.owner.OnChildrenFinish(result, childrenIndex, owner)
}

//OnExit exit
func (r *Root) OnExit() {
	r.curNode = nil
}

//OnUninstall uninstall
func (r *Root) OnUninstall() {
	if len(r.nodeList) != 1 {
		fmt.Println("Root's child is not one!")
	} else {
		r.nodeList[0].OnUninstall()
	}

	r.nodeList = r.nodeList[:0]
}

//WhoAmI who
func (r *Root) WhoAmI() (am string) {
	return r.Owner
}

//SetIndex setindex
func (r *Root) SetIndex(index int) {
	r.IdxInParent = index
}

//SetParent setparent
func (r *Root) SetParent(parent BaseNode) {
	r.Parent = parent
}

//SetTree settree
func (r *Root) SetTree(t *Tree) {
	r.owner = t
}

//IsInit isinit?
func (r *Root) IsInit() (init bool) {
	return r.IsAlInit
}

//ToString tostring
func (r *Root) ToString() {
	r.nodeList[0].ToString()
}
