package testmod

import (
	"errors"
	"fmt"
)

//Hi returns a friendly greeting
func Hi(name string, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi %s!", name)
	case "pt":
		return fmt.Sprintf("Oi  %s!", name)
	case "es":
		return fmt.Sprintf("Hola %s!", name)
	case "fr":
		return fmt.Sprintf("Bonjour %s!", name)
	default:
		return "", errors.New("unknown Language")
	}
}
