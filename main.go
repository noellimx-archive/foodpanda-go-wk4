package main

import (
	"fmt"
	"foodpandagowk4/errhand"
	"foodpandagowk4/panichand"
	"log"
)

func main() {

	errhand.Run()
	i, block1, block1FlagOld := panichand.Run()

	if i != 0 {
		log.Panic(fmt.Sprintf("[panichand returns] hmm please check design.  %d", i))
	}

	if block1.Flag != 1 {
		log.Panic(fmt.Sprintf("[panichand returns] hmm please check design on block1.  %d", i))

	}

	if block1FlagOld != 0 {
		log.Panic(fmt.Sprintf("[panichand returns] hmm please check design on block1.  %d", block1FlagOld))
	}
}
