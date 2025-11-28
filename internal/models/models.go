package models

type Component struct {
	Name string
}


type Metadata struct {
	PackageName string
	Children    []string
	Usages      []string
}

type ItemChangedMsg struct {
	Name string
}

type ComponentMetaMsg struct {
	Metadata Metadata
}

func (c Component) FilterValue() string {
	return c.Name
}

func (c Component) Title() string {
	return c.Name
}

func (c Component) Description() string {
	return ""
}
