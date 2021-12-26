package templates

import (
	celeritas "github.com/lildannylin/go-laverel"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "templates"

	app := &application{
		App: cel,
	}

	return app
}
