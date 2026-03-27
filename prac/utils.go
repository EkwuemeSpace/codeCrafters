package main

import (
	"strconv"
)

func BaseConverter(word string, base int) (int64, error) {

	num, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
func dex(n int64, base int) (string, error) {
	return strconv.FormatInt(n, base), nil
}
