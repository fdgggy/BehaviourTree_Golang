package test

import (
	"fmt"
	"github.com/995933447/gobehaviortree"
)

type Action2 struct {
	gobehaviortree.AINode
}

func NewAction2() *Action2 {
	return &Action2{
		gobehaviortree.AINode{Owner: "action2"},
	}
}
func (a *Action2) OnInstall() {
	fmt.Println("Action2 OnInstall")
	a.IsAlInit = true
}
func (a *Action2) OnUninstall() {
	fmt.Println("Action2 OnUninstall")
	a.IsAlInit = false
}
func (a *Action2) OnEnter() {
	fmt.Println("Action2 OnEnter")
	a.SendParentResult(a, gobehaviortree.ResultFailed)
}

func (a *Action2) OnExit() {
	fmt.Println("Action2 OnExit")
}

func (a *Action2) SetChildCount(count int) {
	fmt.Println("Action2 SetChildCount")
}

func (a *Action2) SetInit(init bool) {
	fmt.Println("Action2 SetInit")
}
func (a *Action2) SetParent(p gobehaviortree.BaseNode) {
	fmt.Println("Action2 SetParent")
	a.Parent = p
}

func (a *Action2) SetIndex(index int) {
	a.IdxInParent = index
}

func (a *Action2) IsInit() (init bool) {
	return a.IsAlInit
}
func (a *Action2) WhoAmI() (am string) {
	return a.Owner
}
func (a *Action2) ToString() {
	fmt.Printf("Owner:%s IdxInParent:%d IsAlInit:%t ParentName:%s\n\n", a.Owner, a.IdxInParent, a.IsAlInit, a.Parent.WhoAmI())
}
