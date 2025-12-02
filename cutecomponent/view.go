package cutecomponent

import(
	zone "github.com/lrstanley/bubblezone"
)

func (m *Model) View() string{
	var layout string
	return zone.Mark(layout)
}
	