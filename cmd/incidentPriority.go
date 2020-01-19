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
	"fmt"
	"log"
	"net/url"
	"strings"

	ps "github.com/sahilsk/hundun/pgclient/schema"

	"github.com/spf13/cobra"
)

// incidentPriorityCmd represents the ack command
var incidentPriorityCmd = &cobra.Command{
	Use:   "priority",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incidentPriority called")

		var desiredPriority string
		var validInputs []string

		if len(incidentPriorityParams.priority) == 2 {
			//convert the 'summary' to priority id
			iv, err := pgclient.List("priorities", url.Values{})
			if err != nil {
				log.Fatal(err)
			}
			priorityList := iv.(ps.Priorities)

			for _, v := range priorityList.Priorities {
				if strings.ToLower(v.Summary) == strings.ToLower(incidentPriorityParams.priority) {
					desiredPriority = v.Id
					break
				}
				validInputs = append(validInputs, v.Summary)
			}
			log.Printf("%v", priorityList)
		}

		if desiredPriority == "" {
			log.Fatalf("Priority: %s, not found. Valid values are: %s", incidentPriorityParams.priority, validInputs)
		}

		incidentPriorityPayloadJSON := []byte(fmt.Sprintf(`
		{
			"incident": {
			  "type": "incident_reference",
			  "status": "acknowledged",
			  "priority": {
				  "id": "%s",
				  "type": "priority_reference"
			  }
			}
		}
		`, desiredPriority))

		iv, err := pgclient.Put("incidents", incidentPriorityParams.incidentID, url.Values{}, incidentPriorityPayloadJSON)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		incident := iv.(ps.IncidentResponse)

		log.Printf("%+v", incident)
		incidentStr, err := json.MarshalIndent(incident, "", "  ")
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Print(string(incidentStr))
	},
}

type IncidentPriorityParams struct {
	incidentID string
	priority   string
}

var incidentPriorityParams IncidentPriorityParams

func init() {
	setCmd.AddCommand(incidentPriorityCmd)
	incidentPriorityCmd.Flags().StringVar(&incidentPriorityParams.incidentID, "id", "", "Incident id")
	incidentPriorityCmd.Flags().StringVar(&incidentPriorityParams.priority, "priority", "", "Incident priority eg. S1, S3, NI. Set mapping of priority id to names in the config file")

}
