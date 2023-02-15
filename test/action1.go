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
}

func (a *Action1) OnUninstall() {
	fmt.Println("Action1 OnUninstall")
}

func (a *Action1) OnEnter() gobehaviortree.Result {
	fmt.Println("Action1 OnEnter")
	return gobehaviortree.ResultSuccess
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

func (a *Action1) WhoAmI() string {
	return a.Name
}

func (a *Action1) Print() {
	fmt.Printf("Owner:%s IdxInParent:%d ParentName:%s\n\n", a.Name, a.IdxInParent, a.Parent.WhoAmI())
}

func (a *Action1) OnChildFinish(result gobehaviortree.Result, childIdx int, owner string) {
}