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

	avs "github.com/sahilsk/hundun/avclient/schema"
	"github.com/spf13/cobra"
)

// avAlarmCmd represents the avAlarm command
var avAlarmCmd = &cobra.Command{
	Use:   "avalarm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("avAlarm called")
		iv, err := avclient.Get("alarms", avalarm_params.AlarmId)
		if err != nil {
			logger.Fatal("Error: %s", err)
		}
		avalarm := iv.(avs.Alarm)
		logger.Info("%v", avalarm)
		fmt.Printf("%v", avalarm)
	},
}

type AVAlarmParams struct {
	AlarmId string
}

var avalarm_params AVAlarmParams

func init() {
	describeCmd.AddCommand(avAlarmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// avAlarmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	avAlarmCmd.Flags().StringVarP(&avalarm_params.AlarmId, "alarm_id", "a", "", "alarm id")
}
