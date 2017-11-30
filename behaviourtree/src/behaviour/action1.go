package behaviour

import (
	"behaviourtree/src/baseai"
	"fmt"
)

type Action1 struct {
	baseai.AiNode
}

func NewAction1() *Action1 {
	return &Action1{
		baseai.AiNode{Owner: "action1"},
	}
}
func (a *Action1) OnInstall() {
	fmt.Println("Action1 OnInstall")
	a.IsAlInit = true
}
func (a *Action1) OnUnstall() {
	fmt.Println("Action1 OnUnstall")
	a.IsAlInit = false
}
func (a *Action1) OnEnter() {
	fmt.Println("Action1 OnEnter")
	a.SendParentResult(a, baseai.B_TRUE)
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
func (a *Action1) SetParent(p baseai.BaseNode) {
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
func (a *Action1) Tostring() {
	fmt.Println("Owner:%s IdxInParent:%d IsAlInit:%t ParentName:%s\n", a.Owner, a.IdxInParent, a.IsAlInit, a.Parent.WhoAmI())
}
