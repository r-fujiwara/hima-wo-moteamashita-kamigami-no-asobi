package main

import(
	"log"
	"net/http"
	"golang.org/x/net/context"
	"github.com/guregu/kami"
)

func greet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}

func main() {
	ctx := context.Background()
	kami.Context = ctx

	kami.Get("/", greet)
	go kami.Serve()
}
