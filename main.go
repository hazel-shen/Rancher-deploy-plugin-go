package main

import (
    "os"
	"fmt"
	"io/ioutil"
	"log"
    "net/http"
    "encoding/base64"
    "github.com/tidwall/sjson"
    s "strings"
)
func basicAuth(username, password string) string {
    auth := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(auth))
}
    
    
func main() {
  
    PLUGIN_PROJECT_API := os.Getenv("PLUGIN_PROJECT_API")
    PLUGIN_DEPLOY_IMAGE := os.Getenv("PLUGIN_DEPLOY_IMAGE")
    RANCHER_ACCESS_KEY := os.Getenv("RANCHER_ACCESS_KEY")
    RANCHER_SECRET_KEY := os.Getenv("RANCHER_SECRET_KEY")
    var client http.Client

    req, err := http.NewRequest("GET", PLUGIN_PROJECT_API, nil)
    req.SetBasicAuth(RANCHER_ACCESS_KEY, RANCHER_SECRET_KEY)

    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    } 

    res, err2 := client.Do(req)
  
	if err2 != nil {
        log.Fatal(err2)
        os.Exit(1)
    }    

	defer res.Body.Close()
    apiResponse, err := ioutil.ReadAll(res.Body)
    var apiResponseStr string = string(apiResponse[:])
    if s.Contains(apiResponseStr, "error") {
        fmt.Printf("%s", apiResponse)
        os.Exit(1)
    }

    value2, _ := sjson.Set(apiResponseStr, `containers.0.image`, PLUGIN_DEPLOY_IMAGE)
    fmt.Println(value2)
    
    req3, err3 := http.NewRequest("PUT", PLUGIN_PROJECT_API, s.NewReader(value2))
    req3.SetBasicAuth(RANCHER_ACCESS_KEY, RANCHER_SECRET_KEY)

    if err3 != nil {
        log.Fatal(err3)
        os.Exit(1)
    } 

    res4, err4 := client.Do(req3)
  
	if err4 != nil {
        log.Fatal(err4)
        os.Exit(1)
    }    

    defer res4.Body.Close()
    apiResponse2, err2 := ioutil.ReadAll(res4.Body)
    var apiResponseStr2 string = string(apiResponse2[:])
    if s.Contains(apiResponseStr2, "error") {
        fmt.Printf("%s", apiResponse2)
        os.Exit(1)
    }
}