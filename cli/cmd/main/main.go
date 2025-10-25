package main

import (
	"log"
	"pusherGo/domain"
	"pusherGo/function"
	"pusherGo/global"
	"pusherGo/global/initialize"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pusher-go",
	Short: "A simple cli tool to retrieve stocks news and push to email",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := initialize.GlobalInit()
		if err != nil {
			return nil
		}
		return nil
	},
}

var news = &cobra.Command{
	Use:   "news",
	Short: "retrieve stock news from API",
	Run: func(cmd *cobra.Command, args []string) {
		// retrieve stock news from API
		news, err := function.GetNews()
		if err != nil {
			log.Printf("GetNews(): %s\n", err.Error())
			defer function.SendError(err)
			panic(err)
		}

		newsText, err := function.FormatNews(news)
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}
		// save news to file
		err = function.WriteFile(&domain.SaveRequest{
			FileName: global.Configs.File.FileNameNews,
			Content:  newsText,
		})
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}
	},
}

var model = &cobra.Command{
	Use:   "model",
	Short: "retrieve stock news from API",
	Run: func(cmd *cobra.Command, args []string) {
		newsText, err := function.ReadFile(&domain.ReadRequest{
			FileName: global.Configs.File.FileNameNews,
		})
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}

		response, err := function.CallModel(&domain.ModelCallRequest{
			Model:   global.Configs.Model.ModelName,
			ApiKey:  global.Configs.Model.ApiKey,
			Content: newsText,
		})
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}

		err = function.WriteFile(&domain.SaveRequest{
			FileName: global.Configs.File.FileNameModelResponse,
			Content:  response.Answer,
		})
		if err != nil {
			defer function.SendError(err)
			panic(err)
		}
	},
}

var email = &cobra.Command{
	Use:   "email",
	Short: "push stock news summarization to email",
	Run: func(cmd *cobra.Command, args []string) {
		// push stock news to email
		answer, err := function.ReadFile(&domain.ReadRequest{
			FileName: global.Configs.File.FileNameModelResponse,
		})
		if err != nil {
			panic(err)
		}
		function.SendEmail(answer)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(news)
	rootCmd.AddCommand(model)
	rootCmd.AddCommand(email)
}

func main() {
	rootCmd.Execute()
}
