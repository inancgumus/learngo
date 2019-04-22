package main

import "html/template"

var tmpl *template.Template

func init() {
	tmpl = template.Must(
		template.New("list.tmpl").
			ParseFiles("list.tmpl"))
}
