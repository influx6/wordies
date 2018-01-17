package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/influx6/wordies/internal"
	"github.com/influx6/wordies/service"
)

var (
	ending        = []byte("\r\n")
	basicSentence = "Miss. Greg left with me and we met at the cafe and went home. After which I got me some coffee; but Miss. Greg appeared and kissed me before I knew we were married"
	basicWords    = internal.LexSentence(basicSentence)
)

func BenchmarkTCPServiceWithWorkerPool(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := service.NewLetterCounter()
	words := service.NewWordCounter()

	var addr = "localhost:6559"
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan chan []string, 0)
	var waiter sync.WaitGroup

	pool := internal.NewWorkerPool(300, 30*time.Second)
	defer pool.Stop()

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		service.TCPService(ctx, false, addr, jobs)
	}()

	// drain channel
	waiter.Add(1)
	go func() {
		defer waiter.Done()
		for job := range jobs {
			func(incoming chan []string) {
				pool.Add(func() {
					for _, word := range <-incoming {
						letters.Compute(word)
						words.Compute(word)
					}
				})
			}(job)
		}
	}()

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriterSize(conn, len(basicSentence)+2)

	b.SetBytes(int64(len(basicSentence) + 2))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		writer.WriteString(basicSentence)
		writer.Write(ending)
	}

	writer.Flush()
	b.StopTimer()

	conn.Close()
	cancel()
	waiter.Wait()
}

func BenchmarkTCPServiceWithUpdatesWithWorkerPool(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := service.NewLetterCounter()
	words := service.NewWordCounter()
	top5 := new(service.Top5WordLetterStat)

	var addr = "localhost:6559"
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan chan []string, 0)
	var waiter sync.WaitGroup

	pool := internal.NewWorkerPool(300, 30*time.Second)
	defer pool.Stop()

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		service.TCPService(ctx, false, addr, jobs)
	}()

	// drain channel
	waiter.Add(1)
	go func() {
		defer waiter.Done()
		for job := range jobs {
			func(incoming chan []string) {
				pool.Add(func() {
					for _, word := range <-incoming {
						letters.Compute(word)
						words.Compute(word)
					}

					ltstat, lttotal := letters.Stat()
					wdstat, wdtotal := words.Stat()
					top5.Update(service.FreshStat{
						Words:        wdstat,
						Letters:      ltstat,
						TotalWords:   wdtotal,
						TotalLetters: lttotal,
					})
				})
			}(job)
		}
	}()

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriterSize(conn, len(basicSentence)+2)

	b.SetBytes(int64(len(basicSentence) + 2))
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		writer.WriteString(basicSentence)
		writer.Write(ending)
	}

	writer.Flush()
	b.StopTimer()

	conn.Close()
	cancel()
	waiter.Wait()
}
