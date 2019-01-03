package main
//https://stackoverflow.com/questions/28682439/go-parse-yaml-file
import (
    "fmt"
    "io/ioutil"
    "path/filepath"

    "gopkg.in/yaml.v2"
)

type Config struct{
    Firewall_network_rules map[string]Options
}

type Options struct{
    Src string
    Dst string
}

func main(){
    fmt.Println("this will parse yaml soon")
    filename, _ := filepath.Abs("./fruits.yml")
    yamlFile, err := ioutil.ReadFile(filename)
    fmt.Println(yamlFile)
    fmt.Println(err)

    if err != nil {
        panic(err)
    }

    var config Config
    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Value: %#v\n", config.Firewall_network_rules)

    keys := make([]string, 0, )
}