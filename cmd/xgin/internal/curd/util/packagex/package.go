package packagex

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetRootPackage() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return ""
	}
	modFile := path + "/go.mod"
	f, err := os.Open(modFile)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	line, err := buf.ReadString('\n')
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return ""
	}
	strs := strings.Split(line, " ")
	if len(strs) != 2 {
		return ""
	}
	return strs[1]
}
