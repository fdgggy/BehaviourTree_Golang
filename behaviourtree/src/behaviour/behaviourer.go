package behaviour

import (
	"behaviourtree/src/baseai"
)

//Behaviourer behaviourer
type Behaviourer struct {
	executor *baseai.Tree
}

//Init init
func (b *Behaviourer) Init() {
	b.executor = baseai.NewTree()
	root := baseai.NewRootNode()

	seq := baseai.NewSequence()
	action1 := NewAction1()
	action2 := NewAction2()
	seq.AddNode(action1, action2)

	root.AddNode(seq)
	b.executor.SetRoot(*root)
}

//Run run
func (b *Behaviourer) Run() {
	b.executor.Run()
}
