// Package aggreagtemyqiita is logic package
package aggreagtemyqiita

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	filename = "app.properties"
)

// Params is a command argument.
type Params struct {
	UserID string
	Token  string
}

// GetParameters get the default value from app.properties.
func GetParameters(params *Params) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "=")

		if len(s) != 2 {
			return err
		}

		switch s[0] {
		case "user":
			params.UserID = s[1]
		case "token":
			params.Token = s[1]
		default:
		}
	}

	return nil
}

// WriteParameters write parameters to app.properties.
func WriteParameters(params *Params) error {
	if err := os.Remove(filename); err != nil {
		return err
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "user=%s\n", params.UserID)
	fmt.Fprintf(file, "token=%s\n", params.Token)

	return nil
}
