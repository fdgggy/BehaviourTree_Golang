package test

import "testing"

func TestTree(t *testing.T) {
	behaviout := NewBehaviour()
	behaviout.Init()
	behaviout.Run()
}
