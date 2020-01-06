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
	"os"
	"log"

	homedir "github.com/mitchellh/go-homedir"
	c "github.com/sahilsk/hundun/config"
	pg "github.com/sahilsk/hundun/pgclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var hundunConfig c.HundunConfig
var pgclient *pg.PgClient


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hundun",
	Short: "A handy pagerduty client",
	Long: `Hundun: God of chaos and noise in some mythology long dead. 
	You can use this client to resolve alerts and add note to them in a breeze
	n. For example:

	hundun get alerts 
	Or
	hundun describe alert <alert_id>
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log.SetFlags( log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile )
	
	cobra.OnInitialize(initConfig)

	cobra.OnInitialize(initPgClient)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hundun.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yml")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".hundun" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config")
		viper.SetConfigName(".hundun")
		viper.SetDefault("pagerduty.email", "")

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&hundunConfig)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}


	if hundunConfig.Pagerduty.ApiKey == "" {
		log.Fatalf("Please make sure your config file has api key and endpoint to connect to.")
	}
}

func initPgClient() {
	log.Printf("Pagerduty client initializing..")
	log.Printf("\nApi key: %s \nUrl: %s", hundunConfig.Pagerduty.ApiKey, hundunConfig.Pagerduty.Url)

	pgclient = &pg.PgClient{
		ApiKey:   hundunConfig.Pagerduty.ApiKey,
		Endpoint: hundunConfig.Pagerduty.Url,
		Email: hundunConfig.Pagerduty.Email,
	}
	pgclient.Init()

}
