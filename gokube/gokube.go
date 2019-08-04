package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo"
)



func execCmd(s string) (string,error){
}


func update(c echo.Context) error {
	deployname := c.QueryParam("deployname")
	image := c.QueryParam("image")
	namespace := c.QueryParam("namespace")
	updateCmd := "kubectl set image deploy" + " " + deployname + " " + deployname + "=" + image + " " + "-n" + " " + namespace
        testcmd := "kubectl get deploy "+deployname+" -n "+namespace
	upcmd := exec.Command("/bin/bash","-c",updateCmd)
	testcmd := exec.Command("/bin/bash","-c",updateCmd)
	var upout bytes.Buffer
	var testout bytes.Buffer
	upcmd.Stdout = &upout
	testcmd.Stdout = &testout
	uperr := upcmd.Run()
	if uperr != nil {
		fmt.Println(uperr)
		panic(uperr)
	}
	testerr := upcmd.Run()
	if testerr != nil {
		fmt.Println(err)
		panic(err)
	}

	return c.String(http.StatusOK, out.String()+"\n"+updateCmd)
}

func main() {
	e := echo.New()
	e.GET("/update", update)
	e.Logger.Fatal(e.Start(":8080"))
}
