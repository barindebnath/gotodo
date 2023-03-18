package template

import (
	"html/template"
	"time"
)

var Templates *template.Template

var tplFunctions = template.FuncMap{
	"monthYYYY":   monthYYYY,
	"ddMonthYYYY": ddMonthYYYY,
	"fullTime":    fullTime,
}

func init() {
	Templates = template.Must(template.New("newTpl").Funcs(tplFunctions).ParseGlob("template/*"))
}

func monthYYYY(t time.Time) string {
	return t.Format("January 2006")
}

func ddMonthYYYY(t time.Time) string {
	return t.Format("02 January 2006")
}

func fullTime(t time.Time) string {
	return t.Format("02 January 2006 15:04:05")
}
