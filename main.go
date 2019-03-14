package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	b, err := Asset("tpl/page.html")
	if err != nil {
		log.Fatalf("unable to get sset: %v", err)
	}

	t, err := template.New("base").Parse(string(b))
	if err != nil {
		log.Fatalf("unable to parse tpl file: %v", err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("unable to execute tpl: %v", err)
	}
	// t, err := template.ParseFiles("tpl/page.html")
	// if err != nil {
	// 	log.Fatalf("unable to parse tpl file: %v", err)
	// }

	// err = t.ExecuteTemplate(os.Stdout, "base", nil)
	// if err != nil {
	// 	log.Fatalf("unable to execute tpl: %v", err)
	// }
}
