package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall"
)

func maxOpenFiles() {
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Println("Error Getting Rlimit ", err)
	}

	if rLimit.Cur < rLimit.Max {
		rLimit.Cur = rLimit.Max
		err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			log.Println("Error Setting Rlimit ", err)
		}
	}
}

func main() {
	//token := os.Getenv("token")

	//bot, _ := tgbotapi.NewBotAPI(token)
	c := http.Client{}
	url := "https://api.telegram.org/bot1574524831:AAEcjMbq_hKCyVlJtY9Qd4I29Wq0WLBOUSw/sendMessage?chat_id=@test123456&text=123"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, _ := c.Do(req)
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
