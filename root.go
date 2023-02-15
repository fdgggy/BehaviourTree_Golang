package gobehaviortree

//Root only one Root
type Root struct {
	AINode
	owner *Tree
}

//NewRoot Root
func NewRoot() *Root {
	return &Root{}
}

//AddNode addnode
func (r *Root) AddNode(node BaseNode) {
	node.SetIdx(0)
	node.SetParent(r)
	r.nodeList = append(r.nodeList, node)
	r.ChildCount = 1
}

//OnInstall install
func (r *Root) OnInstall() {
	r.HasInit = true
}

//OnEnter enter
func (r *Root) OnEnter() {
	if len(r.nodeList) == 0 {
		return
	}
	r.curNode = r.nodeList[0]
	r.ExecNode(r.curNode)
}

//OnChildrenFinish children finish
func (r *Root) OnChildrenFinish(result Result, childrenIdx int, owner string) {
	r.owner.EmitOnChildrenFinish(result, childrenIdx, owner)
}

//OnExit exit
func (r *Root) OnExit() {
	r.curNode = nil
}

//OnUninstall uninstall
func (r *Root) OnUninstall() {
	if len(r.nodeList) == 0 {
		return
	}
	r.nodeList[0].OnUninstall()
	r.nodeList = nil
}

//WhoAmI who
func (r *Root) WhoAmI() (am string) {
	return r.Name
}

//SetIdx setidx
func (r *Root) SetIdx(idx int) {
	r.IdxInParent = idx
}

//SetParent setparent
func (r *Root) SetParent(parent BaseNode) {
	r.Parent = parent
}

//SetTree settree
func (r *Root) SetTree(tree *Tree) {
	r.owner = tree
}

//IsInit isinit?
func (r *Root) IsInit() (init bool) {
	return r.HasInit
}

//Print tostring
func (r *Root) Print() {
	if len(r.nodeList) == 0 {
		return
	}
	r.nodeList[0].Print()
}
