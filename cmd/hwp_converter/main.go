package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func parseHWP(filePath string) (result bytes.Buffer, err error) {

	cmd := exec.Command("unhwp", "text", filePath) // it takes multiple arguments, doing what CLI should do.

	var out bytes.Buffer // bytes.Buffer has both Read and Write interfaces.
	cmd.Stdout = &out    // pour the output into out. I get confused by the way it is written
	// exec.Command produces a pointer *Cmd
	// *Cmd has the following interfaces with it. Stdout is io.Writer, so cmd.Stdout = &out may work..?
	// Shouldn't it be &out = cmd.Stdout? Like we put cmd.Stdout into &out?

	if err := cmd.Run(); err != nil {
		return bytes.Buffer{}, fmt.Errorf("Error:%w", err)
	} // cmd *Cmd has Run() method. Run starts the specified command and waits for it to complete.
	return out, nil
}

func main() {
	// 1. Get the text information.
	text, err := parseHWP("./books/자말묵.hwp")
	if err != nil {
		panic(err)
	}
	// Reading method is basically r(io.Reader).Read(p []byte) returning n, err
	bucket := make([]byte, 1024)
	var dataRead []byte

	for {
		n, err := text.Read(bucket)

		if n > 0 {
			dataRead = append(dataRead, bucket[:n]...)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	log.Printf("%s\n", dataRead)

	err = os.WriteFile("./books/자말묵.txt", dataRead, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	log.Printf("Success!")

}
