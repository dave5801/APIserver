package main
 
import (
    "fmt"
    "regexp"
    "net/url"
)
 /*
    Need
        X - emails
        X - urls
        missing fields
 */

type validator interface{
    validateEmail() bool
    validateURL() bool
}

//MAYBE IMPORT METADATACONFIG HERE

type YmlFile struct {
    Email string
    Url string
}

func (ymlFile YmlFile) validateEmail() bool{
    //regex from http://www.golangprograms.com/regular-expression-to-validate-email-address.html
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(ymlFile.Email)
}

func (ymlFile YmlFile) validateURL() bool{
     _, err := url.ParseRequestURI(ymlFile.Url)
    if err != nil {
        return false
    } else {
        return true
    }
}

//func validEmail(email string) bool{}
func Validate(v validator){
    fmt.Println(v)
    fmt.Println(v.validateEmail())
    fmt.Println(v.validateURL())
}

func main() {
    v := YmlFile{"apptwohotmail.com", "http://www.golangcode.com"}
    Validate(v)
 
}