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

	ps "github.com/sahilsk/hundun/pgclient/schema"
	"github.com/spf13/cobra"
)

// priorityCmd represents the priority command
var priorityCmd = &cobra.Command{
	Use:   "priority",
	Short: "A brief description of your command",
	Long:  `Get priority detail out of a priority id`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("priority called")
		iv, err := pgclient.Get("priorities", priorityParams.id)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
		priority := iv.(ps.PriorityResponse)

		logger.Info("%+v", priority)
		priorityStr, err := json.MarshalIndent(priority, "", "  ")
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Print(string(priorityStr))
	},
}

type PriorityParams struct {
	id string
}

var priorityParams PriorityParams

func init() {
	describeCmd.AddCommand(priorityCmd)
	priorityCmd.Flags().StringVar(&priorityParams.id, "id", "", "priority id")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priorityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priorityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
