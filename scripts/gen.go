package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"
)

type TemplateData struct {
	Domain     string
	StructName string
}

func main() {
	domain := flag.String("domain", "", "Nombre del dominio (ej: order)")
	flag.Parse()

	if *domain == "" {
		log.Fatal("❌ Debes especificar --domain=nombre")
	}

	name := strings.ToLower(*domain)
	structName := strings.Title(name)

	data := TemplateData{
		Domain:     name,
		StructName: structName,
	}

	dir := "internal/" + name
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("❌ Error creando directorio: %v", err)
	}

	templates := map[string]string{
		"handler":    "templates/handler.tmpl",
		"service":    "templates/service.tmpl",
		"repository": "templates/repository.tmpl",
		"":           "templates/model.tmpl",
	}

	for kind, path := range templates {
		var fileName string
		if kind == "" {
			fileName = dir + "/" + name + ".go"
		} else {
			fileName = dir + "/" + name + "_" + kind + ".go"
		}
		generateFile(path, fileName, data)
	}
}

func generateFile(tmplPath, outPath string, data TemplateData) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Fatalf("❌ Error cargando template %s: %v", tmplPath, err)
	}

	f, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("❌ Error creando archivo %s: %v", outPath, err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, data); err != nil {
		log.Fatalf("❌ Error ejecutando template: %v", err)
	}

	log.Printf("✅ Generado: %s", outPath)
}
