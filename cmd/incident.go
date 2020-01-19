/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	ps "github.com/sahilsk/hundun/pgclient/schema"
	"github.com/spf13/cobra"
)

// incidentCmd represents the incident command
var incidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("incident called")

		iv, err := pgclient.Get("incidents", incidentParams.incidentID)
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		incident := iv.(ps.IncidentResponse)

		logger.Info("%+v", incident)

		fmt.Print(incident.ToPrettyString())
	},
}

type IncidentParams struct {
	incidentID string
}

var incidentParams IncidentParams

func init() {
	describeCmd.AddCommand(incidentCmd)
	incidentCmd.Flags().StringVar(&incidentParams.incidentID, "id", "", "Incident id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incidentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incidentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
