package greetings

import "fmt"

// Hello returns a greeting for the named person
func Hello() {
    message := fmt.Sprintf("Hi %v. Welcome!", name)
    return message
}

// Alternative way to write this function
// func Hello () {
//     var message string
//     message = fmt.Sprintf("Hi %v. Welcome!", name)
// }
