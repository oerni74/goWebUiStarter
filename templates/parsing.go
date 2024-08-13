package templates

import "html/template"

func ParseTemplates() map[string]*template.Template {
	parsedTemplates := make(map[string]*template.Template)

	parsedTemplates["home"] = template.Must(parseLayout().ParseFiles("templates/home.gohtml"))
	parsedTemplates["about"] = template.Must(parseLayout().ParseFiles("templates/about.gohtml"))

	return parsedTemplates
}

func parseLayout() *template.Template {
	return template.Must(template.ParseFiles("templates/header.gohtml", "templates/index.gohtml"))
}
