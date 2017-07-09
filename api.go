package github

import (
	"github.com/dforsyth/reflectclient"
	"github.com/dforsyth/reflectclient-github/models"
	"net/http"
)

type GithubServiceV3 struct {
	Repo      func(*RepoParams) (*models.Repo, error)        `rc_method:"GET" rc_path:"/repos/{owner}/{repo}"`
	Repos     func(*ReposParams) ([]*models.Repo, error)     `rc_method:"GET" rc_path:"/user/repos"`
	UserRepos func(*UserReposParams) ([]*models.Repo, error) `rc_method:"GET" rc_path:"/users/{username}/repos"`
}

type RepoParams struct {
	Owner string `rc_feature:"path" rc_name:"owner"`
	Repo  string `rc_feature:"path" rc_name:"repo"`
}

type ReposParams struct {
	Visibility  string `rc_feature:"query" rc_name:"visibility"`
	Affiliation string `rc_feature:"query" rc_name:"affiliation"`
	Type        string `rc_feature:"query" rc_name:"type"`
	Sort        string `rc_feature:"query" rc_name:"sort"`
	Direction   string `rc_feature:"query" rc_name:"direction"`
}

type UserReposParams struct {
	UserName  string `rc_feature:"path" rc_name:"username"`
	Type      string `rc_feature:"query" rc_name:"type"`
	Sort      string `rc_feature:"query" rc_name:"sort"`
	Direction string `rc_feature:"query" rc_name:"direction"`
}

type TokenProvider interface {
	GetUsername() string
	GetToken() string
}

type DefaultTokenProvider struct {
	username, token string
}

func (p *DefaultTokenProvider) GetUsername() string {
	return p.username
}

func (p *DefaultTokenProvider) GetToken() string {
	return p.token
}

func MakeDefaultTokenProvider(username, token string) *DefaultTokenProvider {
	return &DefaultTokenProvider{username: username, token: token}
}

func MakeService(baseUrl string, provider TokenProvider) (*GithubServiceV3, error) {
	service := new(GithubServiceV3)
	client, err := reflectclient.NewBuilder().
		BaseUrl(baseUrl).
		AddRequestTransformer(func(r *http.Request) *http.Request {
			r.Header.Set("Accept", "application/vnd.github.v3+json")
			return r
		}).
		AddRequestTransformer(func(r *http.Request) *http.Request {
			if provider != nil && provider.GetUsername() != "" && provider.GetToken() != "" {
				r.SetBasicAuth(provider.GetUsername(), provider.GetToken())
			}
			return r
		}).
		SetUnmarshaler(&reflectclient.JsonUnmarshaler{}).
		Build()
	if err != nil {
		return nil, err
	}

	client.Init(service)

	return service, nil
}
