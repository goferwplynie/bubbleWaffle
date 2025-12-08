package analyzer

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

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

	if !checkInit(t, named) {
		return false
	}

	if !checkUpdate(t, named, modelInterface, ptr) {
		return false
	}

	if !checkView(t, named) {
		fmt.Println("failed view")
		return false
	}

	return true
}

func verifySignature(t types.Type, pkg *types.Package, name string, params, results int) *types.Signature {
	methodObj, _, _ := types.LookupFieldOrMethod(t, true, pkg, name)
	if methodObj == nil {
		return nil
	}
	if fun, ok := methodObj.(*types.Func); ok {
		if sig, ok := fun.Type().(*types.Signature); ok {
			if sig.Params().Len() == params && sig.Results().Len() == results {
				return sig
			}
		}
	}

	return nil
}

func checkInit(t types.Type, named *types.Named) bool {
	sig := verifySignature(t, named.Obj().Pkg(), "Init", 0, 1)
	if sig == nil {
		return false
	}
	ret := sig.Results().At(0)
	if !isTeaCmd(ret) {
		return false
	}

	return true
}

func isTeaCmd(tVar *types.Var) bool {
	if tVar == nil {
		return false
	}
	if _, ok := tVar.Type().Underlying().(*types.Signature); !ok {
		return false
	}

	return true
}

func checkUpdate(t types.Type, named *types.Named, modelI *types.Interface, ptr *types.Pointer) bool {
	sig := verifySignature(t, named.Obj().Pkg(), "Update", 1, 2)
	if sig == nil {
		return false
	}
	param0 := sig.Params().At(0).Type()
	if _, ok := param0.Underlying().(*types.Interface); !ok {
		return false
	}

	ret1 := sig.Results().At(0).Type()

	isModel := false
	if types.AssignableTo(ret1, modelI) || types.Implements(ret1, modelI) {
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
		if ret1Str == named.Obj().Name() ||
			ret1Str == "*"+named.Obj().Name() ||
			ret1Str == named.Obj().Pkg().Name()+"."+named.Obj().Name() ||
			ret1Str == "*"+named.Obj().Pkg().Name()+"."+named.Obj().Name() {
			isModel = true
		}
	}

	if !isModel {
		return false
	}

	ret2 := sig.Results().At(1)
	if !isTeaCmd(ret2) {
		return false
	}
	return true
}

func checkView(t types.Type, named *types.Named) bool {
	sig := verifySignature(t, named.Obj().Pkg(), "View", 0, 1)
	if sig == nil {
		return false
	}

	ret := sig.Results().At(0).Type()
	isView := false

	if namedRes, ok := ret.(*types.Named); ok {
		if namedRes.Obj().Name() == "View" {
			isView = true
		}
	}
	if basic, ok := ret.(*types.Basic); ok {
		if basic.Kind() == types.String {
			isView = true
		}
	}

	if !isView {
		return false
	}
	return true
}
