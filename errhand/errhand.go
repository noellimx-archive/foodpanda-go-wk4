package errhand

import (
	"errors"
	"fmt"
	"log"
)

type customError struct {
	prepend int
}

func (e customError) Error() string {
	e.prepend += 1
	return fmt.Sprintf("sure fail error 2, Errors should be stateless, that is passed by-copy. %d", e.prepend)
}

var errSureFail error = errors.New("sure fail error")

var errSureFail2 customError

func SureFail() (int, error) {

	return 1, errSureFail
}

func SureFail2() (int, error) {

	return 1, errSureFail2
}

func IdiomaticHandling() {
	_, err := SureFail()

	if err != nil {
		a := 0
		a += 0
	}
}

func IdiomaticHandling2() {
	if _, err := SureFail2(); err != nil {
		a := 0
		a += 0
		log.Printf("[IdiomaticHandling2] %s", err.Error())

	}
}

func Run() {
	IdiomaticHandling()
	i := 0

	for i < 10 {
		IdiomaticHandling2()
		i += 1
	}

}
