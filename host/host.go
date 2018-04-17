package host

import (
  "os"
  "net/url"
  "fmt"
  "github.com/go-resty/resty"
)

type OpenContentHost struct {
  Scheme   string `json:"schema"`
  Host     string `json:"host"`
  User     string `json:"user"`
  Password string `json:"password"`
  Port     string  `json:"port"`
  Baseurl  string  `json:"baseurl"`
}

type OpenContentResponse struct {
  Err error
  StatusCode int
}

func (e *OpenContentResponse) Error() string {
  return e.Error()
}

func getEnvironmentVariable(name, defaultValue string) string {
  value := os.Getenv(name)
  if len(value) == 0 {
    return defaultValue
  }
  return value
}


// Creates a new Open Content host, if environment variables set they are used
// Environment variables :
//
func NewOpenContentHost() OpenContentHost {

  host := OpenContentHost{}
  host.Scheme = getEnvironmentVariable("OC_SCHEME", "http")
  host.Host = getEnvironmentVariable("OC_HOST", "localhost")
  host.User = getEnvironmentVariable("OC_USER", "admin")
  host.Password = getEnvironmentVariable("OC_PWD", "admin")
  host.Port = getEnvironmentVariable("OC_PORT", "8080")
  host.Baseurl = getEnvironmentVariable("OC_BASEURL", "/opencontent")
  return host
}

func (host *OpenContentHost) CreateUrl() string {
  url := url.URL{}
  url.Scheme = host.Scheme
  url.Host = fmt.Sprint(host.Host, ":", host.Port)
  url.Path = host.Baseurl
  return url.String()
}

func statusOk(statusCode int, status []int) bool {
  for _, s := range status {
    if statusCode == s {
      return true
    }
  }
  return false
}

func IsResponseOk(response resty.Response, err error, okstatus... int) bool {
  if err != nil {
    fmt.Println("ERROR: ", err.Error())
    return false
  }

  if ! statusOk(response.StatusCode(), okstatus) {
    var body string
    if len(response.Body()) > 0 {
      body = fmt.Sprint("Body: ", string(response.Body()))
    }
    fmt.Println("ERROR: ", response.Status(), body)
    return false
  }

  return true
}
