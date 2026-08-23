package main

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// nolint: gochecknoglobals
var (
	buildVersion   string
	buildDate      string
	buildCommit    string
	buildGoVersion string
)

var rootCmd = &cobra.Command{
	Use:   "kratos",
	Short: "Kratos: An elegant toolkit for Go microservices.",
	Long:  `Kratos: An elegant toolkit for Go microservices.`,
}

func init() {
	if buildVersion == "" {
		buildVersion = "N/A"
	}

	if buildDate == "" {
		buildDate = "N/A"
	}

	if buildCommit == "" {
		buildCommit = "N/A"
	}

	buildGoVersion = "N/A"
	bi, ok := debug.ReadBuildInfo()
	if ok && bi != nil && bi.GoVersion != "" {
		buildGoVersion = bi.GoVersion
	}

	rootCmd.Version = fmt.Sprintf("droid has version %s built with go %s from %s on %s", buildVersion, buildGoVersion, buildCommit, buildDate)
	// TODO: register commands
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
