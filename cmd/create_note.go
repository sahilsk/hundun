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
	"fmt"
	"log"
	"net/url"

	ps "github.com/sahilsk/hundun/pgclient/schema"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("note called")

		notePayloadJSON := []byte(fmt.Sprintf(`
		{
			"note": {
			  "content": "%s"
			}
		}
		`, createNoteParams.Content))

		iv, err := pgclient.PostChild("incidents", "notes", createNoteParams.IncidentId, url.Values{}, notePayloadJSON)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}

		note := iv.(ps.Note)

		logger.Info("%+v", note)
		fmt.Printf("%s", note.ToPrettyString())

	},
}

type CreateNoteParams struct {
	IncidentId string
	Content    string
}

var createNoteParams CreateNoteParams

func init() {
	createCmd.AddCommand(noteCmd)

	noteCmd.Flags().StringVarP(&createNoteParams.Content, "content", "c", "", "note content")
	noteCmd.Flags().StringVarP(&createNoteParams.IncidentId, "incident_id", "i", "", "incident id")
	noteCmd.MarkFlagRequired("incident_id")
	noteCmd.MarkFlagRequired("content")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// noteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
