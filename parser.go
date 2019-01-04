package main
//https://stackoverflow.com/questions/28682439/go-parse-yaml-file
import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

//structure for yml file
type MetaDataConfig struct{
    Title string
    Version string
    Maintainers []map[string]string
    Company string
    Website string
    Source string
    License string
    Description string
}

func parseMetaDataFromYML(filename string) MetaDataConfig{

    yamlFile, err := ioutil.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var metaDataConfig MetaDataConfig
    
    err = yaml.Unmarshal(yamlFile, &metaDataConfig)
    
    return metaDataConfig
}

func main(){
  
    filename, _ := filepath.Abs("./metadata/test1.yml")
    metaDataConfig := parseMetaDataFromYML(filename)

    fmt.Printf("Title: %#v\n", metaDataConfig.Title)
    fmt.Printf("Version: %#v\n", metaDataConfig.Version)
    fmt.Printf("Maintainers:\n")//NOTE -- Maintainers will likely need to be iterated upon
    fmt.Printf("    Email: %#v\n", metaDataConfig.Maintainers[1]["email"])
    fmt.Printf("    Name:  %#v\n", metaDataConfig.Maintainers[1]["name"])
    fmt.Printf("Company: %#v\n", metaDataConfig.Company)
    fmt.Printf("Website: %#v\n", metaDataConfig.Website)
    fmt.Printf("Source: %#v\n", metaDataConfig.Source)
    fmt.Printf("License: %#v\n", metaDataConfig.License)
    fmt.Printf("Description: %#v\n", metaDataConfig.Description)
    
}