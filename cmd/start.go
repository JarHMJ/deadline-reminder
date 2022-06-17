/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/JarHMJ/deadline-reminder/pkg/receiver/feishu"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var feishuClient *feishu.Feishu = feishu.NewFeishu(viper.GetString("feishu.baseUrl"))

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		c := cron.New()
		_, err := c.AddFunc("0 8 * * *", func() {
			fmt.Println("现在是早上8点，起床study，奥力给！！！")
		})
		if err != nil {
			log.Fatalf("cron add func err: %s \n", err)
		}
		c.Start()
		log.Println("starting")

		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
			syscall.SIGQUIT)

		select {
		case <-s:
			c.Stop()
			log.Println("exiting")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
