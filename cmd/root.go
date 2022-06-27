package cmd

import (
	"fmt"
	"os"

	"github.com/1talent/gotraining/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: config.GetFormattedBuildArgs(),
	Use:     "app",
	Short:   config.ModuleName,
	Long: fmt.Sprintf(`%v

A stateless RESTful JSON service written in Go.
Requires configuration through ENV.`, config.ModuleName),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(`{{printf "\n" .Version}}`)
}
