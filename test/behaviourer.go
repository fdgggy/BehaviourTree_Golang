package test

import (
	"github.com/995933447/gobehaviortree"
)

//Behaviourer behaviourer
type Behaviourer struct {
	executor *gobehaviortree.Tree
}

//Init init
func (b *Behaviourer) Init() {
	b.executor = gobehaviortree.NewTree()
	root := gobehaviortree.NewRoot()

	seq := gobehaviortree.NewSequence()
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

func NewBehaviour() *Behaviourer {
	return &Behaviourer{}
}
