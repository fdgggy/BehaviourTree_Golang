package api

import (
	"behaviourtree/src/behaviour"
)

//NewBehaviour create behaviour tree
func NewBehaviour() *behaviour.Behaviourer {
	return &behaviour.Behaviourer{}
}
