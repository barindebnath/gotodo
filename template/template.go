package template

import (
	"html/template"
	"time"
)

var Templates *template.Template

var tplFunctions = template.FuncMap{
	"monthYYYY": monthYYYY,
}

func init() {
	Templates = template.Must(template.New("newTpl").Funcs(tplFunctions).ParseGlob("template/*"))
}

func monthYYYY(t time.Time) string {
	return t.Format("January 2006")
}
