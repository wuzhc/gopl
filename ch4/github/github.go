package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const api = "https://api.github.com/search/issues"

type SearchResult struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []*Item
}

type Item struct {
	Url           string
	RepositoryUrl string `json:"repository_url"`
	Title         string
	Id            int
	HtmlUrl       string `json:"html_url"`
	CreateAt      string `json:"create_at"`
	UpdateAt      string `json:update_at`
	Body          string
	score         float64
	User          *User
	Labels        []*Label
}

type User struct {
	Login     string
	Id        int
	Url       string
	AvatarUrl string `json:"avatar_url"`
}

type Label struct {
	Id     int
	NodeId string `json:"node_id"`
	Url    string
	Name   string
	Color  string
}

func SearchIssues(terms []string) (*SearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(api + "?q=" + q)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	// 请求失败关闭resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed :%s\n", resp.Status)
	}

	var result SearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *User) GetName() string {
	return u.Login
}
