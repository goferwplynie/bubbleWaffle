package analyzer

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type Component struct {
	Name     string
	Children []Component
	Parents  []Component
}

func LoadComponents(rootPath string) ([]Component, error) {
	var components []Component
	var modelInterface *types.Interface
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedDeps,
		Dir:  rootPath,
	}

	pkgs, err := packages.Load(cfg, rootPath+"/...")
	if err != nil {
		return components, fmt.Errorf("failed to load packages: %w", err)
	}

	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		lookup := scope.Lookup("Model")

		modelInterface = (lookup.Type().Underlying().(*types.Interface))
	}

	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()

		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if _, ok := obj.Type().Underlying().(*types.Struct); ok {
				ptr := types.NewPointer(obj.Type())
				if types.Implements(ptr, modelInterface) {
					fmt.Println(name)
				}
			}
		}
	}

	return components, nil
}
