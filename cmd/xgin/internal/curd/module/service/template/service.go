package template

// ServiceGen defines a template for biz
var ServiceGen = `
package {{.pkg}}
{{.imports}}
{{.types}}
{{.new}}
{{.create}}
{{.update}}
{{.delete}}
{{.find}}
{{.extraMethod}}
`
