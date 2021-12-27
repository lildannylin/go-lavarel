package middleware

import (
	celeritas "github.com/lildannylin/go-laverel"
	"github.com/lildannylin/go-laverel/cmd/cli/templates/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
