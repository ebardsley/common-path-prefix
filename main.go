package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const endMarker = "\000"

type pathMap map[string]*pathMap

func main() {
	p, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(prefixes(strings.Split(strings.TrimSpace(string(p)), "\n")))
}

func prefixes(l []string) string {
	// For each entry, break it down based on the path separator and add it to a
	// trie of directory names.
	root := &pathMap{}
	for _, e := range l {
		m := root
		for _, d := range strings.Split(e, string(filepath.Separator)) {
			if (*m)[d] == nil {
				(*m)[d] = &pathMap{}
			}
			m = (*m)[d]
		}
		(*m)[endMarker] = nil
	}

	// Walk the trie and construct the common path, stopping when we reach the end
	// of a path or there is more than one choice.
	ret := []string{}
	for {
		if len(*root) != 1 {
			break
		}
		if _, ok := (*root)[endMarker]; ok {
			break
		}

		// It is guaranteed that there is only one element.
		for k, v := range *root {
			ret = append(ret, k)
			root = v
		}
	}

	// Special case: if all paths are an absolute path but have no other common
	// elements, the result should be the root directory.
	if len(ret) == 1 && ret[0] == "" && strings.HasPrefix(l[0], string(filepath.Separator)) {
		return string(filepath.Separator)
	}

	return strings.Join(ret, string(filepath.Separator))
}
