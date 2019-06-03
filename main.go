package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	s "strings"

	"github.com/tidwall/pretty"
	"github.com/tidwall/sjson"
)

func main() {

	pluginProjectAPI := os.Getenv(`PLUGIN_PROJECT_API`)
	pluginDeployImage := os.Getenv(`PLUGIN_DEPLOY_IMAGE`)
	rancherAccessKey := os.Getenv(`PLUGIN_ACCESS_KEY`)
	rancherSecretKey := os.Getenv(`PLUGIN_SECRET_KEY`)

	var client http.Client

	req, _ := http.NewRequest("GET", pluginProjectAPI, nil)
	req.SetBasicAuth(rancherAccessKey, rancherSecretKey)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	projectInfo, _ := ioutil.ReadAll(res.Body)
	projectInfoStr := string(projectInfo[:])

	if s.Contains(projectInfoStr, "error") {
		os.Exit(1)
	}

	updatedProjectInfo, _ := sjson.Set(projectInfoStr, `containers.0.image`, pluginDeployImage)

	req, _ = http.NewRequest("PUT", pluginProjectAPI, s.NewReader(updatedProjectInfo))
	req.SetBasicAuth(rancherAccessKey, rancherSecretKey)

	res, err = client.Do(req)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer res.Body.Close()
	updatedResult, _ := ioutil.ReadAll(res.Body)
	updatedResultStr := string(updatedResult)
	updatedResultJSON := pretty.Pretty(updatedResult)

	fmt.Printf("%s", updatedResultJSON)
	if s.Contains(updatedResultStr, "error") {
		os.Exit(1)
	}
}
