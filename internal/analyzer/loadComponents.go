package analyzer

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

type Component struct {
	Name     string
	Children []Component
}

func LoadComponents(rootPath string) ([]Component, error) {
	var components []Component
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports | packages.NeedDeps,
	}

	pkgs, err := packages.Load(cfg, rootPath+"/...")
	if err != nil {
		return components, fmt.Errorf("failed to load packages: %w", err)
	}

	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()

		for _, name := range scope.Names() {
			obj := scope.Lookup(name)
			if _, ok := obj.Type().Underlying().(*types.Struct); ok {
				ptr := types.NewPointer(obj.Type())
				if IsBubbleTeaModel(ptr) {
					fmt.Println(pkg.Name)
					components = append(components, Component{Name: name})
				}
			}
		}
	}

	return components, nil
}

func IsBubbleTeaModel(t types.Type) bool {
	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	pkg := named.Obj().Pkg()

	if !checkInit(t, pkg) {
		return false
	}

	return true
}

func checkMethod(t types.Type, pkg *types.Package, name string, params, results int) *types.Signature {
	obj, _, _ := types.LookupFieldOrMethod(t, true, pkg, name)
	if obj == nil {
		return nil
	}

	if fun, ok := obj.(*types.Func); ok {
		if sig, ok := fun.Type().(*types.Signature); ok {
			if sig.Results().Len() == results && sig.Params().Len() == params {
				return sig
			}
		}
	}

	return nil

}

func checkInit(t types.Type, pkg *types.Package) bool {
	sig := checkMethod(t, pkg, "Init", 0, 1)
	if sig == nil {
		return false
	}

	sig.Results().At(0).Type()

	return false
}
