package hashtables

import (
	"sort"
	"strings"
)

// 12.2 Is An Anonymous Letter Constructible ?

func IsLetterConstructibleFromMagazine(letterText, magazineText string) bool {
	letterDic := map[string]int{}
	magazineDic := map[string]int{}

	for _, char := range strings.Split(letterText, "") {
		if _, ok := letterDic[char]; !ok {
			letterDic[char] = 0
		}
		letterDic[char] += 1
	}

	for _, char := range strings.Split(magazineText, "") {
		if _, ok := magazineDic[char]; !ok {
			magazineDic[char] = 0
		}
		magazineDic[char] += 1
	}

	for text, count := range letterDic {
		mCount, ok := magazineDic[text]
		if !ok || mCount < count {
			return false
		}
	}
	return true
}

func FindAnagrams(dictionary []string) [][]string {
	sortedStringToAnagram := map[string][]string{}
	for _, s := range dictionary {
		chars := sort.StringSlice(strings.Split(s, ""))
		chars.Sort()
		key := bytes.s //strings.Join(chars, "")

		if _, ok := sortedStringToAnagram[key]; !ok {
			sortedStringToAnagram[key] = []string{}
		}
		sortedStringToAnagram[key] = append(sortedStringToAnagram[key], s)
	}
	var result [][]string
	for _, val := range sortedStringToAnagram {
		if len(val) > 1 {
			result = append(result, val)
		}
	}
	return result
}
