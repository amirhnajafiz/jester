package pkg

import "fmt"

func WrapError(trace, message string, e error) error {
	return fmt.Errorf("[%s] %s:\n\t%w", trace, message, e)
}
