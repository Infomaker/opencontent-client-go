package main

import (
  "github.com/hansbringert/opencontent-client/ochost"
  "github.com/hansbringert/opencontent-client/requests"
  "fmt"
  "os"
)

func main() {
  host := ochost.NewOpenContentHost()

  host.Host = "192.168.1.109"

  req := requests.NewPropertiesRequest(host)
  req.SetUuid("f29ee1d9-2e11-4fbd-b323-675f9289edab")
  req.AddProperty("uuid,version")

  properties, err  := req.GetProperties()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  // Print values
  for _, property := range properties.Properties {
    fmt.Println(property.Name, "=", property.Values)
  }
  // propertiesJsonBytes, err := properties.ToJson()
  // fmt.Println(string(propertiesJsonBytes))
}
