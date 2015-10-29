package main

import (
	"bufio"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

var (
	suffix        = kingpin.Flag("suffix", "The suffix to be used when writing the files.").Default("json").Short('f').String()
	critical_size = kingpin.Flag("size", "If the message received is bigger then <size> bytes write the message.").Default("900").Short('s').Int()
	out_dir       = kingpin.Flag("output-dir", "Messages are written into this directory.").Default("/tmp").Short('o').String()
	message_sep   = kingpin.Flag("message-separator", "A character that is the separator between messages.").Default("\n").String()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func persist_message(text string) {
	text_lenth := len(text)
	max_length := int(math.Min(80.0, float64(text_lenth)))

	filename := fmt.Sprintf("%v/critic_%v.%v", *out_dir, time.Now().Unix(), *suffix)

	preview := text[0:max_length]
	dots := ""

	if max_length < text_lenth {
		dots = "..."
	}
	fmt.Printf("\nPersisting message: %v%v as %v\n", preview, dots, filename)

	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	_, err = f.WriteString(text)
	check(err)
}

func process_message(text string) {
	text_size := strings.NewReader(text).Len()

	fmt.Printf("message size: %15.0f kb\n", text_size)
	if text_size > *critical_size {
		persist_message(text)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	kingpin.Parse()

	var sep byte
	sep = (*message_sep)[0]

	for {
		line, err := reader.ReadString(sep)
		if err == io.EOF {
			os.Exit(0)
		}
		check(err)
		go process_message(line)
	}

	//process_message("Hello")
	//process_message("Hello my Dear!")
}
