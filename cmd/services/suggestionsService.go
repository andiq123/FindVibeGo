package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SuggestionsService struct {
	suggestionLink string
}

func NewSuggestionsService() *SuggestionsService {
	suggestionLink := "https://clients1.google.com/complete/search?client=youtube&gs_ri=youtube&ds=yt&q="
	return &SuggestionsService{suggestionLink: suggestionLink}
}

func (s *SuggestionsService) GetSuggestions(searchQuery string) ([]string, error) {
	fullLink := fmt.Sprintf("%s%s", s.suggestionLink, searchQuery)
	resp, err := http.Get(fullLink)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	stringBody := string(bodyBytes)

	startIndex := strings.Index(stringBody, "[")
	endIndex := strings.LastIndex(stringBody, "]")

	if startIndex == -1 || endIndex == -1 {
		return nil, errors.New("failed to find suggestions")
	}

	jsonString := stringBody[startIndex : endIndex+1]

	var data []any
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, len(data[1].([]any)))
	for _, item := range data[1].([]any) {
		if str, ok := item.([]any)[0].(string); ok {
			results = append(results, str)
		}
	}

	return results, nil
}
