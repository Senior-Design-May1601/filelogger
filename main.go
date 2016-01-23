package main

import (
    "bufio"
    "log"
    "os"

    "github.com/Senior-Design-May1601/projectmain/loggerplugin"
)

type TextfilePlugin struct {
    w *bufio.Writer
}

func (x *TextfilePlugin) Log(msg []byte, _ *int) error {
    log.Println(string(msg))
	return nil
}

func main() {
	out, err := os.Create("log-example.txt")
	if err != nil {
        panic(err)
    }
	defer func() {
		if err := out.Close(); err != nil {
			panic(err)
		}
	}()

    w := bufio.NewWriter(out)
    log.SetOutput(out)
    p, err := loggerplugin.NewLoggerPlugin(&TextfilePlugin{w})
    if err != nil {
        log.Fatal(err)
    }

    err = p.Run()
    if err != nil {
        log.Fatal(err)
    }
}
