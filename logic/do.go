package logic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Struct2Pb(path string) string {
	var result string
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	var index int
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lineResult := dealLine(string(line))
		if lineResult == "" {
			continue
		}
		index++
		lineResult = fmt.Sprintf("%s%d; \n", lineResult, index)
		result += lineResult
	}
	return result
}

func dealLine(line string) string {
	if !strings.Contains(line, "json:\"") {
		return ""
	}
	switch LineType(line).GetType() {
	case mapType:
		return ""
	case strType:
		return caculateName(line) + fmt.Sprintf(" %s = ", "string")
	case int64Type:
		return caculateName(line) + fmt.Sprintf(" %s = ", "int64")
	case intType:
		return caculateName(line) + fmt.Sprintf(" %s = ", "int")
	case boolType:
		return caculateName(line) + fmt.Sprintf(" %s = ", "bool")
	}
	return ""
}

type LineType string

const (
	mapType = iota
	strType
	int64Type
	intType
	boolType
)

func (l LineType) GetType() int {
	if strings.Contains(string(l), "map") {
		return mapType
	}
	if strings.Contains(string(l), "string") {
		return strType
	}
	if strings.Contains(string(l), "int64") {
		return int64Type
	}
	if strings.Contains(string(l), "int") {
		return intType
	}
	if strings.Contains(string(l), "int") {
		return boolType
	}
	return -1
}

func caculateName(line string) string {
	//json:""
	index := strings.Index(line, "json:")
	lineCopy := line[index+6:]
	endIndex1 := strings.Index(lineCopy, "\"")
	endIndex2 := strings.Index(lineCopy, ",")
	var endIndex int
	if endIndex2 < 0 {
		endIndex = endIndex1
	}else {
		if endIndex2 >= endIndex1 {
			endIndex = endIndex1
		}else {
			endIndex = endIndex2
		}
	}

	name := lineCopy[0:endIndex]
	return name
}
