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
	SetParent(parent BaseNode)
	SetIdx(idx int)
	GetIdx() int
	Print()
	WhoAmI() string
	SetTree(t *Tree)

	OnInstall()
	OnUninstall()
	OnEnter() Result
	OnExit()
}

//AINode node describe
type AINode struct {
	ChildCount  int
	Parent      BaseNode
	nodeList    []BaseNode
	curNode     BaseNode
	IdxInParent int
	Name   		string
	Tree   		*Tree
}

//SetTree don't delete
func (a *AINode) SetTree(tree *Tree) {
	a.Tree = tree
}

func (a *AINode) SetIdx(idx int) {
	a.IdxInParent = idx
}

func (a *AINode) GetIdx() int {
	return a.IdxInParent
}

func (a *AINode) WhoAmI() string {
	return a.Name
}

func (a *AINode) SetParent(parent BaseNode) {
	a.Parent = parent
}
