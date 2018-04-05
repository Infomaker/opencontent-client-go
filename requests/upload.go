package requests

import (
  "path"
  "github.com/go-resty/resty"
  "fmt"
  "strconv"
  "github.com/Infomaker/opencontent-client-go/host"
  "encoding/json"
)

type UploadRequest struct {
  Host     host.OpenContentHost `json:"host"`
  Id       string                 `json:"id"`
  Batch    bool                   `json:"batch"`
  Source   string                 `json:"source"`
  Primary  FileInfo               `json:"primary"`
  Metadata map[string]FileInfo    `json:"metadata"`
  Preview  FileInfo               `json:"preview`
  Thumb    FileInfo               `json:"thumb`
}

type FileInfo struct {
  Fullpath string `json:"fullpath"`
  Mimetype string `json:"mimetype"`
}

func NewUploadRequest(host host.OpenContentHost) UploadRequest {
  req := UploadRequest{}
  req.Host = host
  req.Metadata = make(map[string]FileInfo)
  return req
}

func (uploadrequest *UploadRequest) PrimaryFile(fileWithFullpath, mimetype string) {
  uploadrequest.Primary.Fullpath = fileWithFullpath
  uploadrequest.Primary.Mimetype = mimetype
}

func (uploadrequest * UploadRequest) MetadataFile(fileWithFullpath, mimetype string) {
  filename := path.Base(fileWithFullpath)
  fileinfo := FileInfo{
    Fullpath: fileWithFullpath,
    Mimetype: mimetype,
  }
  uploadrequest.Metadata[filename] = fileinfo
}

func (uploadrequest *UploadRequest) PreviewFile(fileWithFullpath, mimetype string) {
  uploadrequest.Preview.Fullpath = fileWithFullpath
  uploadrequest.Preview.Mimetype = mimetype
}

func (uploadrequest *UploadRequest) ThumbFile(fileWithFullpath, mimetype string) {
  uploadrequest.Thumb.Fullpath = fileWithFullpath
  uploadrequest.Thumb.Mimetype = mimetype
}


func (req *UploadRequest) GetUrl() string {
  return fmt.Sprint(req.Host.CreateUrl(), "/objectupload?")
}

func (req *UploadRequest) createFilesMap() map[string]string {
  filesmap := make(map[string]string)

  // add the primary file
  filesmap[path.Base(req.Primary.Fullpath)] = req.Primary.Fullpath

  if len(req.Preview.Fullpath) > 0 {
    filesmap[path.Base(req.Preview.Fullpath)] = req.Preview.Fullpath
  }

  if len(req.Thumb.Fullpath) > 0 {
    filesmap[path.Base(req.Thumb.Fullpath)] = req.Thumb.Fullpath
  }

  // add the metadata files
  for key,metadata := range req.Metadata {
    filesmap[key] = metadata.Fullpath
  }
  return filesmap
}

func (req *UploadRequest) createFormMap() map[string]string {
  formmap := make(map[string]string)

  formmap["id"] = req.Id
  formmap["source"] = req.Source
  formmap["batch"] = strconv.FormatBool(req.Batch)


  // add primary file mimetyp
  formmap["file"] = path.Base(req.Primary.Fullpath)
  formmap["file-mimetype"] = req.Primary.Mimetype

  // add preview file mimetyp
  if len(req.Preview.Fullpath) > 0 {
    formmap["preview"] = path.Base(req.Preview.Fullpath)
    formmap["preview-mimetype"] = req.Preview.Mimetype
  }

  // add thumb file thumb
  if len(req.Thumb.Fullpath) > 0 {
    formmap["thumb"] = path.Base(req.Thumb.Fullpath)
    formmap["thumb-mimetype"] = req.Thumb.Mimetype
  }

  index := 1
  indexStr := ""
  for key,metadata := range req.Metadata {
    if index > 1 {
      indexStr = strconv.Itoa(index)
    }
    formmap["metadata" + indexStr] = key
    formmap["metadata" + indexStr + "-mimetype"] = metadata.Mimetype
    index++
  }
  return formmap
}

func (req *UploadRequest) Upload() (*resty.Response, error) {

  filesMap := req.createFilesMap()
  formMap := req.createFormMap()

  resp, err := resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).SetFiles(filesMap).SetFormData(formMap).Post(req.GetUrl())
  return resp, err
}

func (uploadrequest * UploadRequest) toString() string{
  jsonBytes, _ := json.MarshalIndent(uploadrequest, "", "  ")
  return string(jsonBytes)
}