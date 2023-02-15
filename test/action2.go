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
		gobehaviortree.AINode{Name: "action2"},
	}
}
func (a *Action2) OnInstall() {
	fmt.Println("Action2 OnInstall")
	a.HasInit = true
}

func (a *Action2) OnUninstall() {
	fmt.Println("Action2 OnUninstall")
	a.HasInit = false
}

func (a *Action2) OnEnter() gobehaviortree.Result {
	fmt.Println("Action2 OnEnter")
	return gobehaviortree.ResultFailed
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
func (a *Action2) SetParent(parent gobehaviortree.BaseNode) {
	fmt.Println("Action2 SetParent")
	a.Parent = parent
}

func (a *Action2) SetIdx(idx int) {
	a.IdxInParent = idx
}

func (a *Action2) IsInit() bool {
	return a.HasInit
}

func (a *Action2) WhoAmI() string {
	return a.Name
}

func (a *Action2) Print() {
	fmt.Printf("Owner:%s IdxInParent:%d HasInit:%t ParentName:%s\n\n", a.Name, a.IdxInParent, a.HasInit, a.Parent.WhoAmI())
}

func (a *Action2) OnChildFinish(result gobehaviortree.Result, childIdx int, owner string) {
}