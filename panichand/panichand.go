package panichand

import "log"

func Divide(a int, b int) int {
	res := a / b

	log.Printf("[Divide] %d / %d => %d", a, b, res)
	return a / b
}

func defereable(i string) {
	log.Printf("[defereable] %s", i)
}

type Block struct {
	Flag int
}

func Run() (int, *Block, int) {
	defer defereable("beta")

	i := 10
	for i > 1 {

		i -= 1
		Divide(10-i, i)
	}

	state := 0

	var block1 *Block = &Block{}

	defer func() {
		log.Printf("[inline-defer state init] primitive: %d pointer val: %d %p", state, block1.Flag, block1)

		state += 1
		block1.Flag += 1

	}()

	defer func() {
		log.Printf("[inline-defer state after] primitive: %d pointer val: %d %p", state, block1.Flag, block1)

	}()

	defer defereable("omega")

	defer defereable("charlie. should be executed last")

	return state, block1, block1.Flag

}
