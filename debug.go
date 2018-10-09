package resteater

import "fmt"

func debug(msg ...interface{}) (int, error) {
	return fmt.Println("DEBUG:", msg)
}
