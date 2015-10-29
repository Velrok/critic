package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func persist_message(text string) {
	text_lenth := len(text)
	max_length := int(math.Min(80.0, float64(text_lenth)))

	filename := fmt.Sprintf("/tmp/critic_%v.edn", time.Now().Unix())

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
	text_size := float64(strings.NewReader(text).Len()) / 1000

	fmt.Printf("message size: %15.0f kb\n", text_size)
	if text_size > 900 {
		persist_message(text)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(0)
		}
		check(err)
		go process_message(line)
	}

	//process_message("Hello")
	//process_message("Hello my Dear!")
}
