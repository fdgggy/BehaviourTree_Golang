package main

import "behaviourtree/src/api"

func main() {
	behaviout := api.NewBehaviour()
	behaviout.Init()
	behaviout.Run()
}
