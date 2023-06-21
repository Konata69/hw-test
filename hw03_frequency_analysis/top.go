package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type sortItem struct {
	key   string
	value int
}

func Top10(text string) []string {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[А-Яа-я]+(?:-?[А-Яа-я]+)*`)

	splited := re.FindAllString(text, -1)

	words := map[string]int{}

	for _, word := range splited {
		words[word]++
	}

	slice := make([]sortItem, 0)

	for key, word := range words {
		item := sortItem{key: key, value: word}
		slice = append(slice, item)
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i].value != slice[j].value {
			return slice[i].value > slice[j].value
		}
		return slice[i].key < slice[j].key
	})

	sliceLength := len(slice)
	if sliceLength > 10 {
		sliceLength = 10
	}
	result := make([]string, 0)
	for i := 0; i < sliceLength; i++ {
		result = append(result, slice[i].key)
	}

	return result
}
