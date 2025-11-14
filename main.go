package main

import (
	"flag"

	"github.com/benyamin218118/todoService/infra/config"
	"github.com/benyamin218118/todoService/infra/db"
)

func main() {
	runMigrations := flag.Bool("run-migrations", false, "migrate the db and exit")
	config, err := config.Read(config.ENVConfigReader)
	if err != nil {
		panic(err)
	}

	// run the migrations and exit
	if runMigrations != nil && *runMigrations {
		err = db.RunMigrations(config.DBDSN)
		if err != nil {
			panic(err)
		}
		return
	}

	app := App{}
	app.Init(config)
	app.Run()
}
