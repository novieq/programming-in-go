package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"

	"os"
)

//we define a new flag "n", which defaults to off. Note that we get the help (-h) for free
//this defines a boolean flag, -n, stored in the pointer numberFlag, with type *bool
var numberFlag = flag.Bool("n", false, "number each line")

//this function actually reads the files contents and then displays it
func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, e := r.ReadBytes('\n') //read one line at a time
		if e == io.EOF {            //hit stop if we reach end
			break
		}
		if *numberFlag { //if we have to number each line
			fmt.Fprintf(os.Stdout, "%5d  %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func main() {
	//After all the flags are defined this is called to parse the command line to the flags
	flag.Parse()
	//If there are no arguments passed in cmd then you just print what is being entered
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, e := os.Open(flag.Arg(i))
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), e.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}