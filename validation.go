package main
 
import (
    "fmt"
    "regexp"
)
 /*
    Need
        emails
        urls
        missing fields
 */


func main() {
    
    str1 := "apptwohotmail.com"
    str2 := "email@gmail.com"
   

 
    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
 
    fmt.Printf("Pattern: %v\n", re.String()) // print pattern   
    fmt.Printf("\nEmail: %v :%v\n", str1, re.MatchString(str1))
    fmt.Printf("Email: %v :%v\n", str2, re.MatchString(str2))
 
}