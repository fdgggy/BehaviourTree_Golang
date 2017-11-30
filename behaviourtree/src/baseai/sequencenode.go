package baseai

import (
	"fmt"
)

//SequenceAINode sequence node
type SequenceAINode struct {
	AiNode
}

//NewSequence construct
func NewSequence() *SequenceAINode {
	return &SequenceAINode{
		AiNode: AiNode{Owner: "SequenceAINode"},
	}
}

//AddNode add node
func (s *SequenceAINode) AddNode(arg ...BaseNode) {
	for i, v := range arg {
		v.SetIndex(i)
		v.SetParent(s)
		s.nodeList = append(s.nodeList, v)
	}
	s.ChildCount = len(arg)
}

//OnInstall preEnter
func (s *SequenceAINode) OnInstall() {
	s.IsAlInit = true
}

//OnEnter enter
func (s *SequenceAINode) OnEnter() {
	if len(s.nodeList) > 0 {
		s.curNode = s.nodeList[0]
		s.ExcuserNode(s.curNode)
	} else {
		s.OnExit()
		s.SendParentResult(s, B_TRUE)
	}
}

//OnChildrenFinish child is done
func (s *SequenceAINode) OnChildrenFinish(result ResultType, childrenIndex int, owner string) {
	fmt.Printf("%s childNode return %d\n", owner, result)

	if result == B_FALSE {
		s.SendParentResult(s, B_FALSE)
		return
	}
	fmt.Printf("childrenIndex:%d childCount:%d\n", childrenIndex, len(s.nodeList))
	if childrenIndex < len(s.nodeList)-1 {
		s.curNode = s.nodeList[childrenIndex+1]
		s.ExcuserNode(s.curNode)
	} else {
		s.SendParentResult(s, B_TRUE)
	}
}

//OnExit exit
func (s *SequenceAINode) OnExit() {
	s.curNode = nil
}

//OnUnstall uninstall
func (s *SequenceAINode) OnUnstall() {
	for _, v := range s.nodeList {
		v.OnUnstall()
	}
	s.nodeList = s.nodeList[:0]
}

//WhoAmI return ownerself
func (s *SequenceAINode) WhoAmI() (am string) {
	return s.Owner
}

//SetIndex setindex
func (s *SequenceAINode) SetIndex(index int) {
	s.IdxInParent = index
}

//SetParent set parent
func (s *SequenceAINode) SetParent(parent BaseNode) {
	s.Parent = parent
}

//IsInit is init?
func (s *SequenceAINode) IsInit() (init bool) {
	return s.IsAlInit
}

//Tostring dump all
func (s *SequenceAINode) Tostring() {
	for _, v := range s.nodeList {
		v.Tostring()
	}
}
