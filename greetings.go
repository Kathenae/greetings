package greetings

import "fmt"
import "math/rand"
import "errors"

// Hello functions returns a greeting for the name of the person.
func Hello(name string) (string, error) {
  // if no name was provided, return an error with a message
  if name == "" {
    return "", errors.New("empty name")
  }

  // Return a greeting that embeds the name in a message.
  var message string
  message = fmt.Sprintf(randomFormat(), name)
  return message, nil
}

func randomFormat() string {
  // A slice of message formats
  formats := []string {
    "Hi, %v. Welcome!",
    "Great to see you %v",
    "Hail, %v! Well met!",
  }

  // Return a randomly selected message format by specifying
  // a random index for the slices of formats.
  return formats[rand.Intn(len(formats))]
}
