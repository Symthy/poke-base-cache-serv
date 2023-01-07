package client

import (
	"errors"
	"regexp"
	"strconv"
)

var noOffsetErr = errors.New("non exist offset")

func extractOffset(url string) (int, error) {
	re, _ := regexp.Compile("offset=(\\d+)")
	fs := re.FindString(url)
	if fs == "" {
		return 0, noOffsetErr
	}
	return strconv.Atoi(fs[7:])
}
