package neynarsdk

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRetrieveCastsFromFilters(t *testing.T) {
	
	var httpClient *http.Client
	client := NewClient(httpClient,"TEST_API_KEY")
	res, err := client.Feed.RetrieveCastsFromFilters(context.Background(),RetrieveCastsFromFiltersParams{
		FeedType:    FeedType("filter"),
		FilterType:  FilterType("fids"),
		// Fid:         123,
		Fids:        "2",
		// ParentUrl:   "https://example.com/parent",
		// ChannelId:   "channel123",
		// EmbedUrl:    "https://example.com/embed",
		OmitRecasts: true,
		Limit:       1,
		// Cursor:      "abc123",
		// ViewerFid:   456,
	})
	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(res.Casts)
	}
	
}