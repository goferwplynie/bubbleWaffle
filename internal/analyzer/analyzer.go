package analyzer

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"strings"

	"github.com/goferwplynie/bubbleWaffle/internal/models"
	"golang.org/x/tools/go/packages"
)

func AnalyzeComponent(componentName string, rootPath string) (models.Metadata, error) {
	meta := models.Metadata{
		Children: []string{},
		Usages:   []string{},
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedFiles,
		Dir:  rootPath,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return meta, fmt.Errorf("failed to load packages: %w", err)
	}

	var componentPackageName string

	//find component and it's children
	for _, pkg := range pkgs {
		//check if any file in the package is in rootPath/componentName
		isComponentPkg := false
		for _, file := range pkg.GoFiles {
			if strings.Contains(filepath.Dir(file), filepath.Join(rootPath, componentName)) {
				isComponentPkg = true
				break
			}
		}

		if isComponentPkg {
			componentPackageName = pkg.Name
			meta.PackageName = pkg.Name

			for _, file := range pkg.Syntax {
				ast.Inspect(file, func(node ast.Node) bool {
					//going through node ~w~
					if typeSpec, ok := node.(*ast.TypeSpec); ok {
						if structType, ok := typeSpec.Type.(*ast.StructType); ok {
							for _, field := range structType.Fields.List {
								if field.Type != nil {
									if selector, ok := field.Type.(*ast.SelectorExpr); ok {
										if strings.HasSuffix(selector.Sel.Name, "Model") {
											if xIdent, ok := selector.X.(*ast.Ident); ok {
												childName := xIdent.Name + "." + selector.Sel.Name
												meta.Children = append(meta.Children, childName)
											}
										}
									}
								}
							}
						}
					}
					return true
				})
			}
			break
		}
	}

	//find usages
	for _, pkg := range pkgs {
		if pkg.Name == componentPackageName {
			continue
		}

		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(node ast.Node) bool {
				//going through components ~w~ (again)
				if typeSpec, ok := node.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						for _, field := range structType.Fields.List {
							if selector, ok := field.Type.(*ast.SelectorExpr); ok {
								if xIdent, ok := selector.X.(*ast.Ident); ok {
									if xIdent.Name == componentPackageName && selector.Sel.Name == "Model" {
										meta.Usages = append(meta.Usages, pkg.Name)
										return false
									}
								}
							}
						}
					}
				}
				return true
			})
		}
	}

	return meta, nil
}
