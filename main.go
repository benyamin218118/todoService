package main

import (
	"flag"

	"github.com/benyamin218118/todoService/infra/config"
	"github.com/benyamin218118/todoService/infra/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	runMigrations := flag.Bool("run-migrations", false, "migrate the db and exit")
	help := flag.Bool("help", false, "print help")
	config, err := config.Read(config.ENVConfigReader)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	if help != nil && *help {
		flag.PrintDefaults()
		return
	}
	// run the migrations and exit
	if runMigrations != nil && *runMigrations {
		err = db.RunMigrations(config)
		if err != nil {
			panic(err)
		}
		println("done")
		return
	}

	app := App{}
	app.Init(config)
	app.Run()
}
