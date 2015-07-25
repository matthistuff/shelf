package helpers
import (
	"os"
	"fmt"
)

func ErrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrExit(predicate bool, message string) {
	if predicate {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}