package main

import (
	"pusherGo/function"
	"pusherGo/global/initialize"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pusher-go",
	Short: "A simple cli tool to retrieve stocks news and push to email",
}

// Run initialization, load configurations
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the application",
	Run: func(cmd *cobra.Command, args []string) {
		err := initialize.GlobalInit()
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}
	},
}

var news = &cobra.Command{
	Use:   "news",
	Short: "retrieve stock news from API",
	Run: func(cmd *cobra.Command, args []string) {
		// retrieve stock news from API
	},
}

var model = &cobra.Command{
	Use:   "model",
	Short: "retrieve stock news from API",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var email = &cobra.Command{
	Use:   "email",
	Short: "push stock news summarization to email",
	Run: func(cmd *cobra.Command, args []string) {
		// push stock news to email
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(news)
	rootCmd.AddCommand(model)
	rootCmd.AddCommand(email)
}

func main() {
	rootCmd.Execute()
}
