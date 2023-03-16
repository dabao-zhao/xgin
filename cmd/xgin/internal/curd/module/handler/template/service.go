package template

// HandlerGen defines a template for handler
var HandlerGen = `
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
