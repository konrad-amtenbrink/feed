package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/konrad-amtenbrink/feed/internal/config"
	"github.com/konrad-amtenbrink/feed/internal/db"
	"github.com/konrad-amtenbrink/feed/internal/server"
	"github.com/konrad-amtenbrink/feed/internal/storage"
	"github.com/konrad-amtenbrink/feed/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultPort  = 4201
	portFlagName = "port"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Application Server",
		Run:   runStart,
	}
	cmd.Flags().Int(portFlagName, defaultPort, "port to listen on")

	return cmd
}

func runStart(cmd *cobra.Command, args []string) {
	port, err := cmd.Flags().GetInt(portFlagName)
	if err != nil {
		log.WithError(err).Fatal("cmd: failed to get port flag")
		os.Exit(1)
	}

	ctx := cmd.Context()

	cfg := config.MustParseConfig()
	logger.MustInit(cfg.Logging)

	db, _, err := db.NewDatabase(ctx, cfg.DB)
	if err != nil {
		log.WithError(err).Fatalf("failed to connect to database")
	}

	storage, _, err := storage.NewStorage(ctx, cfg.AWS)
	if err != nil {
		log.WithError(err).Fatal("failed to create storage")
	}

	srv := server.New(db, storage)
	go func() {
		if err = srv.Run(port); err != http.ErrServerClosed {
			log.WithError(err).Errorf("failed to run server")
			os.Exit(1)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-signalChan
	log.Warn("signal: shutting down")

	if err := srv.Shutdown(); err != nil {
		log.WithError(err).Error("failed to close server")
		os.Exit(2)
	} else {
		log.Warn("closed server gracefully")
	}

	os.Exit(0)
}
