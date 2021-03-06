package sonar2csv

import (
	"fmt"
	"github.com/FujitsuLaboratories/COMEVIZZ/sonar2csv/util"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SonarServer struct {
	Client   *http.Client
	URL      *url.URL
	Resource string
	Metrics  []string
}

func NewServer(c util.SonarSetting) (*SonarServer, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}
	return &SonarServer{
		Client:   &http.Client{},
		URL:      u,
		Resource: c.Resource,
		Metrics:  c.Metrics,
	}, nil
}

func (s *SonarServer) GetResources() (*http.Response, error) {
	s.URL.Path = "/api/resources"

	q := url.Values{}
	q.Set("resource", s.Resource)
	q.Set("metrics", strings.Join(s.Metrics, ","))
	s.URL.RawQuery = q.Encode()
	return s.Get()
}

func (s *SonarServer) Get() (*http.Response, error) {
	log.Print("Request to ", s.URL.String())
	resp, err := s.Client.Do(&http.Request{
		Method: "GET",
		URL:    s.URL,
	})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}
