package templates

import "html/template"

const (
	IndexTemplateName = "index.gohtml"
	HomeTemplate      = "home"
	AboutTemplate     = "about"
)

func ParseTemplates() map[string]*template.Template {
	parsedTemplates := make(map[string]*template.Template)

	parsedTemplates[HomeTemplate] = template.Must(parseLayout().ParseFiles("templates/home.gohtml"))
	parsedTemplates[AboutTemplate] = template.Must(parseLayout().ParseFiles("templates/about.gohtml"))

	return parsedTemplates
}

func parseLayout() *template.Template {
	return template.Must(template.ParseFiles("templates/header.gohtml", "templates/index.gohtml"))
}
