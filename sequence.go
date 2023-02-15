package gobehaviortree

//SequenceAINode sequence node
type SequenceAINode struct {
	AINode
}

//NewSequence construct
func NewSequence() *SequenceAINode {
	return &SequenceAINode{
		AINode: AINode{Name: "SequenceAINode"},
	}
}

//AddNode add node
func (s *SequenceAINode) AddNodes(nodes ...BaseNode) {
	for i, v := range nodes {
		v.SetIdx(i)
		v.SetParent(s)
		s.nodeList = append(s.nodeList, v)
	}
	s.ChildCount = len(nodes)
}

//OnInstall preEnter
func (s *SequenceAINode) OnInstall() {
	s.HasInit = true
}

//OnEnter enter
func (s *SequenceAINode) OnEnter() {
	if len(s.nodeList) > 0 {
		s.curNode = s.nodeList[0]
		s.ExecNode(s.curNode)
	} else {
		s.OnExit()
		s.SendParentResult(s, ResultSuccess)
	}
}

//OnChildrenFinish child is done
func (s *SequenceAINode) OnChildrenFinish(result Result, childrenIdx int, owner string) {
	if result == ResultFailed {
		s.SendParentResult(s, ResultFailed)
		return
	}
	if childrenIdx < len(s.nodeList)-1 {
		s.curNode = s.nodeList[childrenIdx+1]
		s.ExecNode(s.curNode)
	} else {
		s.SendParentResult(s, ResultSuccess)
	}
}

//OnExit exit
func (s *SequenceAINode) OnExit() {
	s.curNode = nil
}

//OnUninstall uninstall
func (s *SequenceAINode) OnUninstall() {
	for _, v := range s.nodeList {
		v.OnUninstall()
	}
	s.nodeList = s.nodeList[:0]
}

//WhoAmI return ownerself
func (s *SequenceAINode) WhoAmI() (am string) {
	return s.Name
}

//SetIdx setidx
func (s *SequenceAINode) SetIdx(idx int) {
	s.IdxInParent = idx
}

//SetParent set parent
func (s *SequenceAINode) SetParent(parent BaseNode) {
	s.Parent = parent
}

//IsInit is init?
func (s *SequenceAINode) IsInit() (init bool) {
	return s.HasInit
}

//Print dump all
func (s *SequenceAINode) Print() {
	for _, v := range s.nodeList {
		v.Print()
	}
}
