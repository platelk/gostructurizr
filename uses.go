package gostructurizr

type Namer interface {
	Name() string
}

type RelationShipNode struct {
	from, to Namer
	desc     *string
}

func Uses(from, to Namer, desc string) *RelationShipNode {
	return &RelationShipNode{
		from: from,
		to:   to,
		desc: &desc,
	}
}

func (r *RelationShipNode) From() Namer {
	return r.from
}

func (r *RelationShipNode) To() Namer {
	return r.to
}

func (r *RelationShipNode) Description() *string {
	return r.desc
}
