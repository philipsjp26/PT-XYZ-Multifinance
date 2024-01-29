package migration

import (
	"context"
	"flag"
	"fmt"
	"go_playground/internal/infrastructure/config"
	"log"
	"os"

	"github.com/pressly/goose/v3"
)

var (
	flags   = flag.NewFlagSet("db:migrate", flag.ExitOnError)
	dir     = flags.String("dir", "database/migration", "directory with migration files")
	table   = flags.String("table", "db_migration", "migration table name")
	help    = flags.Bool("help", false, "print help")
	verbose = flags.Bool("verbose", false, "enable verbose mode")
)

func DatabaseMigration(cfg *config.Config) {
	flags.Usage = usage
	flags.Parse(os.Args[2:])

	if *verbose {
		goose.SetVerbose(true)
	}
	goose.SetTableName(*table)

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}
	ctx := context.Background()
	database := dsn(cfg)

	db, err := goose.OpenDBWithDriver(cfg.Database.Driver, database)
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "create":
		if err := goose.RunContext(ctx, "create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run create got :%v", err)
		}
		return
	case "fix":
		if err := goose.RunContext(ctx, "fix", nil, *dir); err != nil {
			log.Fatalf("goose run fix got :%v", err)
		}
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("error goose close db :%v", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	command := args[0]
	if err := goose.RunContext(ctx, command, db, *dir, arguments...); err != nil {
		log.Fatalf("error run command got : %v", err)
	}
}
func dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
}

func usage() {
	fmt.Println(usageCommands)
}
