package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"

)


func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// faire une fonction qui prendu ne liste de string, iterer dessus pour appeler hello

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}



func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome",
		"Great to see you %v !",
		"Hail king %v",
	}
	return formats[rand.Intn(len(formats))]

}