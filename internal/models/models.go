package models

type Component struct {
	Name string
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
