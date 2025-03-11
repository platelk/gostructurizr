package gostructurizr

// FilterMode represents the inclusion or exclusion mode for filtered views
type FilterMode string

const (
	Include FilterMode = "Include" // Include elements that match the filter
	Exclude FilterMode = "Exclude" // Exclude elements that match the filter
)

// FilterType represents what kind of filter to apply
type FilterType string

const (
	TagFilter      FilterType = "Tag"      // Filter by tag
	NameFilter     FilterType = "Name"     // Filter by name
	TypeFilter     FilterType = "Type"     // Filter by element type
	RelationFilter FilterType = "Relation" // Filter by relationship
)

// FilterCriteria defines a single filter criterion
type FilterCriteria struct {
	Mode  FilterMode
	Type  FilterType
	Value string
}

// Viewable interface represents any view that can be used as a base view
type Viewable interface {
	Key() *string
}

// FilteredViewNode represents a filtered view in the architecture
type FilteredViewNode struct {
	ViewNode
	baseView      Viewable
	filterCriteria []FilterCriteria
	title         string
	description   string
	key           string
}

// FilteredView creates a new FilteredViewNode
func FilteredView(baseView Viewable) *FilteredViewNode {
	return &FilteredViewNode{
		ViewNode:      *NewViewNode(),
		baseView:      baseView,
		filterCriteria: []FilterCriteria{},
	}
}

// WithTitle sets the title of the filtered view
func (f *FilteredViewNode) WithTitle(title string) *FilteredViewNode {
	f.title = title
	return f
}

// Title returns the title of the filtered view
func (f *FilteredViewNode) Title() string {
	return f.title
}

// WithDescription sets the description of the filtered view
func (f *FilteredViewNode) WithDescription(description string) *FilteredViewNode {
	f.description = description
	return f
}

// Description returns the description of the filtered view
func (f *FilteredViewNode) Description() string {
	return f.description
}

// WithKey sets the key of the filtered view
func (f *FilteredViewNode) WithKey(key string) *FilteredViewNode {
	f.key = key
	return f
}

// Key returns the key of the filtered view
func (f *FilteredViewNode) Key() string {
	return f.key
}

// BaseView returns the base view that this filtered view is based on
func (f *FilteredViewNode) BaseView() Viewable {
	return f.baseView
}

// WithTagFilter adds a tag filter to the filtered view
func (f *FilteredViewNode) WithTagFilter(tag string, mode FilterMode) *FilteredViewNode {
	f.filterCriteria = append(f.filterCriteria, FilterCriteria{
		Mode:  mode,
		Type:  TagFilter,
		Value: tag,
	})
	return f
}

// WithNameFilter adds a name filter to the filtered view
func (f *FilteredViewNode) WithNameFilter(name string, mode FilterMode) *FilteredViewNode {
	f.filterCriteria = append(f.filterCriteria, FilterCriteria{
		Mode:  mode,
		Type:  NameFilter,
		Value: name,
	})
	return f
}

// WithTypeFilter adds a type filter to the filtered view
func (f *FilteredViewNode) WithTypeFilter(elementType string, mode FilterMode) *FilteredViewNode {
	f.filterCriteria = append(f.filterCriteria, FilterCriteria{
		Mode:  mode,
		Type:  TypeFilter,
		Value: elementType,
	})
	return f
}

// FilterCriteria returns all filter criteria for this filtered view
func (f *FilteredViewNode) FilterCriteria() []FilterCriteria {
	return f.filterCriteria
}

// Include adds a tag inclusion filter
func (f *FilteredViewNode) Include(tag string) *FilteredViewNode {
	return f.WithTagFilter(tag, Include)
}

// Exclude adds a tag exclusion filter
func (f *FilteredViewNode) Exclude(tag string) *FilteredViewNode {
	return f.WithTagFilter(tag, Exclude)
}

// IncludeTag is an alias for Include
func (f *FilteredViewNode) IncludeTag(tag string) *FilteredViewNode {
	return f.Include(tag)
}

// AddIncludeTag is an alias for Include
func (f *FilteredViewNode) AddIncludeTag(tag string) *FilteredViewNode {
	return f.Include(tag)
}

// ExcludeTag is an alias for Exclude
func (f *FilteredViewNode) ExcludeTag(tag string) *FilteredViewNode {
	return f.Exclude(tag)
}

// AddExcludeTag is an alias for Exclude
func (f *FilteredViewNode) AddExcludeTag(tag string) *FilteredViewNode {
	return f.Exclude(tag)
}

// WithAutoLayout sets the auto layout of the filtered view
func (f *FilteredViewNode) WithAutoLayout() *FilteredViewNode {
	f.ViewNode.WithAutoLayout()
	return f
}