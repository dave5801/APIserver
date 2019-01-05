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

/*
func validEmail(email string) bool{
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(email)
}

func isValidUrl(Url string) bool {
    _, err := url.ParseRequestURI(Url)
    if err != nil {
        return false
    } else {
        return true
    }
}*/

type validator interface{
    validateEmail() bool
    validateURL() bool
}

type YmlFile struct {
    Email string
    Url string
}

func (ymlFile YmlFile) validateEmail() bool{
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

func main() {
    v := YmlFile{"apptwohotmail.com", "http://www.golangcode.com"}
    fmt.Printf(v.Email)
    /*
    email := "apptwohotmail.com"

    fmt.Printf("\nEmail: %v :%v\n", email, validEmail(email))
    fmt.Println(isValidUrl("http://www.golangcode.com"))

    // = false
    fmt.Println(isValidUrl("golangcode.com"))

    // = false
    fmt.Println(isValidUrl(""))*/

 
}