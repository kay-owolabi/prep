package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	FIELD_DELIM      = "\t"
	MAX_SEARCH_DEPTH = 4
	DIR_NAME         = "test_data"
)

func GetTestData(testDataFile string) [][]string {
	testDataDir := getDefaultTestDataDirPath()
	testData := splitTsvFile(filepath.Join(testDataDir, testDataFile))
	return testData
}

func retErr(err error) interface{} {
	fmt.Fprintf(os.Stderr, "test util err: %v", err)
	return nil
}

func getDefaultTestDataDirPath() string {
	curDir, err := filepath.Abs(".")
	if err != nil {
		return retErr(err).(string)
	}

	for i := 0; i < MAX_SEARCH_DEPTH; i++ {
		path := filepath.Join(curDir, DIR_NAME)
		fileInfo, err := os.Stat(path)
		if !os.IsNotExist(err) && fileInfo.IsDir() {
			return path
		}
		curDir = filepath.Dir(curDir)
		if curDir == "" {
			break
		}
	}
	panic("unable to find test_data dir")
}

func splitTsvFile(tsvfile string) [][]string {
	inputData, err := ioutil.ReadFile(tsvfile)
	if err != nil {
		return retErr(err).([][]string)
	}

	lines := strings.Split(string(inputData), "\n")
	var result [][]string

	for _, line := range lines {
		if line == "" {
			continue
		}
		result = append(result, strings.Split(line, FIELD_DELIM))
	}
	return result
}
