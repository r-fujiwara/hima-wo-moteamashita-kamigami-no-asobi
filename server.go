package main

import(
	"log"
	"fmt"
	"time"
	"net/http"
	"golang.org/x/net/context"
	"github.com/guregu/kami"
)

var WorkQueue = make(chan WorkRequest, 100)

func greet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	
	// Check to make sure the delay is anaywhere from 1 to 10 seconds.
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	if delay.Seconds() < 1 || delay.Seconds() > 10 {
		http.Error(w, "The delay must be beteween 1 and 10 seconds, inclusively.", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")

	// Now,  we take the delay, and the persion's name, and make a WorkRequest out of them
	work := WorkRequest{Name: name, Delay: delay}

	// Push the work onto the queue.
	WorkQueue <- work
	fmt.Println("Work request queued")

	// And let the user know their work request was created.
	w.WriteHeader(http.StatusCreated)
	return
}

func main() {
	ctx := context.Background()
	kami.Context = ctx

	log.Println("start dispatcher...")
	StartDispatcher(2048)

	kami.Get("/work", greet)
	kami.Serve()
}
