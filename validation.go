package main
 
import (
    "fmt"
    "regexp"
)
 /*
    Need
        X - emails
        urls
        missing fields
 */

func validEmail(email string) bool{
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return re.MatchString(email)
}

func main() {
    
    email := "apptwohotmail.com"

    fmt.Printf("\nEmail: %v :%v\n", email, validEmail(email))

 
}