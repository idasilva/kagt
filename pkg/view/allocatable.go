package view

import "github.com/idasilva/kagt/pkg/view/node"

type Allocatable struct {
	node.Node
}

func New() *Allocatable {

	return &Allocatable{}
}
