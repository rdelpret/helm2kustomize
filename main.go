package main

import (
	"os/exec"
	"fmt"
	"log"
	"gopkg.in/yaml.v2"
	"os"
)



type Config struct {
    test string `yaml`
}


func addHelmRepo(repoName string, repoURL string) {
    out, err := exec.Command("helm", "repo", "add", repoName, repoURL).Output()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Helm: %v", string(out))

}

func wgetValuesFile(url string) {
    out, err := exec.Command("wget", "-N", url).Output()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pulling down values.yaml file %v", string(out))
}

func main() {
  // pull in yaml config


  fmt.Printf("%v", cfg)


  // Make sure the helm repo is added
  //addHelmRepo(c.repoName, c.repoURL)

  // wget the values file
  //wgetValuesFile(c.Values)

  // Template out the helm chart

}
