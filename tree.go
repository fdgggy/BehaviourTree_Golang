package gobehaviortree

type (
	OnExecTreeFunc func(tree *Tree)
	OnTreeChildFinishFunc func(tree *Tree, result Result, root *Root)
)

//Tree tree describe
type Tree struct {
	root *Root
	onExec OnExecTreeFunc
	onChildFinish OnTreeChildFinishFunc
}

//NewTree new tree
func NewTree(onExec OnExecTreeFunc, onChildFinish OnTreeChildFinishFunc) *Tree {
	return &Tree{
		onExec: onExec,
		onChildFinish: onChildFinish,
	}
}

//SetRoot setroot
func (t *Tree) SetRoot(node *Root) {
	if t.root != nil {
		t.root.OnUninstall()
	}
	t.root = node
	t.root.SetTree(t)
	t.root.OnInstall()
}

//Run run
func (t *Tree) Run() {
	t.exec()
}

func (t *Tree) exec() {
	if t.onExec != nil {
		t.onExec(t)
	}

	res := t.root.OnEnter()
	if t.onChildFinish != nil {
		t.onChildFinish(t, res, t.root)
	}
}