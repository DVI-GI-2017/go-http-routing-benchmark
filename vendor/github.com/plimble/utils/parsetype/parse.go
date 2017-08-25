package parsetype

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"reflect"
	"strings"
)

type Parser struct {
	Packages     map[string]*Package
	Types        map[string]*Type
	useComposite []*Type
}

type Package struct {
	Package *ast.Package
	Files   map[string]*File
}

type File struct {
	File      *ast.File
	Imports   map[string]string
	Alias     map[string]string
	IsBuiltIn bool
}

type Type struct {
	Name        string
	Type        string
	Doc         *ast.CommentGroup
	Comment     *ast.CommentGroup
	FromPackage string
	Properties  map[string]*Type
	ArrayType   *Type
	MapKey      string
	MapType     *Type
	Tags        reflect.StructTag
	RefType     *Type
	RefPackage  string
	RefTypeName string
	IsPointer   bool
	composite   []string
}

func NewParser() *Parser {
	return &Parser{
		Packages: make(map[string]*Package),
		Types:    make(map[string]*Type),
	}
}

func (p *Parser) MergeComposite() {
	for i := 0; i < len(p.useComposite); i++ {
		for j := 0; j < len(p.useComposite[i].composite); j++ {
			if p.Types[p.useComposite[i].composite[j]] != nil {
				for key, val := range p.Types[p.useComposite[i].composite[j]].Properties {
					p.useComposite[i].Properties[key] = val
				}
			}
		}
	}
}

func (p *Parser) ParseDir(dir string) {
	if _, ok := p.Packages[dir]; !ok {
		fset := token.NewFileSet()
		packages, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
		if err != nil {
			// logrus.Warnln(err)
			return
		}

		p.parse(dir, packages)
	}
}

func (p *Parser) trimGoPath(s string) string {
	paths := os.Getenv("GOPATH")
	goPaths := strings.Split(paths, ":")
	for i := 0; i < len(goPaths); i++ {
		if strings.HasPrefix(s, goPaths[i]) {
			return strings.TrimPrefix(s, path.Join(goPaths[i], "src")+"/")
		}
	}

	return s
}

func (p *Parser) ParseFromGoPath(dir string) {
	paths := os.Getenv("GOPATH")
	goPaths := strings.Split(paths, ":")
	for i := 0; i < len(goPaths); i++ {
		p.ParseDir(path.Join(goPaths[i], "src", dir))
	}
}

func (p *Parser) parse(dir string, packages map[string]*ast.Package) {
	for _, astPackage := range packages {
		if p.Packages[dir] != nil {
			return
		}

		if strings.Contains(astPackage.Name, "_test") {
			continue
		}

		pack := &Package{
			Files:   make(map[string]*File),
			Package: astPackage,
		}

		p.Packages[dir] = pack

		for _, astFile := range astPackage.Files {
			if strings.Contains(astFile.Name.String(), "_test") {
				continue
			}

			file := &File{
				File:    astFile,
				Imports: make(map[string]string),
				Alias:   make(map[string]string),
			}

			pack.Files[astFile.Name.String()] = file

			for _, astDeclaration := range astFile.Decls {
				if generalDeclaration, ok := astDeclaration.(*ast.GenDecl); ok {
					for _, astSpec := range generalDeclaration.Specs {
						if typeSpec, ok := astSpec.(*ast.ImportSpec); ok {
							importPath := strings.Replace(typeSpec.Path.Value, "\"", "", -1)
							packName := importPath
							packNameIndex := strings.LastIndex(importPath, "/")
							if packNameIndex > 0 {
								packName = importPath[packNameIndex+1:]
							}

							if typeSpec.Name != nil {
								file.Alias[typeSpec.Name.String()] = packName
							}

							file.Imports[packName] = importPath
						}

						if typeSpec, ok := astSpec.(*ast.TypeSpec); ok {
							p.parseTypeSpec(p.trimGoPath(dir), file, typeSpec, generalDeclaration.Doc)
						}
					}

				}
			}
		}
	}
}

func (p *Parser) parseTypeSpec(packname string, file *File, astType *ast.TypeSpec, doc *ast.CommentGroup) *Type {
	name := packname + "." + astType.Name.String()
	if _, ok := p.Types[name]; ok && p.Types[name].Type != "" {
		return p.Types[name]
	}

	var t *Type
	if _, ok := p.Types[name]; ok && p.Types[name].Type == "" {
		t = p.Types[name]
	} else {
		t = &Type{}
		p.Types[name] = t
	}

	t.Name = name
	t.Doc = doc
	p.parseType(packname, file, t, astType.Type)

	return p.Types[name]
}

