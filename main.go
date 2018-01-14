package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"strings"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/influx6/faux/flags"
	"github.com/influx6/wordies/service"
)

func main() {
	flags.Run("wordies",
		flags.Command{
			Name:      "stats",
			Action:    stats,
			ShortDesc: "Retrieve current stats from wordies http service.",
			Desc:      "Sends a http requests to retrieve latest stats from the wordies http service.",
			Flags: []flags.Flag{
				&flags.StringFlag{
					Name:    "httpAddr",
					Desc:    "sets the address for the http service",
					Default: "localhost:8080",
				},
				&flags.DurationFlag{
					Name:    "timeout",
					Desc:    "sets the maximum duration allowed to wait to connect to service",
					Default: time.Second * 2,
				},
			},
		},
		flags.Command{
			Name:      "send",
			Action:    send,
			ShortDesc: "Send a string of sentences to the tcp server",
			Desc:      "send provides a command that sends provided sentences to wordies tcp service if running",
			Flags: []flags.Flag{
				&flags.StringFlag{
					Name:    "tcpAddr",
					Desc:    "sets the address for the tcp level service",
					Default: "localhost:5555",
				},
				&flags.DurationFlag{
					Name:    "timeout",
					Desc:    "sets the maximum duration allowed to wait to connect to service",
					Default: time.Second * 2,
				},
			},
		},
		flags.Command{
			Name:      "serve",
			Action:    serve,
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
		})
}

func stats(ctx flags.Context) error {
	timeout, _ := ctx.GetDuration("timeout")
	httpAddr, _ := ctx.GetString("httpAddr")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/stats", httpAddr), nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: timeout}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Request Failed: recieved %d", res.Status)
	}

	if !strings.Contains(res.Header.Get("Content-Type"), "application/json") {
		return fmt.Errorf("HTTP Request Failed: invalid response type")
	}

	var bu bytes.Buffer
	io.Copy(&bu, res.Body)
	fmt.Println(bu.String())
	return nil
}

func send(ctx flags.Context) error {
	if len(ctx.Args()) == 0 {
		return errors.New("a sentence must atleast be provided")
	}

	timeout, _ := ctx.GetDuration("timeout")
	tcpAddr, _ := ctx.GetString("tcpAddr")

	conn, err := net.DialTimeout("tcp", tcpAddr, timeout)
	if err != nil {
		return nil
	}

	sentences := strings.Join(ctx.Args(), " ") + "\r\n"
	writer := bufio.NewWriterSize(conn, len(sentences))
	writer.WriteString(sentences)
	return writer.Flush()
}

func serve(ctx flags.Context) error {
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
}
