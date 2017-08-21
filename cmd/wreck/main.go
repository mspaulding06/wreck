// Copyright 2017 Matt Spaulding.  All rights reserved.

// Package main implements the wreck command line tool
package main

import (
	"fmt"
	"log"

	"github.com/mspaulding06/wreck"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "wreck",
	Short: "This tool gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalln("must use a subcommand")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must set URL!")
		}
		content, err := wreck.Get(args[0])
		if err != nil {
			log.Fatalln("GET request failed: %v", err)
		}
		fmt.Printf(content)
	},
}

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("must set URL!")
		}
		content, err := wreck.Post(args[0])
		if err != nil {
			log.Fatalln("POST request failed: %v", err)
		}
		fmt.Printf(content)
	},
}

func main() {
	cmd.PersistentFlags().StringP("username", "u",
		wreck.UserConfig.GetString("credentials.username"), "Username for credentials")
	cmd.PersistentFlags().StringP("password", "p",
		wreck.UserConfig.GetString("credentials.password"), "Password for credentials")
	wreck.UserConfig.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	wreck.UserConfig.BindPFlag("password", cmd.PersistentFlags().Lookup("password"))
	postCmd.PersistentFlags().StringP("content", "c", "", "Content for POST")
	wreck.UserConfig.BindPFlag("content", postCmd.PersistentFlags().Lookup("content"))
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	cmd.Execute()
}
