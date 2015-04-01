
package main

import (
  "fmt"
)

func main() {
  gh := Github{}
  fmt.Println(gh.Retrieve())

  updatedRepos []string{}
  createdRepos []string{}

  var wg sync.WaitGroup
  var tasks = make(chan Repo)

  maxWorkers := len(repos)

  if maxWorkers > 5 {
    maxWorkers = 5
  }

  fmt.Println("Spinning up", maxWorkers, "workers")

  for i := 0; i < maxWorkers; i++ {
   wg.Add(1)
   go func(wg *sync.WaitGroup) {
     for task := range tasks {
       _, err := os.Stat(task.LocalPath)

       if err == nil {
          updatedRepos = append(updatedRepos, task.RepoName)
          Pull(&task)
          fmt.Println("Pull complete for", task.RepoName)
        } else {
          createdRepos = append(createdRepos, task.RepoName)
          Clone(&task)
          fmt.Println("Cloning complete for", task.RepoName)
        }

     }

      wg.Done()
   }(&wg)
  }

  for i := 0; i < len(repos); i++ {
   if len(repos[i].IntranetUserID) > 0 {
     tasks <- repos[i]
   }
  }


}
