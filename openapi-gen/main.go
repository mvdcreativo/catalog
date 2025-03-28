package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Method      string
	Path        string
	RequestType string
	Tag         string
	Summary     string
}

type OpenAPI struct {
	OpenAPI string              `yaml:"openapi"`
	Info    Info                `yaml:"info"`
	Paths   map[string]PathItem `yaml:"paths"`
}

type Info struct {
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

type PathItem map[string]MethodItem

type MethodItem struct {
	Tags        []string            `yaml:"tags"`
	Summary     string              `yaml:"summary"`
	RequestBody *RequestBody        `yaml:"requestBody,omitempty"`
	Responses   map[string]Response `yaml:"responses"`
}

type RequestBody struct {
	Required bool                       `yaml:"required"`
	Content  map[string]MediaTypeObject `yaml:"content"`
}

type Response struct {
	Description string                     `yaml:"description"`
	Content     map[string]MediaTypeObject `yaml:"content"`
}

type MediaTypeObject struct {
	Schema map[string]string `yaml:"schema"`
}

var endpoints []Endpoint
var groupPrefix string

func main() {
	_ = filepath.Walk("../internal/routes", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			parseFile(path)
		}
		return nil
	})

	openapi := buildOpenAPI(endpoints)

	data, err := yaml.Marshal(openapi)
	if err != nil {
		log.Fatalf("❌ Error generating YAML: %v", err)
	}
	os.WriteFile("../docs/openapi.yaml", data, 0644)
	fmt.Println("✅ openapi.yaml generado correctamente")
}

func parseFile(path string) {
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, path, nil, parser.AllErrors)
	if err != nil {
		log.Printf("Error parsing file %s: %v", path, err)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if ident, ok := sel.X.(*ast.Ident); ok {
			if ident.Name == "group" && isHTTPMethod(sel.Sel.Name) {
				if len(call.Args) >= 2 {
					pathLit, _ := call.Args[0].(*ast.BasicLit)
					fullPath := groupPrefix + strings.Trim(pathLit.Value, `"`)
					requestType := extractRequestType(call.Args)
					endpoints = append(endpoints, Endpoint{
						Method:      strings.ToLower(sel.Sel.Name),
						Path:        fullPath,
						RequestType: requestType,
						Tag:         strings.TrimPrefix(groupPrefix, "/"),
						Summary:     sel.Sel.Name,
					})
				}
			}
		} else if sel.Sel.Name == "Group" {
			if len(call.Args) > 0 {
				if prefix, ok := call.Args[0].(*ast.BasicLit); ok {
					groupPrefix = strings.Trim(prefix.Value, `"`)
				}
			}
		}

		return true
	})
}

func extractRequestType(args []ast.Expr) string {
	for _, arg := range args {
		if call, ok := arg.(*ast.CallExpr); ok {
			if fun, ok := call.Fun.(*ast.IndexExpr); ok {
				if ident, ok := fun.Index.(*ast.SelectorExpr); ok {
					return ident.X.(*ast.Ident).Name + "." + ident.Sel.Name
				}
			}
		}
	}
	return ""
}

func isHTTPMethod(name string) bool {
	return name == "GET" || name == "POST" || name == "PUT" || name == "DELETE"
}

func buildOpenAPI(endpoints []Endpoint) OpenAPI {
	paths := make(map[string]PathItem)

	for _, e := range endpoints {
		methodItem := MethodItem{
			Tags:    []string{e.Tag},
			Summary: e.Summary,
			Responses: map[string]Response{
				"200": {
					Description: "Success",
					Content: map[string]MediaTypeObject{
						"application/json": {
							Schema: map[string]string{
								"$ref": "#/components/schemas/" + getSimpleName(e.RequestType),
							},
						},
					},
				},
			},
		}

		if e.Method == "post" || e.Method == "put" {
			methodItem.RequestBody = &RequestBody{
				Required: true,
				Content: map[string]MediaTypeObject{
					"application/json": {
						Schema: map[string]string{
							"$ref": "#/components/schemas/" + getSimpleName(e.RequestType),
						},
					},
				},
			}
		}

		if paths[e.Path] == nil {
			paths[e.Path] = PathItem{}
		}
		paths[e.Path][e.Method] = methodItem
	}

	return OpenAPI{
		OpenAPI: "3.1.0",
		Info: Info{
			Title:   "Catalog API",
			Version: "1.0",
		},
		Paths: paths,
	}
}

func getSimpleName(qualified string) string {
	parts := strings.Split(qualified, ".")
	return parts[len(parts)-1]
}
