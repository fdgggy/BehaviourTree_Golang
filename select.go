package gobehaviortree

//SequenceAINode sequence node
type SelectAINode struct {
	AINode
}

//NewSequence construct
func NewSelector() *SelectAINode {
	return &SelectAINode{
		AINode: AINode{Name: "SelectAINode"},
	}
}

//AddNode add node
func (s *SelectAINode) AddNodes(nodes ...BaseNode) {
	for i, v := range nodes {
		v.SetIdx(i)
		v.SetParent(s)
		s.nodeList = append(s.nodeList, v)
		s.ChildCount++
		v.OnInstall()
	}
}

//OnInstall preEnter
func (s *SelectAINode) OnInstall() {
}

//OnEnter enter
func (s *SelectAINode) OnEnter() Result {
	defer func() {
		s.curNode = nil
	}()

	if len(s.nodeList) > 0 {
		for _, node := range s.nodeList {
			s.curNode = node
			res := node.OnEnter()
			if res != ResultFailed {
				s.OnExit()
				return res
			}
		}
	}

	s.OnExit()
	return ResultSuccess
}

//OnExit exit
func (s *SelectAINode) OnExit() {
}

//OnUninstall uninstall
func (s *SelectAINode) OnUninstall() {
	for _, v := range s.nodeList {
		v.OnUninstall()
	}
	s.nodeList = nil
}

//WhoAmI return ownerself
func (s *SelectAINode) WhoAmI() (am string) {
	return s.Name
}

//SetIdx setidx
func (s *SelectAINode) SetIdx(idx int) {
	s.IdxInParent = idx
}

//SetParent set parent
func (s *SelectAINode) SetParent(parent BaseNode) {
	s.Parent = parent
}

//Print dump all
func (s *SelectAINode) Print() {
	for _, v := range s.nodeList {
		v.Print()
	}
}


