package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		resp, err := GetHttp("https://jmssub.net/members/getsub.php?service=373767&id=c362f7a5-c613-455f-844b-78fb45cd77f2")
		if err != nil {
			fmt.Println(err)
		}
		return c.String(http.StatusOK, string(resp))
	})
	e.Logger.Fatal(e.Start(":1323"))

}

func GetHttp(url string) (body []byte, err error) {
	var client http.Client
	var resp *http.Response
	client = http.Client{Timeout: 10 * time.Second}

	resp, err = client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()

	return body, nil
}
