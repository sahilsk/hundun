/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"log"
	"net/url"

	ps "github.com/sahilsk/hundun/pgclient/schema"

	"github.com/spf13/cobra"
)

// ackCmd represents the ack command
var ackCmd = &cobra.Command{
	Use:   "ack",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("ack called")

		incidentAckPayloadJSON := []byte(`
		{
			"incident": {
			  "type": "incident_reference",
			  "status": "acknowledged"
			}
		}
		`)

		iv, err := pgclient.Put("incidents", ackParams.incidentID, url.Values{}, incidentAckPayloadJSON)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		incident := iv.(ps.IncidentResponse)

		logger.Info("%+v", incident)
		incidentStr, err := json.MarshalIndent(incident, "", "  ")
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Print(string(incidentStr))
	},
}

type AckParams struct {
	incidentID string
}

var ackParams AckParams

func init() {
	rootCmd.AddCommand(ackCmd)
	ackCmd.Flags().StringVar(&ackParams.incidentID, "id", "", "Incident id")
}
