package template

const ImportTpl = `import (
	"context"
	"errors"

	"{{.bizPackage}}"
	"{{.typesPackage}}"

	"github.com/jinzhu/copier"
	"github.com/op/go-logging"
)
`
