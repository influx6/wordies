Wordies
---------
[![Go Report Card](https://goreportcard.com/badge/github.com/influx6/wordies)](https://goreportcard.com/report/github.com/influx6/wordies)
[![Travis Build Status](https://travis-ci.org/influx6/wordies.svg?branch=master)](https://travis-ci.org/influx6/wordies#)
[![CircleCI](https://circleci.com/gh/influx6/wordies.svg?style=svg)](https://circleci.com/gh/influx6/wordies)

Wordies showcases a simple CLI service that exposes a tcp server and a http server, where the tcp server exposed on port `5555` by default, receives strings of sentences as text ending in newline or `\r\n`, which it counts all letters and word frequencies.

It allows stats to be retrieved from the underline tcp service through the http server exposed over port `8080` by default.


## Install

```bash
go get  github.com/influx6/wordies
```

## Usage

The project installs the `wordies` CLI into your `$GOBIN` or `$GOPATH/bin` path, ensure to have this path exported, so binaries in there can be executed.

```bash
> wordies

Usage: wordies [flags] [command] 

⡿ COMMANDS:
	⠙ stats        Retrieve current stats from wordies http service.

	⠙ send        Send a string of sentences to the tcp server

	⠙ serve        Serve tcp and http word frequency service.

⡿ HELP:
	Run [command] help

⡿ OTHERS:
	Run 'wordies flags' to print all flags of all commands.

⡿ WARNING:
	Uses internal flag package so flags must precede command name. 
	e.g 'wordies -cmd.flag=4 run'

```


To run the wordies service, run `wordies serve`


```bash
>  wordies serve help
  Command: wordies [flags] serve 
  
  ⡿ DESC:
  	serve starts the tcp and http components of the natural language word frequency service.
  
  ⡿ Flags:
  	
  	⠙ serve.workers
  	 Default: 1000
  	 Desc: sets the maximum workers for background language processing requests
  	
  	⠙ serve.job.buffer
  	 Default: 500
  	 Desc: sets the maximum buffer to queue processing jobs
  	
  	⠙ serve.workers.timeout
  	 Default: 30s
  	 Desc: sets the maximum duration allowed for a worker to be idle
  	
  	⠙ serve.httpaddr
  	 Default: localhost:8080
  	 Desc: sets the address for the http level service
  	
  	⠙ serve.tcpaddr
  	 Default: localhost:5555
  	 Desc: sets the address for the tcp level service
  	
  ⡿ Examples:
  	
  ⡿ USAGE:
  	
  	⠙ wordies -serve.workers=1000 serve 
  	
  	⠙ wordies -serve.job.buffer=500 serve 
  	
  	⠙ wordies -serve.workers.timeout=30s serve 
  	
  	⠙ wordies -serve.httpaddr=localhost:8080 serve 
  	
  	⠙ wordies -serve.tcpaddr=localhost:5555 serve 
  	
  ⡿ OTHERS:
  	Commands which respect context.Context, can set timeout by using the -timeout flag.
  	e.g -timeout=4m, -timeout=4h
  
  ⡿ WARNING:
  	Uses internal flag package so flags must precede command name. 
  	e.g 'wordies -cmd.flag=4 run'
  
```

Before running any of the instructions below, ensure to have started the `wordies` service with `wordies serve`.

- Send sentences to the wordies tcp server

```bash
> wordies send "i want to take a moment to relax and take a strong"
```

```bash
> ϟ (wordies:master) wordies send help
  Command: wordies [flags] send 
  
  ⡿ DESC:
  	send provides a command that sends provided sentences to wordies tcp service if running
  
  ⡿ Flags:
  	
  	⠙ send.tcpaddr
  	 Default: localhost:5555
  	 Desc: sets the address for the tcp level service
  	
  	⠙ send.timeout
  	 Default: 2s
  	 Desc: sets the maximum duration allowed to wait to connect to service
  	
  ⡿ Examples:
  	
  ⡿ USAGE:
  	
  	⠙ wordies -send.tcpaddr=localhost:5555 send 
  	
  	⠙ wordies -send.timeout=2s send 
  	
  ⡿ OTHERS:
  	Commands which respect context.Context, can set timeout by using the -timeout flag.
  	e.g -timeout=4m, -timeout=4h
  
  ⡿ WARNING:
  	Uses internal flag package so flags must precede command name. 
  	e.g 'wordies -cmd.flag=4 run'
  
```

- Retrieve latest stats from the wordies service

```bash
> wordies stats
```

```bash
> ϟ (wordies:master) wordies stat help
  Usage: wordies [flags] [command] 
  
  ⡿ COMMANDS:
  	⠙ stats        Retrieve current stats from wordies http service.
  
  	⠙ send        Send a string of sentences to the tcp server
  
  	⠙ serve        Serve tcp and http word frequency service.
  
  ⡿ HELP:
  	Run [command] help
  
  ⡿ OTHERS:
  	Run 'wordies flags' to print all flags of all commands.
  
  ⡿ WARNING:
  	Uses internal flag package so flags must precede command name. 
  	e.g 'wordies -cmd.flag=4 run'
  
```
