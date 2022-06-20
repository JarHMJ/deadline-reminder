/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type feishuConfig struct {
	BaseUrl string `mapstructure:"baseUrl"`
}

type config struct {
	Feishu   feishuConfig `mapstructure:"feishu"`
	Deadline *time.Time   `mapstructure:"deadline"`
	Name     string       `mapstructure:"name"`
}

var cfgFile string
var cfg config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "deadline-reminder",
	Aliases: []string{"dlr"},
	Short:   "A deadline reminder",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.deadline-reminder.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		log.Printf("home :%s", home)

		// Search config in home directory with name ".deadline-reminder" (without extension).
		viper.AddConfigPath(home)
		//viper.SetConfigType("yaml")
		viper.SetConfigName(".deadline-reminder")
	}

	bindEnv()
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatalln(err)
	}

	// https://github.com/spf13/viper/issues/496
	deadline := viper.GetTime("deadline")
	viper.Set("deadline", deadline)

	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unmarshal config err: %s \n", err)
	}
	log.Printf("cfg :%+v", cfg)
}

func bindEnv() {
	viper.BindEnv("deadline", "deadline")
	viper.BindEnv("feishu.baseUrl", "feishu.baseUrl")
	viper.BindEnv("name", "name")
}
