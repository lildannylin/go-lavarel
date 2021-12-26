package myapp

import celeritas "github.com/lildannylin/go-laverel"

type application struct {
	App *celeritas.Celeritas
}

func main() {
	c := initApplication()
	c.App.ListenAndServer()
}
