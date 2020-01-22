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

	"github.com/spf13/cobra"
)

// resolveCmd represents the resolve command
var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("resolve called")

		incidentResolvePayloadJSON := []byte(fmt.Sprintf(`
		{
			"incident": {
			  "type": "incident_reference",
			  "status": "resolved",
			  "resolution": "%s"
			}
		}
		`, resolveParams.resolution))

		incident, err := pgclient.Put("incidents", resolveParams.incidentID, url.Values{}, incidentResolvePayloadJSON)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		log.Printf("%+v", incident)
		incidentStr, err := json.MarshalIndent(incident, "", "  ")
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Print(string(incidentStr))
	},
}

type ResolveParams struct {
	incidentID string
	resolution string
}

var resolveParams ResolveParams

func init() {
	rootCmd.AddCommand(resolveCmd)
	resolveCmd.Flags().StringVarP(&resolveParams.incidentID, "incident_id", "i", "", "Incident id")
	resolveCmd.Flags().StringVarP(&resolveParams.resolution, "resolution", "r", "", "The resolution for this incident.")
	resolveCmd.MarkFlagRequired("incident_id")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resolveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// resolveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
