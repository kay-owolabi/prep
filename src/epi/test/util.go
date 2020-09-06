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

type Compare interface {
	Len() int
	Compare(i, j int, other interface{}) int
}

type StringSlice []string

func (s StringSlice) Len() int { return len(s) }

func (s StringSlice) Compare(i, j int, other interface{}) int {
	return strings.Compare(s[i], other.(StringSlice)[j])
}

type IntSlice []int

func (I IntSlice) Len() int { return len(I) }

func (I IntSlice) Compare(i, j int, other interface{}) int {
	return (I[i] - other.(IntSlice)[j]) / 1
}

func LexicographyComp(l1, l2 Compare) bool {
	var result int
	for ok, i, j := true, 0, 0; ok; i, j, ok = i+1, j+1, result == 0 {
		if i >= l1.Len() {
			if j >= l2.Len() {
				return false
			} else {
				return true
			}
		}
		if j >= l2.Len() {
			return false
		}
		result = l1.Compare(i, j, l2)
	}
	return result < 0
}

const FailureInfo = "Failure info\n" + "\texpected: %v\n" + "\tgot: %v\n"
