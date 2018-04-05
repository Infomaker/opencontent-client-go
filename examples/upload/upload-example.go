package main

import (
  "github.com/hansbringert/opencontent-client/requests"
  "github.com/hansbringert/opencontent-client/ochost"
  "fmt"
  "os"
)

func main() {
  host := ochost.NewOpenContentHost()

  // upload an article
  req := requests.NewUploadRequest(host)
  req.Id = "43e27fec-7262-4153-8355-2367c2c39b6d"
  req.Source = "some source"
  req.PrimaryFile("/local/infomaker/projects/opencontent-learn/1-upload-article/my-article.xml", "application/vnd.infomaker.se-lab.article")
  req.MetadataFile("/local/infomaker/projects/opencontent-learn/1-upload-article/my-article.xml", "application/vnd.infomaker.se-lab.article")

  resp, err := req.Upload()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  fmt.Println(resp.Status(), string(resp.Body()))


  // upload an image
  req = requests.NewUploadRequest(host)
  req.Id = "43e27fec-7262-4153-8355-2367c2c39b6c"
  req.Source = "some source for image"
  req.PrimaryFile("/local/infomaker/projects/opencontent-learn/2-upload-image/streckgubbe.jpg", "image/jpeg")
  req.MetadataFile("/local/infomaker/projects/opencontent-learn/2-upload-image/image-meta.xml", "application/vnd.infomaker.se-lab.image")
  req.PreviewFile("/local/infomaker/projects/opencontent-learn/2-upload-image/preview.jpg", "image/jpeg")
  req.ThumbFile("/local/infomaker/projects/opencontent-learn/2-upload-image/preview.jpg", "image/jpeg")

  resp, err = req.Upload()
  if err != nil {
    fmt.Println("ERROR:", err.Error())
    os.Exit(1)
  }

  fmt.Println(resp.Status(), string(resp.Body()))

}
