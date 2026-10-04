package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func run() error {

	website := flag.String("w", "", "URL to website")

	flag.Parse()

	if *website == "" {
		return fmt.Errorf("-w flag is required. This flag set the URL to download")
	}

	split := strings.Split(*website, "/")
	name := split[len(split)-1]

	if !(strings.Contains(*website, "https://")) && !(strings.Contains(*website, "http://")) {
		*website = fmt.Sprintf("https://%s", *website)
	}

	resp, err := http.Get(*website)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	length := resp.ContentLength / 1000000 // Formated into MB
	fmt.Println(length)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
