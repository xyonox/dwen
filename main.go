package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func run() error {

	website := flag.String("w", "", "URL to website")

	flag.Parse()

	if *website == "" {
		return fmt.Errorf("-w flag is required. This flag set the URL to download")
	}

	resp, err := http.Get(*website)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err

	}

	file, err := os.Create("test.md")
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	err = os.WriteFile("test.md", []byte(fmt.Sprintf("---\n%s\n---\n", string(body))), 0644)
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
