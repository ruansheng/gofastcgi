package main

import (
	"fastcgi-client/fastcgi"
	"fmt"
	"os"
)

func main()  {
	host := "127.0.0.1"
	port := 9000
	client, err := fastcgi.New(host, port)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	reqParams := "name=zhangsan"

	env := make(map[string]string)
	env["REQUEST_METHOD"] = "GET"
	env["SCRIPT_FILENAME"] = "/usr/local/php7/test/index.php"
	env["QUERY_STRING"] = reqParams

	content, err := client.Request(env, reqParams)
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	fmt.Printf("content: %s", content)
}
