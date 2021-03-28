package methodAndInterface

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getMicrosoft() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}
