package cmd

import "fmt"

func PanicIfError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: \n %s", msg, err))
	}
}
