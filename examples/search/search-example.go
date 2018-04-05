package main

import (
  "fmt"
  "os"
  "github.com/gin-gonic/gin/json"
  "bytes"
  "github.com/Infomaker/opencontent-client-go/host"
  "github.com/Infomaker/opencontent-client-go/requests"
)

func main() {
  host := host.NewOpenContentHost()

  host.Host = "127.0.0.1"

  req := requests.NewSearchRequest(host)
  req.SetQuery("Product:*")
  req.AddHighlightIndexField("Product")
  searchResponse, err := req.Search()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  jsonBytes, err := json.MarshalIndent(searchResponse, "", "  ")
  jsonBytes = bytes.Replace(jsonBytes, []byte("\\u003c"), []byte("<"), -1)
  jsonBytes = bytes.Replace(jsonBytes, []byte("\\u003e"), []byte(">"), -1)
  jsonBytes = bytes.Replace(jsonBytes, []byte("\\u0026"), []byte("&"), -1)

  fmt.Println(string(jsonBytes))
}