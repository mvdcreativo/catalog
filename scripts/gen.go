package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type TemplateData struct {
	Domain     string
	StructName string
	Module     string
}

func main() {
	domain := flag.String("domain", "", "Nombre del dominio (ej: order)")
	flag.Parse()

	if *domain == "" {
		log.Fatal("❌ Debes especificar --domain=nombre")
	}

	module := getModuleName()
	if module == "" {
		log.Fatal("❌ No se pudo detectar el nombre del módulo desde go.mod")
	}

	name := strings.ToLower(*domain)
	structName := strings.Title(name)

	data := TemplateData{
		Domain:     name,
		StructName: structName,
		Module:     module,
	}

	createInternalFiles(data)
	createRoutesFile(data)
	updateRoutesGo(data)
	printBootstrapInstructions(data)
}

func getModuleName() string {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "module ") {
		return strings.TrimSpace(strings.TrimPrefix(lines[0], "module"))
	}
	return ""
}

func createInternalFiles(data TemplateData) {
	dir := "internal/" + data.Domain
	_ = os.MkdirAll(dir, os.ModePerm)

	templates := map[string]string{
		"handler":    "templates/handler.tmpl",
		"service":    "templates/service.tmpl",
		"repository": "templates/repository.tmpl",
		"":           "templates/model.tmpl",
	}

	for kind, path := range templates {
		var fileName string
		if kind == "" {
			fileName = dir + "/" + data.Domain + ".go"
		} else {
			fileName = dir + "/" + data.Domain + "_" + kind + ".go"
		}
		renderTemplate(path, fileName, data)
	}
}

func createRoutesFile(data TemplateData) {
	path := fmt.Sprintf("internal/routes/%s_routes.go", data.Domain)
	content := fmt.Sprintf(`package routes

import (
	"github.com/gin-gonic/gin"
	"%s/internal/%s"
)

func %sRoutes(rg *gin.RouterGroup, h *%s.%sHandler) {
	group := rg.Group("/%ss")
	group.GET("", h.FindAll)
	group.POST("", h.Insert)
	group.GET("/:id", h.FindByID)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
`, data.Module, data.Domain, data.StructName, data.Domain, data.StructName, data.Domain)

	_ = os.WriteFile(path, []byte(content), 0644)
	log.Printf("✅ %s_routes.go creado", data.Domain)
}

func updateRoutesGo(data TemplateData) {
	path := "internal/routes/routes.go"
	content, _ := os.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	var buffer []string

	insertedImport := false
	insertedParam := false
	insertedCall := false

	for _, line := range lines {
		if line == "import (" && !insertedImport {
			buffer = append(buffer, line)
			buffer = append(buffer, fmt.Sprintf(`	"%s/internal/%s"`, data.Module, data.Domain))
			insertedImport = true
			continue
		}

		if strings.Contains(line, "SetupRoutes(") && !insertedParam {
			buffer = append(buffer, line)
			buffer = append(buffer, fmt.Sprintf("	%sHandler *%s.%sHandler,", data.Domain, data.Domain, data.StructName))
			insertedParam = true
			continue
		}

		if strings.Contains(line, "api := r.Group") && !insertedCall {
			buffer = append(buffer, line)
			buffer = append(buffer, fmt.Sprintf("	%sRoutes(api, %sHandler)", data.StructName, data.Domain))
			insertedCall = true
			continue
		}

		buffer = append(buffer, line)
	}

	_ = os.WriteFile(path, []byte(strings.Join(buffer, "\n")), 0644)
	log.Println("✅ routes.go actualizado")
}

func printBootstrapInstructions(data TemplateData) {
	fmt.Println("\n🔧 Debes actualizar manualmente tu archivo bootstrap.go con lo siguiente:\n")

	fmt.Println("🔹 1. Agregar import:")
	fmt.Printf("\t\"%s/internal/%s\"\n\n", data.Module, data.Domain)

	fmt.Println("🔹 2. Agregar en el struct App:")
	fmt.Printf("\t%sHandler *%s.%sHandler\n\n", data.StructName, data.Domain, data.StructName)

	fmt.Println("🔹 3. Agregar en el slice modules:")
	fmt.Printf("\tregister%s,\n\n", data.StructName)

	fmt.Println("🔹 4. Agregar como parámetro en SetupRoutes:")
	fmt.Printf("\tapp.%sHandler,\n\n", data.StructName)

	fmt.Println("🔹 5. Agregar esta función al final del archivo:")
	fmt.Printf(`
func register%s(app *App) {
	repo := %s.New%sRepository(app.MongoClient, app.Config.DbName, "%ss")
	service := %s.New%sService(repo)
	app.%sHandler = %s.New%sHandler(service)
}
`, data.StructName, data.Domain, data.StructName, data.Domain,
		data.Domain, data.StructName, data.StructName, data.Domain, data.StructName)
}

func renderTemplate(tmplPath, outPath string, data TemplateData) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("❌ Error cargando plantilla %s: %v", tmplPath, err)
	}

	f, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("❌ Error creando archivo %s: %v", outPath, err)
	}
	defer f.Close()

	_ = tmpl.Execute(f, data)
	log.Printf("✅ Generado: %s", outPath)
}
