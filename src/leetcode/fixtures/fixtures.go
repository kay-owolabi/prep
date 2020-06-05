package fixtures

import (
	"fmt"
	"io/ioutil"
	"leetcode/trees"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(filename string, numOfLines int, funcName ...string) []string {
	var methodName string
	if len(funcName) > 0 {
		for _, name := range funcName {
			name = strings.TrimSpace(name)
			if name != "" {
				methodName = name
				break
			}
		}
	}

	if methodName == "" {
		methodName = "readLines"
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fixtures %s: %v\n", methodName, err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			lines = append(lines[:i], lines[i+1:]...)
			i--
		}
	}

	if len(lines) < numOfLines {
		fmt.Fprintf(os.Stderr, "fixtures %s: %s should have at least %d non empty lines\n",
			methodName, filename, numOfLines)
		return nil
	}

	return lines
}

func readIntArray(str string) []int {
	var array []int
	regex := regexp.MustCompile(`(-?\d+,?)+`)
	commaSeparated := regex.FindString(str)
	items := strings.Split(commaSeparated, ",")
	for _, digits := range items {
		num, _ := strconv.Atoi(digits)
		array = append(array, num)
	}
	return array
}

func read2DIntArray(line string) [][]int {
	var array [][]int
	regex := regexp.MustCompile(`\[(-?\d+,?)+]`)
	intArrays := regex.FindAllString(line, -1)
	for _, intArrayStr := range intArrays {
		array = append(array, readIntArray(intArrayStr))
	}
	return array
}

func createTree(input []int) *trees.TreeNode {
	var root *trees.TreeNode
	queue := []**trees.TreeNode{&root}

	for _, item := range input {
		node := new(trees.TreeNode)
		node.Val = item
		*(queue[0]) = node
		queue = append(queue, &node.Left, &node.Right)
		queue = queue[1:]
	}
	return root
}

func readWord(line string) string {
	regex := regexp.MustCompile(`\w+`)
	str := regex.FindString(line)
	return str
}

func ReadCriticalConnectionsTest(filename string) (int, [][]int) {
	var count int
	var array [][]int

	lines := readLines(filename, 2, "ReadCriticalConnectionsTest")

	count, err := strconv.Atoi(lines[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "fixtures ReadCriticalConnectionsTest: first Line should be an int %v\n", err)
	}
	array = read2DIntArray(lines[1])
	return count, array
}

func ReadInvertBinaryTreeTest(filename string) *trees.TreeNode {
	lines := readLines(filename, 1, "ReadInvertBinaryTreeTest")
	arrayNodes := readIntArray(lines[0])
	return createTree(arrayNodes)
}

func ReadClosetPointOrigin(filename string) (points [][]int, K int) {
	lines := readLines(filename, 2, "ReadClosetPointOrigin")
	K, _ = strconv.Atoi(lines[1])
	return read2DIntArray(lines[0]), K
}

func ReadEditDistance(filename string) (word1 string, word2 string) {
	lines := readLines(filename, 2, "ReadEditDistance")
	word1 = readWord(lines[0])
	word2 = readWord(lines[1])
	return word1, word2
}

func ReadTwoCityScheduling(filename string) (costs [][]int) {
	lines := readLines(filename, 1, "ReadTwoCityScheduling")
	return read2DIntArray(lines[0])
}
