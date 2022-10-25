package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func constructCRD(crd string) string {

	f, _ := os.Create("test.yaml")

	f.WriteString(fmt.Sprintf("apiVersion: v1\nkind: ConfigMap\nmetadata:\n name: test-game\ndata:\n dashbaord.json:  | \n   %s", crd))

	cmd := exec.Command("kubectl", "apply", "-f", "/Users/mangaalangommeetei/Documents/learning/test.yaml")

	out, err := cmd.Output()

	if err != nil {

		fmt.Println(string(out))

		return err.Error()
	}

	fmt.Println(string(out))

	return "done"

}

func uploadJSON(c *gin.Context) {

	file, err := c.GetRawData()

	if err != nil {
		c.JSON(500, err.Error())

		return
	}

	var out bytes.Buffer
	json.Indent(&out, file, "    ", "  ")

	print := constructCRD(out.String())

	c.JSON(200, print)

}
