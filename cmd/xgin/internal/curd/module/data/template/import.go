package template

const ImportTpl = `import (
	"context"

	"{{.bizPackage}}"

	"gorm.io/gorm"
	"github.com/op/go-logging"
)
`
