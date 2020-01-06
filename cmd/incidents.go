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
	"log"
	"net/url"
	"time"

	ps "github.com/sahilsk/hundun/pgclient/schema"
	"github.com/spf13/cobra"
)

// incidentsCmd represents the incidents command
var incidentsCmd = &cobra.Command{
	Use:   "incidents",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incidents called")
		var queryParams url.Values = url.Values{}

		queryParams["since"] = []string{filters.since}
		queryParams["until"] = []string{filters.until}
		queryParams["date_range"] = []string{filters.dateRange}
		queryParams["statuses[]"] = filters.statuses
		queryParams["incident_key"] = []string{filters.incidentKey}
		queryParams["service_ids[]"] = filters.serviceIds
		queryParams["team_ids[]"] = filters.teamIds
		queryParams["user_Ids[]"] = filters.userIds
		queryParams["urgencies[]"] = filters.urgencies
		queryParams["time_zone"] = []string{filters.timeZone}
		queryParams["sort_by[]"] = filters.sortBy
		queryParams["include[]"] = filters.include

		if filters.sinceRelative != "" {
			m, _ := time.ParseDuration(fmt.Sprintf("-%s", filters.sinceRelative))
			now := time.Now().UTC()
			then := now.Add(m)
			thenFormatted := then.Format(time.RFC3339)
			log.Printf("Since Relative: %s", thenFormatted)
			queryParams["since"] = []string{thenFormatted}
		}

		iv, err := pgclient.List("incidents", queryParams)
		incidentList := iv.(ps.IncidentsResponse)

		//log.Printf("%+v", incidentList)

		//Print pretty json
		data, _ := incidentList.ToPrettyString()

		log.Printf("%s", string(data))

		log.Printf("Total records fetched: %d", len(incidentList.Incidents))
		if len(incidentList.Incidents) > 0 {
			log.Printf("Incident: %s", incidentList.Incidents[0].Summary)
		}

		if err != nil {
			log.Fatalf("Can't pull incidents: %s", err)
		}
	},
}

func init() {
	getCmd.AddCommand(incidentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incidentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incidentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	incidentsCmd.Flags().StringVar(&filters.since, "since", "", "The start of the date range over which you want to search.")
	incidentsCmd.Flags().StringVar(&filters.until, "until", "", "The end of the date range over which you want to search.")
	incidentsCmd.Flags().StringVar(&filters.dateRange, "date_range", "", "When set to all, the since and until parameters and defaults are ignored.")
	incidentsCmd.Flags().StringArrayVar(&filters.statuses, "statuses", []string{"acknowledged", "triggered", "resolved"}, "Statuses to filter: triggered, acknowledged or resolved")
	incidentsCmd.Flags().StringVar(&filters.incidentKey, "incident_key", "", "Incident de-duplication key")
	incidentsCmd.Flags().StringArrayVar(&filters.serviceIds, "service_ids", []string{}, `Returns only the incidents associated with the passed service(s). This expects one or more service IDs.
	`)
	incidentsCmd.Flags().StringArrayVar(&filters.teamIds, "team_ids", []string{}, `An array of team IDs. Only results related to these teams will be returned. Account must have the teams ability to use this parameter.	`)
	incidentsCmd.Flags().StringArrayVar(&filters.userIds, "user_ids", []string{}, `Returns only the incidents currently assigned to the passed user(s). This expects one or more user IDs`)
	incidentsCmd.Flags().StringArrayVar(&filters.urgencies, "urgencies", []string{}, "Array of the urgencies of the incidents to be returned. Defaults to all urgencies. Account must have the urgencies ability to do this")

	incidentsCmd.Flags().StringVar(&filters.timeZone, "time_zone", "UTC", "Time zone in which dates in the result will be rendered.")
	incidentsCmd.Flags().StringArrayVar(&filters.sortBy, "sort_by", []string{}, `Used to specify both the field you wish to sort the results on (incident_number/created_at/resolved_at/urgency), as well as the direction (asc/desc) of the results. The sort_by field and direction should be separated by a colon. A maximum of two fields can be included, separated by a comma. Sort direction defaults to ascending. The account must have the urgencies ability to sort by the urgency.`)
	incidentsCmd.Flags().StringArrayVar(&filters.include, "include", []string{}, "Include: users, services, etc")

	incidentsCmd.Flags().StringVar(&filters.sinceRelative, "since_relative", "", `
					Relative duration to fetch incidents since. 
					eg. '1m30s' retrives incidents created 90seconds ago. 
					A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix,
					such as "300ms", "-1.5h" or "2h45m". 
					Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
					`)

}
