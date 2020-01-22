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

	homedir "github.com/mitchellh/go-homedir"
	av "github.com/sahilsk/hundun/avclient"
	c "github.com/sahilsk/hundun/config"
	log "github.com/sahilsk/hundun/logger"
	pg "github.com/sahilsk/hundun/pgclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var hundunConfig c.HundunConfig
var pgclient *pg.PgClient
var avclient *av.AVClient
var verbose bool
var logger *log.Clogger

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
		fmt.Printf("Error executing root: %s", err)
		os.Exit(-1)
	}
}

func init() {

	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hundun.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	logger = log.NewLogger(verbose)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yml")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatal("%s", err)
			os.Exit(-1)
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
		logger.Info("Using config file: %s", viper.ConfigFileUsed())
	}
	err := viper.Unmarshal(&hundunConfig)
	if err != nil {
		logger.Fatal("unable to marshal config to YAML: %v", err)
	}

	if hundunConfig.Pagerduty.ApiKey == "" || hundunConfig.Pagerduty.Url == "" {
		logger.Fatal("Please make sure your config file has api key and pagerduty api url to connect to.")
	}

	if hundunConfig.Alienvault.ApiKey == "" || hundunConfig.Alienvault.Cookie == "" || hundunConfig.Alienvault.Url == "" {
		logger.Fatal("Please make sure your config file has alienvault api key or cookie, and url.")
	}

	pgclient = pg.NewPgClient(hundunConfig.Pagerduty.ApiKey,
		hundunConfig.Pagerduty.Url,
		hundunConfig.Pagerduty.Email,
		verbose)

	avclient = av.NewAVClient(hundunConfig.Alienvault.ApiKey, hundunConfig.Alienvault.Url,
		hundunConfig.Alienvault.Cookie,
		verbose)
}
