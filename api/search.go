package api

import (
	"context"
	"fmt"
	googleSearch "github.com/rocketlaunchr/google-search"
)

func SearchByText(text string) ([]googleSearch.Result, error) {
	ctx := context.Background()
	fmt.Println(googleSearch.Search(ctx, text))
	return googleSearch.Search(ctx, text)
}
