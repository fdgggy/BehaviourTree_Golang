package gobehaviortree

//Root only one Root
type Root struct {
	AINode
	owner *Tree
	onChildFinish func(root *Root, result Result, child BaseNode)
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
	node.OnInstall()
}

//OnInstall install
func (r *Root) OnInstall() {
}

//OnEnter enter
func (r *Root) OnEnter() Result {
	if len(r.nodeList) == 0 {
		return ResultSuccess
	}
	r.curNode = r.nodeList[0]
	res := r.curNode.OnEnter()
	if r.onChildFinish != nil {
		r.onChildFinish(r, res, r.curNode)
	}
	r.OnExit()
	return res
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

//Print tostring
func (r *Root) Print() {
	if len(r.nodeList) == 0 {
		return
	}
	r.nodeList[0].Print()
}
