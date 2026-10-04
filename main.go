package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/macos")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	body, err := io.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)
		return

	}

	fmt.Println(string(body))

	file, err := os.Create("test.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.Close()
	if err != nil {
		return
	}

	err = os.WriteFile("test.md", []byte(fmt.Sprintf("---\n%s\n---\n", string(body))), 0644)
	if err != nil {
		return
	}

}
