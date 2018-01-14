package service

import (
	"context"
	"fmt"
	"log"

	"bufio"
	"net"
	"net/textproto"
	"sync"
)

var (
	maxDataSize = 1024 * 1024 // 1mb
)

// TCPService initializes a simple tcp based connection which listens for requests form clients, which
// then gets processed by both letters and word counters in provided worker pool.
func TCPService(ctx context.Context, verbose bool, addr string, pool WorkerPool, letters *LetterCounter, words *WordCounter, nf StatMetric) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		defer listener.Close()
		<-ctx.Done()
	}()

	if verbose {
		log.Printf("TCP Service listening on %+q\n", addr)
	}

	var waiter sync.WaitGroup

	for {
		newConn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(*net.OpError); ok && !ne.Temporary() {
				break
			}
			continue
		}

		waiter.Add(1)
		go func(cn net.Conn) {
			defer waiter.Done()
			if err := handleClientConnection(cn, pool, letters, words, nf); err != nil {
				if verbose {
					log.Printf("client connection closed: %+s\n", cn.RemoteAddr())
				}
			}
		}(newConn)
	}

	waiter.Wait()
	return nil
}

// handleClientConnection handles the internal logic necessary to listen to messages provided by connected clients.
// We will read on a line by line basis, that is all text must have the \r\n ending attached.
func handleClientConnection(conn net.Conn, pool WorkerPool, lt *LetterCounter, wd *WordCounter, nf StatMetric) error {
	reader := bufio.NewReaderSize(conn, maxDataSize)
	defer reader.Reset(nil)

	textReader := textproto.NewReader(reader)

	for {
		data, err := textReader.ReadLine()
		if err != nil {
			//if err == io.EOF {
			//fmt.Printf("Read Err on %+q: %+q\n", conn.RemoteAddr(), err)
			//}

			return err
		}

		// transform to string, since we are expecting sentences
		// also we need to keep this data uncorruptable, preferable to
		// turn into string here than copy into another slice.
		func(sentence string) {
			if verbose {
				fmt.Println("Recieved: ", sentence)
			}
			pool.Add(func() {
				for _, word := range lexSentence(data) {
					lt.Compute(word)
					wd.Compute(word)
				}

				// Send lastest stat updates if required.
				if nf != nil {
					letters, letterCount := lt.Stat()
					words, wordCount := wd.Stat()
					nf.Update(FreshStat{
						Words:        words,
						Letters:      letters,
						TotalWords:   wordCount,
						TotalLetters: letterCount,
					})
				}
			})
		}(string(data))
	}

	return nil
}
