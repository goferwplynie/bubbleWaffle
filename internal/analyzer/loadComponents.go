package analyzer

import (
	"fmt"
	"go/types"
	"time"

	"golang.org/x/tools/go/packages"
)

type Component struct {
	Name     string
	Children []Component
	Parents  []Component
}

func LoadComponents(rootPath string) ([]Component, error) {
	start := time.Now()
	var components []Component
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles | packages.NeedTypes | packages.NeedTypesInfo,
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
					components = append(components, Component{Name: pkg.Name})
				}
			}
		}
	}

	fmt.Println(time.Since(start))
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

	if !checkUpdate(t, pkg) {
		return false
	}

	if !checkView(t, pkg) {
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

func isTeaCmd(t types.Type) bool {
	if sig, ok := t.Underlying().(*types.Signature); ok {
		if sig.Results().Len() == 1 {
			return true
		}
	}
	return false
}

func checkInit(t types.Type, pkg *types.Package) bool {
	sig := checkMethod(t, pkg, "Init", 0, 1)
	if sig == nil {
		return false
	}

	restype := sig.Results().At(0).Type()
	if !isTeaCmd(restype) {
		return false
	}
	return true
}

func checkUpdate(t types.Type, pkg *types.Package) bool {
	sig := checkMethod(t, pkg, "Update", 1, 2)
	if sig == nil {
		return false
	}

	var implements bool
	res1 := sig.Results().At(0).Type()

	if types.Identical(res1, t) {
		implements = true
	} else if ptr, ok := t.(*types.Pointer); ok && types.Identical(res1, ptr.Elem()) {
		implements = true
	}

	if !implements {
		return false
	}

	res2 := sig.Results().At(1).Type()
	if !isTeaCmd(res2) {
		return false
	}

	return true
}

func checkView(t types.Type, pkg *types.Package) bool {
	sig := checkMethod(t, pkg, "View", 0, 1)

	res := sig.Results().At(0).Type()

	if basic, ok := res.(*types.Basic); ok {
		if basic.Kind() != types.String {
			return false
		}
	}

	return true
}
