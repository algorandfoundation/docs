package main

import (
	_ "embed"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//go:embed template.tmpl
var configTemplate string
var re = regexp.MustCompile(`\r?\n`)

// resolveLocalTemplate searches for the localTemplate.go file in the given path.
func resolveLocalTemplate(pathStr string) (string, error) {
	var pathTolocalTemplate string
	err := filepath.Walk(pathStr, func(path string, info fs.FileInfo, err error) error {
		// looking for a specific file.
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, "localTemplate.go") {
			pathTolocalTemplate = path
		}
		return nil
	})
	return pathTolocalTemplate, err
}

// DocParts are used to store data for each configuration
type DocParts struct {
	Name        string
	Description string
	Default     string
	Type        string
}

// TemplateFields is passed to the go text template.
type TemplateFields struct {
	NodeDocs []DocParts
}

// parseDefault extracts the most recent default value from the tag.
func parseDefault(tag string) string {
	unwrap := func(str string, r byte) string {
		if str[0] == r {
			str = str[1:]
		}
		if len(str) > 0 && str[len(str)-1] == r {
			str = str[:len(str)-1]
		}
		return str
	}
	tag = unwrap(tag, '`')
	versions := strings.Split(tag, " ")
	version := versions[len(versions)-1]
	value := strings.Split(version, ":")[1]
	value = unwrap(value, '"')
	return value
}

// parseType converts the parameter type into a string.
func parseType(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.Ident:
		return fmt.Sprintf("%s", t)
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", t.X, t.Sel)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", t.Key, t.Value)
	default:
		fmt.Fprintf(os.Stderr, "unknown configuration type: %s\n", t)
		return fmt.Sprintf("%s", t)
	}
}

// parseFile extracts the doc parts from a go file.
func parseFile(filePath string) ([]DocParts, error) {
	docs := make([]DocParts, 0)
	src, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("parseFile: unable to read source: %w", err)
	}

	fset := token.NewFileSet()
	parsed, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parseFile: unable to parse file: %w", err)
	}

	ast.Inspect(parsed, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// skip functions
			return false
		case *ast.TypeSpec:
			// skip everything but the "Local" type
			return "Local" == x.Name.Name
		case *ast.Field:
			missingDoc := false
			if len(x.Names) != 1 {
				missingDoc = true
			}
			if x.Doc.Text() == "" {
				missingDoc = true
			}
			if missingDoc {
				for _, name := range x.Names {
					fmt.Fprintf(os.Stderr, "missing comment for type: %s\n", name)
				}
				return true
			}

			// Grab common parts
			doc := DocParts{
				Name:        x.Names[0].Name,
				Description: re.ReplaceAllString(strings.TrimSpace(x.Doc.Text()), "<br>"),
				Default:     parseDefault(x.Tag.Value),
				Type:        parseType(x.Type),
			}

			docs = append(docs, doc)
		}

		return true
	})

	return docs, nil
}

func main() {
	var pathStr string
	flag.StringVar(&pathStr, "path", "", "Path to go-algorand, used to resolve localTemplate.go which is used for the parameter table.")
	flag.Parse()

	if pathStr == "" {
		fmt.Fprintf(os.Stderr, "Must provide path to go-algorand with -path.")
		os.Exit(1)
	}

	pathToLocalTemplate, err := resolveLocalTemplate(pathStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred searching '%s': %s", pathStr, err)
		os.Exit(1)
	}

	docs, err := parseFile(pathToLocalTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred parsing file '%s': %s", pathToLocalTemplate, err)
		os.Exit(1)
	}

	ut, err := template.New("configDoc").Parse(configTemplate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred parsing the template: %s", err)
		os.Exit(1)
	}

	err = ut.Execute(os.Stdout, TemplateFields{NodeDocs: docs})
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while executing the template: %s", err)
		os.Exit(1)
	}
}
