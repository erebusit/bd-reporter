package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/erebusit/bd-reporter/cmd/model"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	fBdToken       string
	fBdProject     string
	fBdVersion     string
	fBdEnvironment string
	isCircleCi     bool
)

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.Flags().StringVarP(&fBdProject, "project", "p", os.Getenv("CIRCLE_PROJECT_REPONAME"), "Project name. Defaults to CIRCLE_PROJECT_REPONAME.")
	reportCmd.Flags().StringVarP(&fBdToken, "bd-token", "b", os.Getenv("BD_TOKEN"), "Build Dash API token. Will be fetched from BD_TOKEN by default.")
	reportCmd.Flags().StringVarP(&fBdVersion, "version", "v", os.Getenv("CIRCLE_BUILD_NUM"), "Version string. Will be fetched from CIRCLE_BUILD_NUM by default.")
	reportCmd.Flags().StringVarP(&fBdEnvironment, "environment", "e", "production", "Environment name. Set to 'production' unless specified.")

	v := os.Getenv("CIRCLECI")
	if strings.Compare(v, "true") != 0 {
		isCircleCi = false
	} else {
		isCircleCi = true
	}
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "reports the deployment status",
	Run: func(cmd *cobra.Command, args []string) {
		if !isCircleCi {
			log.Fatalf("expected CIRCLECI to be set to true, but was %t", isCircleCi)
		}
		b, err := json.Marshal(&model.DeploymentReport{
			Status:      model.Error,
			Project:     fBdProject,
			Version:     fBdVersion,
			Environment: fBdEnvironment,
		})
		if err != nil {
			log.Fatal(err)
		}
		rsp, err := http.Post("https://deployments.eu.quickci.io/deployments", "application/json", bytes.NewBuffer(b))
		if err != nil {
			log.Fatal(err)
		}
		if rsp.StatusCode != 204 {
			log.Printf("error: received non-successful status code, but we don't want to crash your build pipe so we'll just log it for now: %s", rsp.Status)
		}
	},
}
