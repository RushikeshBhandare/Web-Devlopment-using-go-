package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "Check.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}

}