func (p *Parser) parseType(packname string, file *File, t *Type, astType ast.Expr) {
	switch expr := astType.(type) {
	case *ast.Ident:
		if basicTypes[expr.String()] {
			t.Type = expr.String()
		} else {
			// composite in pkg
			// t.composite = t.RefPackage + "." + t.RefTypeName
			t.Type = "ref"
			t.RefPackage = packname
			t.RefTypeName = expr.String()
			if _, ok := p.Types[t.RefPackage+"."+t.RefTypeName]; !ok {
				p.Types[t.RefPackage+"."+t.RefTypeName] = &Type{}
			}

			t.RefType = p.Types[t.RefPackage+"."+t.RefTypeName]
		}
	case *ast.ArrayType:
		t.Type = "array"
		t.ArrayType = &Type{}
		p.parseType(packname, file, t.ArrayType, expr.Elt)
	case *ast.MapType:
		t.Type = "map"
		t.MapType = &Type{}
		t.MapKey = fmt.Sprint(expr.Key)
		p.parseType(packname, file, t.MapType, expr.Value)
	case *ast.SelectorExpr:
		t.Type = "ref"
		t.RefTypeName = expr.Sel.Name
		t.RefPackage = fmt.Sprint(expr.X)

		if t.RefPackage == "time" {
			t.Type = "time"
			break
		}

		//check is alias
		for key, val := range file.Alias {
			if key == t.RefPackage {
				t.RefPackage = val
			}
		}
		p.ParseFromGoPath(file.Imports[t.RefPackage])
		if ref, ok := p.Types[file.Imports[t.RefPackage]+"."+t.RefTypeName]; ok {
			t.RefType = ref
		}
	case *ast.StarExpr:
		t.IsPointer = true
		p.parseType(packname, file, t, expr.X)
	case *ast.StructType:
		t.Type = "struct"
		t.Properties = make(map[string]*Type)
		for j := 0; j < len(expr.Fields.List); j++ {
			field := expr.Fields.List[j]

			switch {
			case len(field.Names) == 0:
				propName := ""
				if astSelectorExpr, ok := field.Type.(*ast.SelectorExpr); ok {
					fieldPackName := astSelectorExpr.X.(*ast.Ident).Name
					//check is alias
					for key, val := range file.Alias {
						if key == fieldPackName {
							fieldPackName = val
						}
					}
					propName = fieldPackName + "." + strings.TrimPrefix(astSelectorExpr.Sel.Name, "*")
				} else if astTypeIdent, ok := field.Type.(*ast.Ident); ok {
					propName = packname + "." + astTypeIdent.Name
				} else if astStarExpr, ok := field.Type.(*ast.StarExpr); ok {
					if astSelectorExpr, ok := astStarExpr.X.(*ast.SelectorExpr); ok {
						fieldPackName := astSelectorExpr.X.(*ast.Ident).Name
						//check is alias
						for key, val := range file.Alias {
							if key == fieldPackName {
								fieldPackName = val
							}
						}
						propName = fieldPackName + "." + strings.TrimPrefix(astSelectorExpr.Sel.Name, "*")
					}

					if astIdent, ok := astStarExpr.X.(*ast.Ident); ok {
						propName = packname + "." + astIdent.Name
					}
				} else {
					panic(fmt.Errorf("Something goes wrong: %#v", field.Type))
				}

				p.parseProp(packname, file, t, field, propName, true)
			case len(field.Names) > 0:
				for i := 0; i < len(field.Names); i++ {
					p.parseProp(packname, file, t, field, field.Names[i].String(), false)
				}
			default:
				p.parseProp(packname, file, t, field, field.Names[0].String(), false)
			}
		}
	default:
		t.Type = "unknown"
	}
}

func (p *Parser) parseProp(packname string, file *File, t *Type, astField *ast.Field, propName string, isComposite bool) {
	fieldType := &Type{}
	fieldType.Name = propName
	fieldType.Doc = astField.Doc
	fieldType.Comment = astField.Comment
	if astField.Tag != nil {
		fieldType.Tags = reflect.StructTag(strings.Trim(astField.Tag.Value, "`"))
	}

	// composite out side pkg
	if isComposite {
		t.composite = append(t.composite, fieldType.Name)
		p.useComposite = append(p.useComposite, t)
	} else {
		p.parseType(packname, file, fieldType, astField.Type)
		t.Properties[fieldType.Name] = fieldType
	}

}

var basicTypes = map[string]bool{
	"bool":       true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"float32":    true,
	"float64":    true,
	"string":     true,
	"complex64":  true,
	"complex128": true,
	"byte":       true,
	"rune":       true,
	"uintptr":    true,
	"error":      true,
	"Time":       true,
}
