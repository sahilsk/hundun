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

// notesCmd represents the notes command
var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("notes called")

		iv, err := pgclient.ListChild("incidents", "notes", notesParams.IncidentId, url.Values{})
		if err != nil {
			log.Fatal("Can't pull note list: %s", err)
		}
		noteList := iv.(ps.Notes)

		fmt.Printf("%s", noteList.ToPrettyString())

	},
}

type NotesParams struct {
	IncidentId string
}

var notesParams NotesParams

func init() {
	getCmd.AddCommand(notesCmd)

	notesCmd.Flags().StringVarP(&notesParams.IncidentId, "incident_id", "i", "", "Incident id")
	notesCmd.MarkFlagRequired("incident_id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
