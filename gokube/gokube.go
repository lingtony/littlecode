package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

//func checkErr(err error) {
//	if err != nil {
//		fmt.Println(err)
//		panic(err)
//	}
//}
//
////执行shell脚本
//func execShell(s string) (string, error) {
//	cmd := exec.Command("/bin/bash", s)
//
//	var out bytes.Buffer
//	cmd.Stdout = &out
//
//	err := cmd.Run()
//	checkErr(err)
//
//	return out.String(), err
//}

func update(c echo.Context) error {
	deployname := c.QueryParam("deployname")
	image := c.QueryParam("image")
	namespase := c.QueryParam("namespace")
	updateCmd := "set image deploy" + " " + deployname + " " + deployname + "=" + image + " " + "-n" + " " + namespase
	testcmd := "get po -n " + namespase
	cmd := exec.Command("/usr/bin/kubelet", testcmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.String(http.StatusOK, out.String()+"\n"+updateCmd)
}

func main() {
	e := echo.New()
	e.GET("/update", update)
	e.Logger.Fatal(e.Start(":1323"))
}
