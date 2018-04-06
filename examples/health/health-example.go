package main

import (
  "fmt"
  "os"

  "github.com/Infomaker/opencontent-client-go/host"
  "github.com/Infomaker/opencontent-client-go/requests"
)

func main() {
  host := host.NewOpenContentHost()

  host.Host = "34.241.192.167"

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
