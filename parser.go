package main

import (
    "fmt"
    "io/ioutil"
    "regexp" 
    "net/url" 
    "reflect"
    "log"
    "gopkg.in/yaml.v2"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

//NOTE - this will be moved to another file
type validator interface{
    validateMaintainerEmail() bool
    validateURL() bool
    validateFields() bool
}

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

//var people []Person
//var arrayOfValidMetaDataConfigFiles []MetaDataConfig

func (metaDataConfig MetaDataConfig) validateMaintainerEmail() bool{
    //CITATION: regex from http://www.golangprograms.com/regular-expression-to-validate-email-address.html
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(metaDataConfig.Maintainers[0]["email"])
    
}

func (metaDataConfig MetaDataConfig) validateURL() bool{

    _, webSiteErr := url.ParseRequestURI(metaDataConfig.Website)
    _, sourceErr := url.ParseRequestURI(metaDataConfig.Source)
    
    if webSiteErr != nil || sourceErr != nil {
        return false
    } else {
        return true
    }
}

func (metaDataConfig MetaDataConfig) validateFields() bool{

    copyConfig := reflect.ValueOf(metaDataConfig)
    ymlFieldValues := make([]interface{}, copyConfig.NumField())

    for i:= 0; i < copyConfig.NumField(); i++{
        ymlFieldValues[i] = copyConfig.Field(i).Interface()
        if copyConfig.Field(i).Interface() == ""{
            return false
        }
    }

    return true
}

func Validate(v validator) bool{

    if v.validateFields()==false || v.validateMaintainerEmail()==false || v.validateURL() == false{
        return false
    }else{
        return true
    }    
}

func returnValidConfigFiles(configFilDirPath string) []MetaDataConfig {

    files, err := ioutil.ReadDir(configFilDirPath)
    if err != nil {
        log.Fatal(err)
    }

    var arrayOfMetaDataConfigFiles[] MetaDataConfig 

    for _, f := range files {

            yamlFile, err := ioutil.ReadFile(configFilDirPath + f.Name())
            if err != nil {
                panic(err)
            }
            var metaDataConfig MetaDataConfig
    
            err = yaml.Unmarshal(yamlFile, &metaDataConfig)
            if Validate(metaDataConfig) == true{
                arrayOfMetaDataConfigFiles = append(arrayOfMetaDataConfigFiles, metaDataConfig)
            }    
    }

    return arrayOfMetaDataConfigFiles
}

func GetConfigs(w http.ResponseWriter, r *http.Request) {
    
    arrayOfValidMetaDataConfigFiles := returnValidConfigFiles("./metadata/")
    
    json.NewEncoder(w).Encode(arrayOfValidMetaDataConfigFiles)
}

func main(){

    router := mux.NewRouter()

    //arrayOfValidMetaDataConfigFiles := returnValidConfigFiles("./metadata/")
    //fmt.Println(arrayOfValidMetaDataConfigFiles)

    router.HandleFunc("/configs", GetConfigs).Methods("GET")

    fmt.Println("Server is Running...")
    log.Fatal(http.ListenAndServe(":8000", router))
    
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