package command

import (
	"os"

	"github.com/maurana/nuswantara/core/server"
	"github.com/maurana/nuswantara/core/migrate"
	"github.com/maurana/nuswantara/core/logger"
	cli "github.com/urfave/cli/v2"
)

func NuswantaraServer() error {
	app := cli.NewApp()
	app.Name = "Nuswantara Framework"
	app.Description = "Elegant Framework For Golang"

	app.Commands = []*cli.Command{
		{
			Name:        "migrations",
			Description: " migration up database",
			Action: func(c *cli.Context) error {
				err := migrate.Up()
				if err != nil {
					return err
				}
				logger.Log().Info().Msg("migrations successfully")
				return nil
			},
		},
		{
			Name:        "rollbacks",
			Description: "migration down database",
			Action: func(c *cli.Context) error {
				err := migrate.Down()
				if err != nil {
					return err
				}
				logger.Log().Info().Msg("rollbacks successfully")
				return nil
			},
		},
		{
			Name:        "steps",
			Description: "migration steps, it will migrate up if n > 0, and down if n < 0",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "n"},
			},
			Action: func(c *cli.Context) error {
				err := migrate.Steps(c.Int("n"))
				if err != nil {
					return err
				}
				logger.Log().Info().Msgf("steps %d successfully", c.Int("n"))
				return nil
			},
		},
		{
			Name:        "drop",
			Description: "migration drop deletes database",
			Action: func(c *cli.Context) error {
				err := migrate.Drop()
				if err != nil {
					return err
				}
				logger.Log().Info().Msg("drop successfully")
				return nil
			},
		},
		{
			Name:        "start",
			Description: "start the server",
			Action: func(c *cli.Context) error {
				return server.Startup()
			},
		},
		{
			Name:        "launch",
			Description: "migration up database with start the server",
			Action: func(c *cli.Context) error {
				err := migrate.Up()
				if err != nil {
					return err
				}
				return server.Startup()
			},
		},
	}

	return app.Run(os.Args)
}