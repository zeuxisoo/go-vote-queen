package main

import (
    "os"
    "fmt"
    "flag"
    "errors"
)

type Config struct {
    Key     string
    Area    string
}

func newConfig() (c *Config) {
    return &Config{}
}

func (c *Config) check() (string, error) {
    if c.Key == "" {
        return "", errors.New("Please enter api key")
    }

    if c.Area == "" {
        return "", errors.New("Please enter api area")
    }

    return "", nil
}

func usage() {
    const usage = `Go-Queen: a simple auto vote program

Usage:

    go-queen [-k KEY] [-a AREA]
    go-queen -h | --help

Options:

    -k,             The api key for request the proxy api
    -a,             The area of proxy list
    -h, --help      Output help information
`

    fmt.Printf(usage)
    os.Exit(0)
}


func main() {
    var key, area string
    var help bool

    config := newConfig()

    flag.StringVar(&key,  "k",    "",    "Porxy api key")
    flag.StringVar(&area, "a",    "JP",  "Proxy area")
    flag.BoolVar(&help,   "h",    false, "Show help message")
    flag.BoolVar(&help,   "help", false, "Show help message")

    flag.Usage = usage

    flag.Parse()

    if help {
        usage()
    }

    config.Key  = key
    config.Area = area

    if _, err := config.check(); err != nil {
        fmt.Printf("Arguments error: %s", err)
    }else{
        fmt.Println("Hello world")
    }
}
