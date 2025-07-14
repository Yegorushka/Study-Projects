package main

import (
	"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"strings"
	"regexp"
)

func request(req string) {
	if req == "list"{
		resp, _ := http.Get("http://localhost:8080/" + req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

	} else {
		resp, _ := http.Get("http://localhost:8080/" + req)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile(`name=([^\s&]+)`)
		name := re.FindStringSubmatch(req)
		f, _ := os.Create(name[1])
		defer f.Close()
		_, _ = f.Write(body)
		fmt.Println("Файл успешно скачен!")
	}


}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Выберите функцию\n1. List\n2. Download")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

    if input == "1"{
    	request("list")
    } else {
    	fmt.Println("Введите имя файла")
    	input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
    	request("download?name=" + input)
    }


}
