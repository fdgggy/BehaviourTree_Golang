package test

import (
	"fmt"
	"github.com/995933447/gobehaviortree"
)

type Action1 struct {
	gobehaviortree.AINode
}

func NewAction1() *Action1 {
	return &Action1{
		gobehaviortree.AINode{Owner: "action1"},
	}
}
func (a *Action1) OnInstall() {
	fmt.Println("Action1 OnInstall")
	a.IsAlInit = true
}
func (a *Action1) OnUninstall() {
	fmt.Println("Action1 OnUninstall")
	a.IsAlInit = false
}
func (a *Action1) OnEnter() {
	fmt.Println("Action1 OnEnter")
	a.SendParentResult(a, gobehaviortree.ResultSuccess)
}

func (a *Action1) OnExit() {
	fmt.Println("Action1 OnExit")
}

func (a *Action1) SetChildCount(count int) {
	fmt.Println("Action1 SetChildCount")
}

func (a *Action1) SetInit(init bool) {
	fmt.Println("Action1 SetInit")
}
func (a *Action1) SetParent(p gobehaviortree.BaseNode) {
	fmt.Println("Action1 SetParent")
	a.Parent = p
}

func (a *Action1) SetIndex(index int) {
	a.IdxInParent = index
}

func (a *Action1) IsInit() (init bool) {
	return a.IsAlInit
}
func (a *Action1) WhoAmI() (am string) {
	return a.Owner
}
func (a *Action1) ToString() {
	fmt.Printf("Owner:%s IdxInParent:%d IsAlInit:%t ParentName:%s\n\n", a.Owner, a.IdxInParent, a.IsAlInit, a.Parent.WhoAmI())
}
