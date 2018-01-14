package service

import (
	"bufio"
	"context"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/icrowley/fake"
)

var addr = "localhost:64559"

type goRecieveNotificaitons struct {
	top5   *Top5WordLetterStat
	waiter sync.WaitGroup
}

func (gh goRecieveNotificaitons) Update(fr FreshStat) {
	gh.waiter.Add(1)
	defer gh.waiter.Done()
	gh.top5.Update(fr)
}

func BenchmarkTCPServiceWithConstantMessage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := NewLetterCounter()
	words := NewWordCounter()

	pool := NewWorkerPool(300, 5*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	var waiter sync.WaitGroup

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		TCPService(ctx, false, addr, pool, letters, words, nil)
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

	pool.Stop()
	waiter.Wait()
}

func BenchmarkTCPServiceWithVariableMessage(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := NewLetterCounter()
	words := NewWordCounter()

	pool := NewWorkerPool(300, 5*time.Second)

	ctx, cancel := context.WithCancel(context.Background())

	var waiter sync.WaitGroup

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		TCPService(ctx, false, addr, pool, letters, words, nil)
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

	pool.Stop()
	waiter.Wait()
}

func BenchmarkTCPServiceWithUpdateCall(b *testing.B) {
	b.StopTimer()
	b.ReportAllocs()

	letters := NewLetterCounter()
	words := NewWordCounter()
	top5 := new(Top5WordLetterStat)

	var doneNotifier goRecieveNotificaitons
	doneNotifier.top5 = top5

	var waiter sync.WaitGroup

	pool := NewWorkerPool(100, 5*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		TCPService(ctx, false, addr, pool, letters, words, doneNotifier)
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
	pool.Stop()
	doneNotifier.waiter.Wait()
	waiter.Wait()
}
