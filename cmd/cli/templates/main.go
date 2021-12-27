package templates

import (
	celeritas "github.com/lildannylin/go-laverel"
	"github.com/lildannylin/go-laverel/cmd/cli/templates/data"
	"github.com/lildannylin/go-laverel/cmd/cli/templates/handlers"
	"github.com/lildannylin/go-laverel/cmd/cli/templates/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	//c := initApplication()
	//c.App.ListenAndServer()
}
