package requests

import (
  "fmt"
  "github.com/go-resty/resty"
  "encoding/json"
  "errors"
  "github.com/Infomaker/opencontent-client-go/host"
)

type HealthRequest struct {
  Host host.OpenContentHost
}

func NewHealthRequest(host host.OpenContentHost) HealthRequest {
  req := HealthRequest{}
  req.Host = host
  return req
}

func (req *HealthRequest) GetUrl() string {
  return fmt.Sprint(req.Host.CreateUrl(), "/health")
}

func (req *HealthRequest) execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().Get(req.GetUrl())
}

func (req *HealthRequest) Health() (HealthResponse, error){
  response, err := req.execute()
  if err != nil {
    return HealthResponse{}, err
  }
  if response.StatusCode() != 200 {
    return HealthResponse{}, errors.New(response.Status() + " " + string(response.Body()))
  }
  var health HealthResponse
  err = json.Unmarshal(response.Body(), &health)
  if err != nil {
    return health, err
  }

  return health, nil
}

// Response from health
type HealthResponse  struct {
  Indexer             bool `json:"indexer"`
  Index               bool `json:"index"`
  Database            bool `json:"database"`
  FileSystem          bool `json:"filesystem"`
  FreeSystemDiskSpace int  `json:"freeSystemDiskSpace"`
  MaximumMemory       int `json:"maximumMemory"`
  CurrentMemory       int `json:"currentMemory"`
  ActiveConfiguration Configuration `json:"activeConfiguration"`
  TempConfiguration   Configuration `json:"tempConfiguration"`
}


type Configuration struct {
  Checksum     string `json:"checksum"`
  LastModified string `json:"lastModified"`
}

func (hr *HealthResponse) ToJson() ([]byte, error){
  bytes, err :=  json.MarshalIndent(hr, " ", " ", )
  return bytes, err
}

func HealthResponseFromBytes(jsonBytes []byte) (HealthResponse, error) {
  health := HealthResponse{}
  err := json.Unmarshal(jsonBytes, &health)
  return health, err
}
