package utils

import (
	"fmt"
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateId() string{
	id, err := gonanoid.New(16)

	if(err != nil){
		log.Printf("Gonanoid error occured: %v", err)
	}

	return id

}


func StringToBool(str string) (*bool, error) {

  falseValue := false
  trueValue := true

  switch str {
  case "0":
    return &falseValue, nil // Return pointer to false
  case "1":
    return &trueValue, nil // Return pointer to true
  case "":
    return nil, nil // Return nil for empty string
  default:
    return nil, fmt.Errorf("invalid boolean string: %s", str)
  }
}