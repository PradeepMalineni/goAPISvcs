package getdata

import (
	"fmt"
	"net/http"
)

func CallHttpData() {
	httpbinurl := "http://httpbin.org/get"

	res, err := http.Get(httpbinurl)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)

}
