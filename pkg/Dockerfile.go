package ScriptHelper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func readDockerfile() (map[string]string, error) {
	fi, err := os.Open("Dockfile")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, nil
	}
	defer fi.Close()

	r := make(map[string]string)

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		line := string(a)

		strArr := strings.Split(line, ` `)
		if len(strArr) == 0 {
			continue
		}
		r[strArr[0]] = strings.Join( strArr[1:] , ` `)
		fmt.Println(line)
	}

	return r, nil
}

func writeDockerfile(dockerContent map[string]string){
	fi, err := os.Open("Dockfile")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	bw := bufio.NewWriter(fi)

	for k, v := range dockerContent {
		bw.WriteString(k + " " + v)
	}
}