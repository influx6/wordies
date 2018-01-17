package service_test

import (
	"bufio"
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/influx6/faux/tests"
	"github.com/influx6/wordies/service"
)

var (
	ending       = []byte("\r\n")
	top5Expected = service.Top5Stat{
		Count:   26,
		Letters: []string{"E", "A", "M", "T", "I"},
		Words:   []string{"me", "and", "miss.", "greg", "i"},
	}
)

func TestTCPService(t *testing.T) {
	letters := service.NewLetterCounter()
	words := service.NewWordCounter()
	top5 := new(service.Top5WordLetterStat)

	jobs := make(chan chan string, 1)

	addr := "localhost:4559"
	var waiter sync.WaitGroup
	waiter.Add(1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go func() {
		defer waiter.Done()
		service.TCPService(ctx, false, addr, jobs)
	}()

	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		cancel()
		tests.FailedWithError(err, "Should have successfully connected to service.")
	}
	tests.Passed("Should have successfully connected to service.")

	writer := bufio.NewWriterSize(conn, len(basicSentence))
	writer.WriteString(basicSentence)
	writer.Write(ending)
	if err := writer.Flush(); err != nil {
		tests.FailedWithError(err, "Should have successfully written data to tcp service.")
	}
	tests.Passed("Should have successfully written data to tcp service.")

	conn.Close()
	cancel()

	job := <-jobs
	for word := range job {
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

	stat := top5.Stat()
	if stat.Count != top5Expected.Count {
		tests.Info("Received: %d", stat.Count)
		tests.Info("Expected: %d", top5Expected.Count)
		tests.Failed("Should have gotten exact count stats from http service")
	}
	tests.Passed("Should have gotten exact count stats from http service")

	if !compareSlices(stat.Letters, top5Expected.Letters) {
		tests.Info("Received: %+q", stat.Letters)
		tests.Info("Expected: %+q", top5Expected.Letters)
		tests.Failed("Should have gotten exact letters stats from http service")
	}
	tests.Passed("Should have gotten exact letters stats from http service")

	if !compareSlices(stat.Words, top5Expected.Words) {
		tests.Info("Received: %+q", stat.Words)
		tests.Info("Expected: %+q", top5Expected.Words)
		tests.Failed("Should have gotten exact words stats from http service")
	}
	tests.Passed("Should have gotten exact words stats from http service")

	waiter.Wait()
}
