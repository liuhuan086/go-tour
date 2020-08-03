package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime string
var duration string

// 每一个子命令，都需要到项目的cmd/root中的rootCmd进行注册，否则无法使用
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间处理格式",
	Long:  "时间处理格式",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "Get now time",
	Long:  "Get now time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("The output is: %s, %d", nowTime.Format("2006-01-02 15:04:05"),
			nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "compute the time that to need",
	Long:  "compute the time that to need",
	Run: func(cmd *cobra.Command, args []string) {
		var CurrentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			CurrentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			CurrentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				// judge is calculateTime has space
				t, _ := strconv.Atoi(calculateTime)
				CurrentTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(CurrentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("The output is: %s, %d", calculateTime.Format(layout),
			calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "",
		"The time that need to compute, the availability unit is timestamp or the time has been formatted")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "",
		`time of duration, the availability unit is "ns", "us" (or "us"), "ms", "m", "h"`)
}
