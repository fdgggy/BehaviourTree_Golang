package gobehaviortree

//Result result enums
type Result int

const (
	ResultNil Result = iota
	ResultFailed
	ResultSuccess
	ResultRunning
)

//BaseNode baseNode
type BaseNode interface {
	IsInit() bool
	SetParent(parent BaseNode)
	SetIdx(idx int)
	Print()
	WhoAmI() string
	SetTree(t *Tree)

	OnInstall()
	OnUninstall()
	OnEnter()
	OnExit()
	ExecNode(child BaseNode)
	SendParentResult(exitNode BaseNode, result Result)
	OnChildrenFinish(result Result, childIdx int, childrenOwnerName string)
}

//AINode node describe
type AINode struct {
	ChildCount  int
	HasInit     bool
	Parent      BaseNode
	nodeList    []BaseNode
	curNode     BaseNode
	IdxInParent int
	Name   string
}

//ExecNode excuse a node
func (a *AINode) ExecNode(child BaseNode) {
	if child == nil {
		return
	}

	if child.IsInit() == false {
		child.OnInstall()
	}

	child.OnEnter()
}

//SendParentResult send to him parent
func (a *AINode) SendParentResult(exitNode BaseNode, result Result) {
	exitNode.OnExit()
	if a.Parent != nil {
		a.Parent.OnChildrenFinish(result, a.IdxInParent, a.Name)
	}
}

//OnChildrenFinish don't delete
func (a *AINode) OnChildrenFinish(result Result, childIdx int, owner string) {
}

//SetTree don't delete
func (a *AINode) SetTree(t *Tree) {
}