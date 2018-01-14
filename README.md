Wordies
---------

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

```

