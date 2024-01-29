package cmd

import (
	"go_playground/cmd/http"
	"go_playground/internal/infrastructure/config"
	"go_playground/pkg/migration"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Start() {
	rootCommand := &cobra.Command{}

	cmd := []*cobra.Command{
		{
			Use:   "serve",
			Short: "Run server",
			Run: func(cmd *cobra.Command, args []string) {
				http.Runner()
			},
		},
		{
			Use:   "db:migrate",
			Short: "Database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.DatabaseMigration(config.Configuration())
			},
		},
		{
			Use:   "create",
			Short: "Database create",
			Run: func(cmd *cobra.Command, args []string) {
				migration.DatabaseMigration(config.Configuration())
			},
		},
	}
	rootCommand.AddCommand(cmd...)
	if err := rootCommand.Execute(); err != nil {
		log.Fatalf("error run command using cobra : %v", err)
		os.Exit(1)
	}
}
