package templatego

import (
	"html/template"
	"log"
	"os"
)

var TemplateMap map[string]*template.Template

func init() {

	GOPATH := os.Getenv("GOPATH")

	var base = GOPATH + "/src/github.com/itsmeadi/cart/frontend/vegefoods/"

	TemplateMap = make(map[string]*template.Template)
	myTemplates := make(map[string]string)
	myTemplates["index"] = base + "products.html"
	myTemplates["product_detail"] = base + "product-single.html"
	myTemplates["login"] = base + "login.html"
	myTemplates["cart"] = base + "cart.html"

	for k, v := range myTemplates {
		//t := template.Must(template.New("question").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`))
		t := template.Must(template.ParseFiles(v))
		TemplateMap[k] = t

	}

	//
	//for k, v := range myTemplates {
	//	t := template.Must(template.New(k).ParseFiles(v))
	//	TemplateMap[k] = t
	//}

	log.Print("Templated init complete")
}

// templateFile defines the contents of a template to be stored in a file, for testing.
type templateFile struct {
	name     string
	contents string
}
