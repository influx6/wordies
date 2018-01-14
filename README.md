Wordies
---------
[![Go Report Card](https://goreportcard.com/badge/github.com/influx6/wordies)](https://goreportcard.com/report/github.com/influx6/wordies)
[![Travis Build Status](https://travis-ci.org/travis-ci/github.com/influx6/wordies)](https://goreportcard.com/report/github.com/influx6/wordies)

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


Before running any of the instructions below, ensure to have started the `wordies` service with `wordies serve`.

- Send sentences to the wordies tcp server

```bash
> wordies send "i want to take a moment to relax and take a strong"
```

- Retrieve latest stats from the wordies service

```bash
> wordies stats
```
