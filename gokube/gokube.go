package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//执行shell脚本
func execShell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

func update(c echo.Context) error {
	// 从请求参数里获取 team 和 member 的值
	deployname := c.QueryParam("deployname")
	member := c.QueryParam("member")

	ret, _ := execShell("/tmp/test.sh")

	return c.String(http.StatusOK, "ret:"+ret+",team:"+team+", member:"+member)
}

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":1323"))
}
