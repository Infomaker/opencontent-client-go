package requests

import (
  "github.com/hansbringert/opencontent-client/ochost"
  "github.com/go-resty/resty"
  "fmt"
  "strings"
  "errors"
  "encoding/json"
)

type SearchRequest struct {
  Host                ochost.OpenContentHost
  Start               int
  Limit               int
  Property            []string
  Properties          []string
  Filters             []string
  Contenttype         []string
  Deleted             bool
  Timezone            string
  Q                   string
  Facet               bool
  FacetIndexField     []string
  FacetLimit          int
  FacetMinCount       int
  FacetDateMinCount   string
  FacetDateStart      string
  FacetDateEnd        string
  SortIndexField      []string
  SortAscending       []string
  SortName            string
  HighlightIndexField []string
  HighlightPre        string
  HighlitePost        string

  Quiet               bool // only display uuids
}

func NewSearchRequest(host ochost.OpenContentHost) SearchRequest {
  req := SearchRequest{}
  req.Q = "*:*"
  req.Host = host
  req.Start = 0
  req.Limit = 15
  return req
}

func (req *SearchRequest) SetQuery(query string) SearchRequest {
  req.Q = query
  return *req
}

func (req *SearchRequest) SetStart(start int) SearchRequest {
  req.Start = start
  return *req
}

func (req *SearchRequest) SetLimit(limit int) SearchRequest {
  req.Limit = limit
  return *req
}

func (req *SearchRequest) AddHighlightIndexField(field string) SearchRequest {
  req.HighlightIndexField = append(req.HighlightIndexField, field)
  return *req
}

func (req *SearchRequest) GetUrl() string {
  httpQuery := "?"
  httpQuery = fmt.Sprint(httpQuery, "start=", req.Start, "&")
  httpQuery = fmt.Sprint(httpQuery, "limit=", req.Limit, "&")
  httpQuery = fmt.Sprint(httpQuery, "q=", req.Q, "&")
  for  _, hif := range req.HighlightIndexField {
    httpQuery = fmt.Sprint(httpQuery, "highlight.indexfield=", hif, "&")
  }
  return fmt.Sprint(req.Host.CreateUrl(), "/search", strings.TrimSuffix(httpQuery, "&"))
}

func (req *SearchRequest) execute() (*resty.Response, error) {
  return resty.SetDisableWarn(true).R().SetBasicAuth(req.Host.User, req.Host.Password).Get(req.GetUrl())
}

func (req *SearchRequest) Search() (SearchResponse, error){
  response, err := req.execute()
  if err != nil {
    return SearchResponse{}, err
  }
  if response.StatusCode() != 200 {
    return SearchResponse{}, errors.New(response.Status() + " " + string(response.Body()))
  }
  var searchResponse SearchResponse
  err = json.Unmarshal(response.Body(), &searchResponse)
  if err != nil {
    return searchResponse, err
  }
  return searchResponse, nil
}


//
// Search result
//
type SearchResponse struct {
  Hits  Hits `json:"hits"`
  Facet FacetFields `json:"facet"`
  Stats Stats `json:"stats"`
  Highlight interface{} `json:"highlight"`
}

type Hits struct {
  TotalHits    int   `json:"totalHits"`
  Hits         []Hit `json:"hits"`
  IncludedHits int   `json:"includedHits"`
}

type Hit struct {
  Id         string `json:"id"`
  Versions   []Version `json:"versions"`
  NoVersions int `json:"noversions"`
}

type Version struct {
  Id         int `json:"id"`
  Properties interface{} `json:"properties"`
}

type FacetFields struct {
  Fields []FacetField `json:"fields"`
}

type FacetField struct {
  Year        int `json:"year"`
  Month       int `json:"month"`
  Day         int `json:"day"`
  FacetField  string `json:"facetField"`
  Frequencies []Frequency `json:"frequencies"`
}

type Frequency struct {
  Term      string `json:"term"`
  Frequency int `json:"frequency"`
}

type Stats struct {
  Duration int `json:"duration"`
}
