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

func Read2DIntArray(filename string) (int, [][]int) {
	var count int
	var array [][]int

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fixtures Read2DIntArray: %v\n", err)
		return 0, nil
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		fmt.Fprintf(os.Stderr, "fixtures Read2DIntArray: %s should have at least 2 lines\n", filename)
		return 0, nil
	}

	firstLine := len(lines)
	for index, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if firstLine > index {
			count, err = strconv.Atoi(line)
			firstLine = index
			if err != nil {
				fmt.Fprintf(os.Stderr, "fixtures Read2DIntArray: first Line should be an int %v\n", err)
			}
			continue
		}

		regex := regexp.MustCompile(`\d+,\d+`)
		strPairs := regex.FindAllString(line, -1)
		for _, strPair := range strPairs {
			itemsStr := strings.Split(strPair, ",")
			x, err := strconv.Atoi(itemsStr[0])
			if err != nil {
				continue
			}

			y, err := strconv.Atoi(itemsStr[1])
			if err != nil {
				continue
			}

			intPair := []int{x, y}
			array = append(array, intPair)
		}

	}

	return count, array
}

func ReadTree(filename string) *trees.TreeNode {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fixtures ReadTree: %v\n", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 1 {
		fmt.Fprintf(os.Stderr, "fixtures Read2DIntArray: %s should have at least 1 lines\n", filename)
		return nil
	}

	var root *trees.TreeNode
	var queue []**trees.TreeNode
	queue = append(queue, &root)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		regex := regexp.MustCompile(`\d+`)
		allInts := regex.FindAllString(line, -1)
		for _, strInt := range allInts {
			x, err := strconv.Atoi(strInt)
			if err != nil {
				queue = queue[1:]
				continue
			}
			node := new(trees.TreeNode)
			node.Val = x
			*queue[0] = node
			queue = append(queue, &node.Left, &node.Right)
			queue = queue[1:]
		}
	}

	return root
}
