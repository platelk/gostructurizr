package gostructurizr

type GroupNode[K any] struct {
	value K
}

func Group[K any]() *GroupNode[K] {
	return &GroupNode[K]{}
}
