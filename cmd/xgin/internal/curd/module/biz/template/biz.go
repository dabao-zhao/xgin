package template

// BizGen defines a template for biz
var BizGen = `
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
