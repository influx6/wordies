package service

import (
	"context"
	"fmt"
	"io"
	"log"

	"bufio"
	"net"
	"net/textproto"
	"sync"

	"github.com/influx6/wordies/internal"
)

var (
	maxDataSize = 1024 * 1024 // 1mb
)

// TCPService initializes a simple tcp based server, which listens for requests from clients, which
// then processed into words and send into provided jobs channel. It closes said channel when
// the server get's closed.
func TCPService(ctx context.Context, verbose bool, addr string, jobs chan chan []string) error {
	defer close(jobs)

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
			if err := handleClientConnection(cn, verbose, jobs); err != nil {
				if verbose {
					log.Printf("client connection closed: %+s\n", cn.RemoteAddr())
				}
			}
		}(newConn)
	}

	waiter.Wait()
	return nil
}

// handleClientConnection handles the internal logic necessary to listen to messages from a client net.Conn.
// We will read on a line by line basis, that is all text must have the \r\n ending attached.
// It processed received sentences into words and feeds it into provided job channel.
func handleClientConnection(conn net.Conn, verbose bool, jobs chan chan []string) error {
	reader := bufio.NewReaderSize(conn, maxDataSize)
	defer reader.Reset(nil)

	textReader := textproto.NewReader(reader)

	for {
		data, err := textReader.ReadLine()
		if err != nil {
			if err == io.EOF && verbose {
				fmt.Printf("Read Err on %+q: %+q\n", conn.RemoteAddr(), err)
			}

			return err
		}

		sentence := string(data)
		if verbose {
			fmt.Printf("Recieved: %+q\n", sentence)
		}

		job := make(chan []string, 1)
		job <- internal.LexSentence(sentence)
		jobs <- job
	}

	return nil
}
