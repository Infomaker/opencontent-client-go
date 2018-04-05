package main

import (
  "fmt"
  "os"
  "github.com/hansbringert/opencontent-client/requests"
  "github.com/hansbringert/opencontent-client/ochost"
)

func main() {
  host := ochost.NewOpenContentHost()

  host.Host = "192.168.1.109"

  req := requests.NewHealthRequest(host)
  health, err := req.Health()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  healthJson, _ := health.ToJson()
  fmt.Println(string(healthJson))

  healthResponse, err := requests.HealthResponseFromBytes(healthJson)
  // get information from healthResponse struct
  fmt.Println("healthResponse.Indexer = ", healthResponse.Indexer)
}
