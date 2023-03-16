package template

// ServiceGen defines a template for service
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
