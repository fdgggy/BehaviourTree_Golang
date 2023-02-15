package test

import (
	"fmt"
	"github.com/995933447/gobehaviortree"
)

//Behaviourer behaviourer
type Behaviourer struct {
	executor *gobehaviortree.Tree
}

//Init init
func (b *Behaviourer) Init() {
	b.executor = gobehaviortree.NewTree(nil, func(tree *gobehaviortree.Tree, result gobehaviortree.Result, root *gobehaviortree.Root) {
		fmt.Printf("node:%s run over!\n", root.WhoAmI())
	})
	root := gobehaviortree.NewRoot()

	seq := gobehaviortree.NewSequence()
	action1 := NewAction1()
	action2 := NewAction2()
	seq.AddNodes(action1, action2)

	root.AddNode(seq)
	b.executor.SetRoot(root)
}

//Run run
func (b *Behaviourer) Run() {
	b.executor.Run()
}

func NewBehaviour() *Behaviourer {
	return &Behaviourer{}
}
