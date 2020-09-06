package main

import "fmt"

func main() {
	//testFileName := "/Users/koowolab/gitroot/kay-owolabi/prep/src/leetcode/fixtures/twocityscheduling/test1.in"
	//result := solutions.TwoCitySchedCost(fixtures.ReadTwoCityScheduling(testFileName))
	//fmt.Printf("%v\n", result)

	input := []string{"un", "iq", "ue", "the", "a", "there", "answer", "any", "by", "bye", "their"}
	root := Trie{
		char:     0,
		isWord:   false,
		children: map[string]*Trie{},
	}
	for _, str := range input {
		if len(str) == 0 {
			continue
		}
		if root.children[Char(str[0]).String()] == nil {
			root.children[Char(str[0]).String()] = &Trie{Char(str[0]), false, map[string]*Trie{}}
		}
		root.children[string(str[0])].Insert(str)
	}

	for _, str := range []string{"the", "these", "their", "thaw", "there"} {
		if len(str) == 0 {
			continue
		}
		key := Char(str[0]).String()
		if root.children[key] == nil {
			continue
		}
		fmt.Printf("%v\n", root.children[key].Find(str))
	}
}

/*
 def maxLength(self, A):
        dp = [set()]
        for a in A:
            if len(set(a)) < len(a): continue
            a = set(a)
            for c in dp[:]:
                if a & c: continue
                dp.append(a | c)
        return max(len(a) for a in dp)
*/

func maxLength(arr []string) int {
	dp := map[string]bool{"": true}
	for _, a := range arr {
		if len(set(a)) < len(a) {
			continue
		}

		for c, _ := range dp {
			if intersection(a, c) {
				continue
			}
			dp[a+c] = true
		}
	}
	var res int
	for str, _ := range dp {
		if len(str) > res {
			res = len(str)
		}
	}
	return res
}

func intersection(a string, c string) bool {
	zero := 'a'
	var aInt, cInt int
	for _, char := range a {
		aInt |= 1 << (char - zero)
	}
	for _, char := range c {
		cInt |= 1 << (char - zero)
	}

	return (aInt & cInt) != 0
}

func set(str string) string {
	mapChar := map[rune]bool{}
	var setRune []rune
	for _, c := range str {
		if mapChar[c] {
			continue
		}
		mapChar[c] = true
		setRune = append(setRune, c)
	}
	return string(setRune)
}

type Char byte

func (c Char) String() string {
	return string([]byte{byte(c)})
}

type Trie struct {
	char     Char
	isWord   bool
	children map[string]*Trie
}

func (t *Trie) Insert(key string) {
	if len(key) == 0 {
		return
	}
	if len(key) == 1 {
		t.isWord = true
		return
	}

	if t.children[Char(key[1]).String()] == nil {
		t.children[Char(key[1]).String()] = &Trie{Char(key[1]), false, map[string]*Trie{}}
	}
	t.children[Char(key[1]).String()].Insert(key[1:])
}

func (t *Trie) Find(key string) *Trie {
	if len(key) == 0 {
		return nil
	}

	if len(key) == 1 {
		if t.char == Char(key[0]) && t.isWord {
			return t
		}
		return nil
	}

	if t.children[Char(key[1]).String()] != nil {
		return t.children[Char(key[1]).String()].Find(key[1:])
	}
	return nil
}
