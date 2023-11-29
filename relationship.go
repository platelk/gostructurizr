package gostructurizr

type Namer interface {
	Name() string
}

type InteractionStyle string

const (
	Asynchronous InteractionStyle = "asynchronous"
	Synchronous  InteractionStyle = "synchronous"
)

type RelationShipNode struct {
	from, to         Namer
	desc             *string
	tech             *string
	interactionStyle *InteractionStyle
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

func (r *RelationShipNode) WithTechnology(tech string) *RelationShipNode {
	r.tech = &tech
	return r
}

func (r *RelationShipNode) Technology() *string {
	return r.tech
}

func (r *RelationShipNode) WithInteractionStyle(i InteractionStyle) *RelationShipNode {
	r.interactionStyle = &i
	return r
}
