package creator

import (
	"fmt"
	"os"
	"slices"
	"text/template"
)

func CreateComponent(path string, name string, opts ...*ComponentOptions) error {
	components, err := GetComponents(path)
	if err != nil {
		return err
	}
	if slices.Contains(components, path) {
		return fmt.Errorf("component already exists")
	}
	var options *ComponentOptions
	if len(opts) > 0 {
		options = opts[0]
	}
	fmt.Println(options)
	mt := template.Must(template.New("mt").Parse(ModelTemplate))
	it := template.Must(template.New("it").Parse(InitTemplate))
	ut := template.Must(template.New("ut").Parse(UpdateTemplate))
	vt := template.Must(template.New("vt").Parse(ViewTemplate))

	templates := map[string]*template.Template{
		"model.go":  mt,
		"init.go":   it,
		"update.go": ut,
		"view.go":   vt,
	}

	if options != nil {
		if options.StyleFile {
			templates["style.go"] = template.Must(template.New("st").Parse(StylesTemplate))
		}
		if options.KeybindsFile {
			templates["keys.go"] = template.Must(template.New("kt").Parse(KeyBindsTemplate))
		}
		if options.BubbleZone {
			templates["view.go"] = template.Must(template.New("bt").Parse(BubbleZoneView))
		}
	}

	if err := os.Mkdir(name, 0755); err != nil {
		return err
	}

	for fileName, templ := range templates {
		f, err := os.Create(name + "/" + fileName)
		if err != nil {
			return err
		}
		templ.Execute(f, struct {
			Name string
		}{name})
	}

	return nil
}

func GetComponents(path string) ([]string, error) {
	var components []string
	files, err := os.ReadDir(path)
	if err != nil {
		return components, err
	}
	for _, v := range files {
		if v.IsDir() {
			components = append(components, v.Name())
		}
	}
	return components, nil
}
