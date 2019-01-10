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

//functions for validating the yml files
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

//return true if maintainer email matches the regular expression
func (metaDataConfig MetaDataConfig) validateMaintainerEmail() bool{
    //CITATION: regex from http://www.golangprograms.com/regular-expression-to-validate-email-address.html
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(metaDataConfig.Maintainers[0]["email"])
    
}

//return true if website and source urls are valid
func (metaDataConfig MetaDataConfig) validateURL() bool{

    _, webSiteErr := url.ParseRequestURI(metaDataConfig.Website)
    _, sourceErr := url.ParseRequestURI(metaDataConfig.Source)
    
    if webSiteErr != nil || sourceErr != nil {
        return false
    } else {
        return true
    }
}

//return true if no empty fields are missing from .yml
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

//returns true if .yml passes checks for emails, urls, or missing fields
func Validate(v validator) bool{

    if v.validateFields()==false || v.validateMaintainerEmail()==false || v.validateURL() == false{
        return false
    }else{
        return true
    }    
}

//parse and validate .yml files in a directory
func returnValidConfigFiles(configFilDirPath string) []MetaDataConfig {

    //parse file firectory
    files, err := ioutil.ReadDir(configFilDirPath)
    if err != nil {
        log.Fatal(err)
    }

    var arrayOfMetaDataConfigFiles[] MetaDataConfig 

    //parse individual .yml files
    for _, f := range files {

            yamlFile, err := ioutil.ReadFile(configFilDirPath + f.Name())
            if err != nil {
                panic(err)
            }
            var metaDataConfig MetaDataConfig
    
            err = yaml.Unmarshal(yamlFile, &metaDataConfig)

            //append valid .yml files to an array
            if Validate(metaDataConfig) == true{
                arrayOfMetaDataConfigFiles = append(arrayOfMetaDataConfigFiles, metaDataConfig)
            }    
    }

    return arrayOfMetaDataConfigFiles
}

//GET Request for all validated .yml files
func GetConfigs(w http.ResponseWriter, r *http.Request) {

    //create list of valid .yml files
    arrayOfValidMetaDataConfigFiles := returnValidConfigFiles("./metadata/")
    
    //convert the list of .yml files to JSON
    json.NewEncoder(w).Encode(arrayOfValidMetaDataConfigFiles)
}

//main method
func main(){

    //new server
    router := mux.NewRouter()

    //assign url path to get request
    router.HandleFunc("/configs", GetConfigs  ).Methods("GET")

    fmt.Println("Server is Running...")
    log.Fatal(http.ListenAndServe(":8000", router))

    
}