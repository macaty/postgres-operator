package cmd

/*
 Copyright 2017 Crunchy Data Solutions, Inc.
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

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	"github.com/crunchydata/postgres-operator/pgo/api"
	"github.com/spf13/cobra"
	"os"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test cluster connectivity",
	Long: `TEST allows you to test the connectivity for a cluster. For example:

	pgo test mycluster
	pgo test --selector=env=research
	pgo test all`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("test called")
		if Selector == "" && len(args) == 0 {
			fmt.Println(`Error: You must specify the name of the clusters to test.`)
		} else {
			if OutputFormat != "" && OutputFormat != "json" {
				fmt.Println("Error: Only 'json' is currently supported for the --output flag value.")
				os.Exit(2)
			}
			showTest(args, Namespace)
		}
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&Selector, "selector", "s", "", "The selector to use for cluster filtering.")
	testCmd.Flags().StringVarP(&OutputFormat, "output", "o", "", "The output format. Currently, json is the only supported value.")
}

func showTest(args []string, ns string) {

	log.Debugf("showCluster called %v", args)

	log.Debugf("selector is %s", Selector)
	if len(args) == 0 && Selector != "" {
		args = make([]string, 1)
		args[0] = "all"
	}

	for _, arg := range args {
		response, err := api.ShowTest(httpclient, arg, Selector, &SessionCredentials, ns)
		if err != nil {
			fmt.Println("Error: " + err.Error())
			os.Exit(2)
		}

		if response.Status.Code != msgs.Ok {
			fmt.Println("Error: " + response.Status.Msg)
			os.Exit(2)
		}

		if OutputFormat == "json" {
			b, err := json.MarshalIndent(response, "", "  ")
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println(string(b))
			return
		}

		if len(response.Results) == 0 {
			fmt.Println("Nothing found.")
			return
		}

		for _, result := range response.Results {
			fmt.Println("")
			fmt.Printf("cluster : %s \n", result.ClusterName)
			for _, v := range result.Items {
				fmt.Printf("%s%s is ", TreeBranch, v.PsqlString)
				if v.Working {
					fmt.Printf("%s\n", GREEN("Working"))
				} else {
					fmt.Printf("%s\n", RED("NOT working"))
				}
			}
		}

	}
}
