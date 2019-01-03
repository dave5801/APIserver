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
}

type NameAndEmail struct{
    Name string
    Email string
}

/*
type Options struct{
    Src string
    Dst string
}*/

func main(){
    fmt.Println("this will parse yaml soon")
  
    filename, _ := filepath.Abs("./metadata/test1.yml")
     
    yamlFile, err := ioutil.ReadFile(filename)
    fmt.Println(yamlFile)

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
    fmt.Printf("Value: %#v\n", config.Maintainers)

}