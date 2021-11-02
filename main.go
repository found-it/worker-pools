package main

import (
	"github.com/found-it/worker-pools/worker/pool"
)

func main() {

	filler := []string{
		"thing0",
		"thing1",
		"thing2",
		"thing3",
		"thing4",
		"thing5",
		"thing6",
		"thing7",
		"thing8",
		"thing9",
	}

	pool.NewWorkerPool(3).WorkOn(&filler)
}
