package main

import (
  "fmt"
  "os"
  "github.com/Infomaker/opencontent-client-go/host"
  "github.com/Infomaker/opencontent-client-go/requests"
)

func main() {
  host := host.NewOpenContentHost()

  host.Host = "192.168.1.109"

  req := requests.NewDeleteRequest(host)
  req.SetUuid("110151d0-9467-4334-b4e5-3f7e3a60c226")
  err := req.Delete()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

}
