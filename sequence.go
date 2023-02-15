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
		s.ChildCount++
		v.OnInstall()
	}
}

//OnInstall preEnter
func (s *SequenceAINode) OnInstall() {
}

//OnEnter enter
func (s *SequenceAINode) OnEnter() Result {
	if len(s.nodeList) > 0 {
		for _, node := range s.nodeList {
			if s.curNode != nil && node.GetIdx() <= s.curNode.GetIdx() {
				continue
			}

			s.curNode = node
			res := node.OnEnter()
			if res != ResultSuccess {
				if res == ResultFailed || s.curNode.GetIdx() == len(s.nodeList) {
					s.curNode = nil
				}

				s.OnExit()
				return res
			}
		}
	}

	s.OnExit()
	s.curNode = nil
	return ResultSuccess
}

//OnExit exit
func (s *SequenceAINode) OnExit() {
}

//OnUninstall uninstall
func (s *SequenceAINode) OnUninstall() {
	for _, v := range s.nodeList {
		v.OnUninstall()
	}
	s.nodeList = nil
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

//Print dump all
func (s *SequenceAINode) Print() {
	for _, v := range s.nodeList {
		v.Print()
	}
}
