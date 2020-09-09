// +build !simple

package command

import (
	"github.com/micro/cli/v2"
	"github.com/owncloud/ocis-thumbnails/pkg/command"
	"github.com/owncloud/ocis-thumbnails/pkg/flagset"
	"github.com/owncloud/ocis/pkg/config"
	"github.com/owncloud/ocis/pkg/register"

	svcconfig "github.com/owncloud/ocis-thumbnails/pkg/config"
)

// ThumbnailsCommand is the entrypoint for the thumbnails command.
func ThumbnailsCommand(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "thumbnails",
		Usage:    "Start thumbnails server",
		Category: "Extensions",
		Flags:    flagset.ServerWithConfig(cfg.Thumbnails),
		Action: func(c *cli.Context) error {
			thumbnailsCommand := command.Server(configureThumbnails(cfg))

			if err := thumbnailsCommand.Before(c); err != nil {
				return err
			}

			return cli.HandleAction(thumbnailsCommand.Action, c)
		},
	}
}

func configureThumbnails(cfg *config.Config) *svcconfig.Config {
	cfg.Thumbnails.Log.Level = cfg.Log.Level
	cfg.Thumbnails.Log.Pretty = cfg.Log.Pretty
	cfg.Thumbnails.Log.Color = cfg.Log.Color

	if cfg.Tracing.Enabled {
		cfg.Thumbnails.Tracing.Enabled = cfg.Tracing.Enabled
		cfg.Thumbnails.Tracing.Type = cfg.Tracing.Type
		cfg.Thumbnails.Tracing.Endpoint = cfg.Tracing.Endpoint
		cfg.Thumbnails.Tracing.Collector = cfg.Tracing.Collector
		cfg.Thumbnails.Tracing.Service = cfg.Tracing.Service
	}

	return cfg.Thumbnails
}

func init() {
	register.AddCommand(ThumbnailsCommand)
}
