package main

import "fmt"

var WorkerQueue chan Worker

func StartDispatcher(nworkers int) {
	// First, initialize the channel we are going to but the worker's work channels into.
	WorkerQueue = make(chan Worker, nworkers)

	// Now, create all of our workers
	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWokrer(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				fmt.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					fmt.Printf("Dispatching work request to worker%d\n", worker.ID)
					worker.Work <- work
				}()
			}
		}
	}()
}
