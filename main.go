package main

import (
	"foodpandagowk4/errhand"
	"foodpandagowk4/panichand"
	"log"
)

func main() {

	errhand.Run()

	wantFail := false
	i, block1, block1FlagOld := panichand.Run(wantFail)

	if i != 0 {
		log.Panicf("[panichand returns] hmm please check design.  %d", i)
	}

	if block1FlagOld != 0 {
		log.Panicf("[panichand returns] hmm please check design on block1.  %d", block1FlagOld)
	}

	if block1 == nil {
		log.Panicf("[panichand returns] hmm please check design on block1.")

	}

	if block1.Flag != 1 {
		log.Panicf("[panichand returns] hmm please check design on block1.  %d", block1.Flag)
	}

	wantFail = true
	i, block1, block1FlagOld = panichand.Run(wantFail)

	if i != 0 {
		log.Panicf("[panichand returns] hmm please check design.  %d", i)
	}

	if block1FlagOld != 0 {
		log.Panicf("[panichand returns] hmm please check design on block1.  %d", block1FlagOld)
	}

	if block1 != nil {
		log.Panicf("[panichand returns] hmm please check design on block1.")

	}
}
