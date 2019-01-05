package main
//https://stackoverflow.com/questions/28682439/go-parse-yaml-file
import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "regexp" //note - for validation
   // "net/url" //note - for validation
    "metaDataConfig"

    "gopkg.in/yaml.v2"
)

//NOTE - this will be moved to another file
type validator interface{
    validateMaintainerEmail() bool
   // validateURL() bool
}

//structure for yml file
/*
type MetaDataConfig struct{
    Title string
    Version string
    Maintainers []map[string]string
    Company string
    Website string
    Source string
    License string
    Description string
}*/

func (metaDataConfig MetaDataConfig) validateMaintainerEmail() bool{
    //regex from http://www.golangprograms.com/regular-expression-to-validate-email-address.html
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(metaDataConfig.Maintainers[0]["email"])
}

func Validate(v validator) string{
 
    if v.validateMaintainerEmail() == false{
        return "MetaData is invalid"
    }else{
        return "MetaData is valid"
    }
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
    isMetaDataValid := Validate(metaDataConfig)
    fmt.Println(isMetaDataValid)

    /*
    fmt.Printf("Title: %#v\n", metaDataConfig.Title)
    fmt.Printf("Version: %#v\n", metaDataConfig.Version)
    fmt.Printf("Maintainers:\n")
    fmt.Printf("    Email: %#v\n", metaDataConfig.Maintainers[0]["email"])
    fmt.Printf("    Name:  %#v\n", metaDataConfig.Maintainers[0]["name"])
    fmt.Printf("Company: %#v\n", metaDataConfig.Company)
    fmt.Printf("Website: %#v\n", metaDataConfig.Website)
    fmt.Printf("Source: %#v\n", metaDataConfig.Source)
    fmt.Printf("License: %#v\n", metaDataConfig.License)
    fmt.Printf("Description: %#v\n", metaDataConfig.Description)*/
    
}