package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrefixes(t *testing.T) {
	test := func(a string, b []string) {
		require.Equal(t, a, prefixes(b), fmt.Sprintf("%#v", b))
	}

	test("", []string{"", ""})
	test("", []string{""})
	test("", []string{})
	test("", nil)
	test("/", []string{"/", "/"})
	test("/", []string{"/"})
	test("/a", []string{"/a"})
	test("a", []string{"a"})
	test("a/b/c", []string{"a/b/c"})

	test("/a/b/c", []string{
		"/a/b/c",
		"/a/b/c/d",
	})
	test("/a/b/c", []string{
		"/a/b/c",
		"/a/b/c/d/e",
	})
	test("/a/b/c", []string{
		"/a/b/c",
		"/a/b/c/",
	})
	test("/a/b/c", []string{
		"/a/b/c/",
		"/a/b/c/d/e",
	})
	test("/", []string{
		"/a/b/c",
		"/d/e/f",
	})
	test("", []string{
		"/a/b/c",
		"d/e/f",
	})
	test("a/b/c", []string{
		"a/b/c/x/y/z",
		"a/b/c/d/e",
	})
	test("a/b/c", []string{
		"a/b/c/x/y/z",
		"a/b/c/x/y",
		"a/b/c/x/",
		"a/b/c/d/e",
	})
}
