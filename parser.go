package main
//https://stackoverflow.com/questions/28682439/go-parse-yaml-file
import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)


type Config struct{
    Title string
    Version string
    Maintainers []map[string]string
    Company string
    Website string
    Source string
    License string
    Description string
}

func main(){
  
    filename, _ := filepath.Abs("./metadata/test1.yml")
     
    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Title: %#v\n", config.Title)
    fmt.Printf("Version: %#v\n", config.Version)
    fmt.Printf("Maintainers:\n")
    fmt.Printf("    Email: %#v\n", config.Maintainers[0]["email"])
    fmt.Printf("    Name:  %#v\n", config.Maintainers[0]["name"])
    fmt.Printf("Company: %#v\n", config.Company)
    fmt.Printf("Website: %#v\n", config.Website)
    fmt.Printf("Source: %#v\n", config.Source)
    fmt.Printf("License: %#v\n", config.License)
    fmt.Printf("Description: %#v\n", config.Description)

}