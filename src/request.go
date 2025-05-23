package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	url := "http://localhost:8080/user"


	resp, err := client.R().
		SetBody(map[string]interface{}{
			"Name": "Robô",
			"DisplayName": "Robô",
		}).
		Post(url)

	if err != nil {
		fmt.Println("[erro]: ", err)
		fmt.Scanln()
	}

	if resp.StatusCode() == 200 {
		fmt.Println("[200] ok")
		fmt.Scanln()
	}
}
