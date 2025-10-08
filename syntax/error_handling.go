package syntax

import "fmt"

func ShouldBeError() error {
	return fmt.Errorf("this is an error")
}
