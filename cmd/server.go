package cmd

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1talent/gotraining/internal/api"
	"github.com/1talent/gotraining/internal/api/router"
	"github.com/1talent/gotraining/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	probeFlag   string = "probe"
	migrateFlag string = "migrate"
	seedFlag    string = "seed"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the server",
	Long: `Starts the stateless RESTful JSON server

Requires configuration through ENV and
and a fully migrated PostgreSQL database.`,
	Run: func(cmd *cobra.Command, args []string) {

		runServer()
	},
}

func init() {
	// serverCmd.Flags().BoolP(probeFlag, "p", false, "Probe readiness before startup.")
	// serverCmd.Flags().BoolP(migrateFlag, "m", false, "Apply migrations before startup.")
	// serverCmd.Flags().BoolP(seedFlag, "s", false, "Seed fixtures into database before startup.")

	rootCmd.AddCommand(serverCmd)
}

func runServer() {
	config := config.DefaultServiceConfigFromEnv()

	zerolog.TimeFieldFormat = time.RFC3339Nano

	s := api.NewServer(config)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := s.InitDB(ctx); err != nil {
		cancel()
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}
	cancel()

	router.Init(s)

	go func() {
		if err := s.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info().Msg("Server closed")
			} else {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

}
