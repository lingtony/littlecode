package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)

func update(c echo.Context) error {
	deployname := c.QueryParam("deployname")
	image := c.QueryParam("image")
	namespace := c.QueryParam("namespace")
	updatecmd := "kubectl set image deploy" + " " + deployname + " " + deployname + "=" + image + " " + "-n" + " " + namespace
	cmd := exec.Command("/bin/bash", "-c", updatecmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.String(http.StatusOK, out.String())
}

func checkDeploy(c echo.Context) error {
	deployname := c.QueryParam("deployname")
	namespace := c.QueryParam("namespace")
	testcmd := "kubectl get deploy " + deployname + " -n " + namespace
	cmd := exec.Command("/bin/bash", "-c", testcmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return c.String(http.StatusOK, out.String())
}
func main() {
	e := echo.New()
	e.GET("/update", update)
	e.GET("/checkdeploy", checkDeploy)
	e.Logger.Fatal(e.Start(":8080"))
}
