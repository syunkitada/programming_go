package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	v1models "github.com/syunkitada/go-samples/generator-sample/pkg/v1/models"
	v2models "github.com/syunkitada/go-samples/generator-sample/pkg/v2/models"
)

type Handler struct {
	PkgPath string
	PkgName string
	Name    string
	Method  string
	Path    string
}

type Api struct {
	Version       string
	Handlers      []Handler
	HandlerModels []interface{}
}

type Spec struct {
	Apis []Api
}

func main() {
	spec := Spec{
		Apis: []Api{
			Api{
				Version: "v1",
				HandlerModels: []interface{}{
					v1models.GetUser{},
					v1models.PostUser{},
					v1models.GetRole{},
					v1models.GetProject{},
				},
			},
			Api{
				Version: "v2",
				HandlerModels: []interface{}{
					v2models.GetUser{},
				},
			},
		},
	}

	rootPath := "/home/owner/go/src/github.com/syunkitada/go-samples/generator-sample"

	for _, api := range spec.Apis {
		for _, handler := range api.HandlerModels {
			handlerType := reflect.TypeOf(handler)
			pkgPath := handlerType.PkgPath()
			splitedPath := strings.Split(pkgPath, "/")
			pkgName := splitedPath[len(splitedPath)-1]
			name := handlerType.Name()

			var method string
			if strings.HasPrefix(name, "Get") {
				method = "GET"
			} else if strings.HasPrefix(name, "Post") {
				method = "POST"
			}
			path := name[len(method):]
			path = strings.ToLower(path)

			api.Handlers = append(api.Handlers, Handler{
				PkgPath: pkgPath,
				PkgName: pkgName,
				Name:    handlerType.Name(),
				Method:  method,
				Path:    path,
			})
		}

		t := template.Must(template.ParseFiles("./templates/api.go.tmpl"))
		filePath := fmt.Sprintf(filepath.Join(rootPath, "pkg", api.Version, "api.go"))
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}

		if err := t.Execute(f, api); err != nil {
			log.Fatal(err)
		}

		cmd := "goimports -w " + filePath
		if out, err := exec.Command("sh", "-c", cmd).CombinedOutput(); err != nil {
			log.Fatalf("Failed cmd: %s, out=%s, err=%v", cmd, out, err)
		}
	}

}
