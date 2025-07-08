package utils

import "fmt"

func StringifyError(err error) string {
	if err == nil {
		return "nil"
	}
	return fmt.Sprintf("error(%q)", err.Error())
}
