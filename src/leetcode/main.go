package main

import (
	"fmt"
	"leetcode/fixtures"
	"leetcode/solutions"
)

func main() {
	testFileName := "/Users/koowolab/gitroot/kay-owolabi/prep/src/leetcode/fixtures/twocityscheduling/test1.in"
	result := solutions.TwoCitySchedCost(fixtures.ReadTwoCityScheduling(testFileName))
	fmt.Printf("%v\n", result)
}
