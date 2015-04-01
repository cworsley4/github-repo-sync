package main

import (
  "encoding/json"
  _ "fmt"
  "io/ioutil"
  "net/http"
)

type Github struct{}

type Repo struct {
  SshUrl   string `json:"ssh_url"`
  CloneUrl string `json:"clone_url"`
  GitUrl   string `json:"git_url"`
  Fork     bool   `json:"fork"`
}

func (g Github) Retrieve() []Repo {
  api := "https://api.github.com/users/cworsley4/repos"
  resp, _ := http.Get(api)
  body, _ := ioutil.ReadAll(resp.Body)

  var repos []Repo

  err := json.Unmarshal(body, &repos)

  if err != nil {
    panic(err)
  }

  return repos
}
