package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/influx6/faux/flags"
	"github.com/influx6/wordies/service"
)

func main() {
	flags.Run("wordies", flags.Command{
		Name:      "serve",
		ShortDesc: "Serve tcp and http word frequency service.",
		Desc:      "serve starts the tcp and http components of the natural language word frequency service.",
		Flags: []flags.Flag{
			&flags.IntFlag{
				Name:    "workers",
				Desc:    "sets the maximum workers for background language processing requests",
				Default: 1000,
			},
			&flags.DurationFlag{
				Name:    "workers.timeout",
				Desc:    "sets the maximum duration allowed for a worker to be idle",
				Default: time.Second * 30,
			},
			&flags.StringFlag{
				Name:    "httpAddr",
				Desc:    "sets the address for the tcp level service",
				Default: "localhost:8080",
			},
			&flags.StringFlag{
				Name:    "tcpAddr",
				Desc:    "sets the address for the tcp level service",
				Default: "localhost:5555",
			},
		},
		Action: func(ctx flags.Context) error {
			workers, _ := ctx.GetInt("workers")
			timeout, _ := ctx.GetDuration("workers.timeout")
			pools := service.NewWorkerPool(workers, timeout)

			words := service.NewWordCounter()
			letters := service.NewLetterCounter()
			top5 := new(service.Top5WordLetterStat)

			router := mux.NewRouter()
			router.Path("/stats").HandlerFunc(service.Top5Stats(top5)).Methods("GET")

			httpAddr, _ := ctx.GetString("httpAddr")
			var httpServer http.Server
			httpServer.Addr = httpAddr
			httpServer.Handler = router
			httpServer.SetKeepAlivesEnabled(true)

			defer httpServer.Shutdown(ctx)

			go func() {
				log.Printf("HTTP Service listening on %+q\n", httpAddr)

				if err := httpServer.ListenAndServe(); err != nil {
					log.Fatal(err)
				}
			}()

			tcpAddr, _ := ctx.GetString("tcpAddr")
			return service.TCPService(ctx, tcpAddr, pools, letters, words, top5)
		},
	})
}
