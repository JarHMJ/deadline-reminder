/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/JarHMJ/deadline-reminder/pkg/receiver/feishu"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
		feishuClient := feishu.NewFeishu(cfg.Feishu.BaseUrl)

		c := cron.New()
		_, err := c.AddFunc("0 8 * * *", func() {
			log.Println("现在是早上8点，起床study，奥力给！！！")
			now := time.Now()
			days := calculateDays(&now, cfg.Deadline)
			msg := fmt.Sprintf("距离%s还有 %d 天 <at user_id=\\\"all\\\">所有人</at>", cfg.Name, days)
			feishuClient.Notify(msg)
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

func checkConfig() {
	if cfg.Feishu.BaseUrl == "" {
		log.Fatalf("feishu.baseUrl is missing")
	}

	if cfg.Deadline == nil {
		log.Fatalf("deadline is missing")
	}

	if cfg.Name == "" {
		log.Fatalf("Name is missing")
	}
}

func calculateDays(now, deadline *time.Time) int {
	//deadline.
	newNow := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	newDeadline := time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, time.Local)

	return int(newDeadline.Sub(newNow).Hours() / 24)
}
