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
		gobehaviortree.AINode{Name: "action1"},
	}
}

func (a *Action1) OnInstall() {
	fmt.Println("Action1 OnInstall")
	a.HasInit = true
}

func (a *Action1) OnUninstall() {
	fmt.Println("Action1 OnUninstall")
	a.HasInit = false
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
func (a *Action1) SetParent(parent gobehaviortree.BaseNode) {
	fmt.Println("Action1 SetParent")
	a.Parent = parent
}

func (a *Action1) SetIdx(idx int) {
	a.IdxInParent = idx
}

func (a *Action1) IsInit() bool {
	return a.HasInit
}

func (a *Action1) WhoAmI() string {
	return a.Name
}

func (a *Action1) Print() {
	fmt.Printf("Owner:%s IdxInParent:%d HasInit:%t ParentName:%s\n\n", a.Name, a.IdxInParent, a.HasInit, a.Parent.WhoAmI())
}

func (a *Action1) OnChildrenFinish(result gobehaviortree.Result, childIdx int, owner string) {
}