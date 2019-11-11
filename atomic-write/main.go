package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	xFlag = flag.Bool("x", false, "Exclude O_APPEND flag")
	bFlag = flag.Bool("b", false, "Use buffer writer")
)

func write(f *os.File) (int, error) {
	if *xFlag {
		f.Seek(0, os.SEEK_END)
	}
	f.Write([]byte{'a'})
	return 0, nil
}

func main() {
	flag.Parse()
	fmt.Println("x :=", *xFlag)

	fFlags := os.O_CREATE | os.O_WRONLY
	if !*xFlag {
		fmt.Println("Enable append flag")
		fFlags |= os.O_APPEND
	}

	args := flag.Args()
	if len(args) != 2 {
		panic("invalid args")
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(args[0], fFlags, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	for i := 0; i < num; i++ {
		if *bFlag {
			writer.WriteByte('a')
		} else {
			write(f)
		}
	}
}
