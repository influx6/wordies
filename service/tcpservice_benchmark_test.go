package service_test

import (
	"bufio"
	"context"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/icrowley/fake"
	"github.com/influx6/wordies/service"
)

func BenchmarkTCPServiceWithConstantMessage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := service.NewLetterCounter()
	words := service.NewWordCounter()

	var addr = "localhost:6559"
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan chan []string, 0)
	var waiter sync.WaitGroup

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
			for _, word := range <-job {
				letters.Compute(word)
				words.Compute(word)
			}
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

func BenchmarkTCPServiceWithVariableMessage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := service.NewLetterCounter()
	words := service.NewWordCounter()

	var addr = "localhost:4559"
	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan chan []string, 0)

	var waiter sync.WaitGroup

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
			for _, word := range <-job {
				letters.Compute(word)
				words.Compute(word)
			}
		}
	}()

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriterSize(conn, 32500)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		writer.WriteString(fake.Sentences())
		writer.Write(ending)
	}

	writer.Flush()
	b.StopTimer()

	conn.Close()
	cancel()
	waiter.Wait()
}

func BenchmarkTCPServiceWithUpdateCall(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := service.NewLetterCounter()
	words := service.NewWordCounter()
	top5 := new(service.Top5WordLetterStat)

	var waiter sync.WaitGroup

	var addr = "localhost:5567"

	jobs := make(chan chan []string, 0)
	ctx, cancel := context.WithCancel(context.Background())

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
			for _, word := range <-job {
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
