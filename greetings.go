package greetings

import "fmt"

// Hello functions returns a greeting for the name of the person.
func Hello(name string) string {
  // Return a greeting that embeds the name in a message.
  var message string
  message = fmt.Sprintf("Hi, %v. Welcome!", name)
  return message
}
