package utils

import (
	"distivity/types"
	"strings"
)

func FormatAuthors(authors []types.Author) string {
	var formattedAuthors []string
	for _, author := range authors {
		formattedAuthors = append(formattedAuthors, author.Codename)
	}
	return strings.Join(formattedAuthors, ", ")
}
