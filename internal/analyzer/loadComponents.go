package analyzer

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"time"

	"golang.org/x/tools/go/packages"
)

type Component struct {
	Name     string
	Children []Component
}

const BubbleTeaMockSource = `
package bubbletea

type Msg interface{}

type Cmd func() Msg

type View struct {}

type Model interface {
	Init() Cmd
	Update(Msg) (Model, Cmd)
	View() View
}
`

// MockImporter intercepts imports for bubbletea and serves a mock package.
type MockImporter struct {
	mockPackage *types.Package
	cache       map[string]*types.Package
}

func NewMockImporter() (*MockImporter, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "mock_bubbletea.go", BubbleTeaMockSource, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mock source: %w", err)
	}

	conf := types.Config{
		Importer: importer.Default(),
	}
	pkg, err := conf.Check("charm.land/bubbletea/v2", fset, []*ast.File{file}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to check mock package: %w", err)
	}

	return &MockImporter{
		mockPackage: pkg,
		cache:       make(map[string]*types.Package),
	}, nil
}

func (m *MockImporter) Import(path string) (*types.Package, error) {
	if path == "charm.land/bubbletea/v2" {
		return m.mockPackage, nil
	}

	name := "mock_" + path
	pkg := types.NewPackage(path, name)
	m.cache[path] = pkg
	return pkg, nil
}

func LoadComponents(rootPath string) ([]Component, error) {
	start := time.Now()
	var components []Component

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax,
	}
	pkgs, err := packages.Load(cfg, rootPath+"/...")
	if err != nil {
		return components, fmt.Errorf("failed to load packages: %w", err)
	}

	mockImporter, err := NewMockImporter()
	if err != nil {
		return components, fmt.Errorf("failed to create mock importer: %w", err)
	}

	obj := mockImporter.mockPackage.Scope().Lookup("Model")
	if obj == nil {
		return components, fmt.Errorf("mock Model not found")
	}
	modelInterface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		return components, fmt.Errorf("mock Model is not an interface")
	}

	for _, pkg := range pkgs {
		if len(pkg.Syntax) == 0 {
			continue
		}

		conf := types.Config{
			Importer: mockImporter,
			Error:    func(err error) {},
		}

		info := &types.Info{
			Defs: make(map[*ast.Ident]types.Object),
			Uses: make(map[*ast.Ident]types.Object),
		}

		conf.Check(pkg.PkgPath, pkg.Fset, pkg.Syntax, info)

		for _, obj := range info.Defs {
			if obj == nil {
				continue
			}
			if _, ok := obj.(*types.TypeName); !ok {
				continue
			}
			if _, ok := obj.Type().Underlying().(*types.Struct); ok {
				ptr := types.NewPointer(obj.Type())
				if IsBubbleTeaModel(ptr, modelInterface) {
					components = append(components, Component{Name: pkg.Name})
				}
			}
		}
	}
	fmt.Println(time.Since(start))
	return components, nil
}

func IsBubbleTeaModel(t types.Type, modelInterface *types.Interface) bool {
	if types.Implements(t, modelInterface) {
		return true
	}
	if types.AssignableTo(t, modelInterface) {
		return true
	}

	ptr, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}

	updateObj, _, _ := types.LookupFieldOrMethod(t, true, named.Obj().Pkg(), "Update")
	if updateObj == nil {
		return false
	}
	if fun, ok := updateObj.(*types.Func); ok {
		if sig, ok := fun.Type().(*types.Signature); ok {
			if sig.Params().Len() != 1 || sig.Results().Len() != 2 {
				return false
			}

			param0 := sig.Params().At(0).Type()
			if _, ok := param0.Underlying().(*types.Interface); !ok {
				return false
			}

			ret1 := sig.Results().At(0).Type()

			isModel := false
			if types.AssignableTo(ret1, modelInterface) || types.Implements(ret1, modelInterface) {
				isModel = true
			}
			if !isModel {
				if types.Identical(ret1, t) {
					isModel = true
				} else if types.Identical(ret1, ptr.Elem()) {
					isModel = true
				}
			}

			if !isModel {
				ret1Str := ret1.String()
				if ret1Str == named.Obj().Name() || ret1Str == "*"+named.Obj().Name() ||
					ret1Str == named.Obj().Pkg().Name()+"."+named.Obj().Name() ||
					ret1Str == "*"+named.Obj().Pkg().Name()+"."+named.Obj().Name() {
					isModel = true
				}
			}

			if !isModel {
				return false
			}

			ret2 := sig.Results().At(1).Type()
			if _, ok := ret2.Underlying().(*types.Signature); !ok {
				return false
			}
		}
	}

	viewObj, _, _ := types.LookupFieldOrMethod(t, true, named.Obj().Pkg(), "View")
	if viewObj == nil {
		return false
	}
	if fun, ok := viewObj.(*types.Func); ok {
		if sig, ok := fun.Type().(*types.Signature); ok {
			if sig.Params().Len() != 0 || sig.Results().Len() != 1 {
				return false
			}
			res0 := sig.Results().At(0).Type()
			isView := false

			if namedRes, ok := res0.(*types.Named); ok {
				if namedRes.Obj().Name() == "View" {
					isView = true
				}
			}

			if !isView {
				return false
			}
		}
	}

	return true
}
