package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for i := range strs {
		temp := strs[i]
		key := sortString(temp)
		groups[key] = append(groups[key], strs[i])
	}
	ret := make([][]string, 0)
	for _, v := range groups {
		ret = append(ret, v)
	}
	return ret
}

func main() {
	strs := []string{
		// "eat", "tea", "tan", "ate", "nat", "bat",
		"a",
	}
	fmt.Println(groupAnagrams(strs))
}
