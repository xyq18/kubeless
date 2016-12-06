/*
Copyright 2016 Skippbox, Ltd.

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
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var topicCreateCmd = &cobra.Command{
	Use:   "create <topic_name> FLAG",
	Short: "create a topic to Kubeless",
	Long:  `create a topic to Kubeless`,
	Run: func(cmd *cobra.Command, args []string) {
		master, _ := cmd.Flags().GetString("master")
		if master == "" {
			master = "localhost"
		}

		if len(args) != 1 {
			logrus.Fatal("Need exactly one argument - topic name")
		}

		//TODO
		//topicName := args[0]
		//utils.CreateKubelessTopic(topicName, master)
	},
}