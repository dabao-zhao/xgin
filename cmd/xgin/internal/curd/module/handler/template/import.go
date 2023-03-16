package template

const ImportTpl = `import (
	"context"
	"errors"

	"{{.servicesPackage}}"
	"{{.typesPackage}}"

	"github.com/jinzhu/copier"
	"github.com/op/go-logging"
	"github.com/gin-gonic/gin"
)
`
