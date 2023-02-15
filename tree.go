package gobehaviortree

type (
	OnExecTreeFunc func(tree *Tree)
	OnTreeChildrenFinishFunc func(tree *Tree, result Result, childrenIdx int, owner string)
)

//Tree tree describe
type Tree struct {
	root *Root
	onExec OnExecTreeFunc
	onChildrenFinish OnTreeChildrenFinishFunc
}

//NewTree new tree
func NewTree(onExec OnExecTreeFunc, onChildrenFinish OnTreeChildrenFinishFunc) *Tree {
	return &Tree{
		onExec: onExec,
		onChildrenFinish: onChildrenFinish,
	}
}

//SetRoot setroot
func (t *Tree) SetRoot(node *Root) {
	t.root = node
	t.root.SetTree(t)
}

//Run run
func (t *Tree) Run() {
	t.exec()
}

func (t *Tree) exec() {
	t.emitOnExec()

	if !t.root.IsInit() {
		t.root.OnInstall()
	}

	t.root.OnEnter()
}

func (t *Tree) emitOnExec() {
	if t.onExec == nil {
		return
	}
	t.onExec(t)
}

//OnChildrenFinish chidren finished
func (t *Tree) EmitOnChildrenFinish(result Result, childrenIdx int, owner string) {
	if t.onChildrenFinish == nil {
		return
	}
	t.onChildrenFinish(t, result, childrenIdx, owner)
}
