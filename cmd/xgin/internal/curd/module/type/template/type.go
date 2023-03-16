package template

// TypeTpl defines a template for type in type.
const TypeTpl = `
type (
	Create{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Create{{.upperStartCamelObject}}Response struct {
		{{.fieldsTagOnlyJson}}
	}

	Update{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Update{{.upperStartCamelObject}}Response struct {
		{{.fieldsTagOnlyJson}}
	}

	Delete{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	Delete{{.upperStartCamelObject}}Response struct {
	}

	FindOne{{.upperStartCamelObject}}Request struct {
		{{.fields}}
	}
	FindOne{{.upperStartCamelObject}}Response struct {
		{{.fieldsTagOnlyJson}}
	}
)
`

// TypeGen defines a template for type
var TypeGen = `
package {{.pkg}}
{{.imports}}
{{.types}}
`
