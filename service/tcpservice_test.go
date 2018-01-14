package service

import (
	"bufio"
	"context"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/influx6/faux/tests"
)

var (
	ending       = []byte("\r\n")
	top5Expected = Top5Stat{
		Count:   26,
		Letters: []string{"E", "A", "M", "T", "I"},
		Words:   []string{"me", "and", "miss.", "greg", "i"},
	}
)

type goHeadStat struct {
	top5 *Top5WordLetterStat
	done chan struct{}
}

func (gh goHeadStat) Update(fr FreshStat) {
	gh.top5.Update(fr)
	close(gh.done)
}

func TestTCPService(t *testing.T) {
	letters := NewLetterCounter()
	words := NewWordCounter()
	top5 := new(Top5WordLetterStat)

	var doneNotifier goHeadStat
	doneNotifier.top5 = top5
	doneNotifier.done = make(chan struct{}, 0)

	pool := NewWorkerPool(300, 3*time.Second)
	defer pool.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	addr := "localhost:4559"
	var waiter sync.WaitGroup

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		TCPService(ctx, false, addr, pool, letters, words, doneNotifier)
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

	<-doneNotifier.done

	conn.Close()
	cancel()

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
