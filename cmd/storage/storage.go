package storage

import (
	"os"

	"github.com/konrad-amtenbrink/feed/config"
	"github.com/konrad-amtenbrink/feed/logger"
	"github.com/konrad-amtenbrink/feed/storage"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultFile  = ""
	fileFlagName = "file"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "storage",
		Short: "Storage upload",
		Run:   runUpload,
	}
	cmd.Flags().String(fileFlagName, defaultFile, "port to listen on")

	return cmd
}

func runUpload(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString(fileFlagName)
	if err != nil || file == "" {
		log.WithError(err).Fatal("cmd: failed to get file name")
		os.Exit(1)
	}

	ctx := cmd.Context()

	cfg := config.MustParseConfig()
	logger.MustInit(cfg.Logging)

	storage, _, err := storage.NewStorage(ctx, cfg.AWS)
	if err != nil {
		log.WithError(err).Fatal("cmd: failed to create storage")
		os.Exit(1)
	}

	err = storage.Upload(file)
	if err != nil {
		log.WithError(err).Fatal("cmd: failed to upload file")
		os.Exit(1)
	}

	os.Exit(0)
}
