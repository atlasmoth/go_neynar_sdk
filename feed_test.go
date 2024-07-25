package neynarsdk

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestRetrieveCastsFromFilters(t *testing.T) {
	var httpClient *http.Client
	client := NewClient(httpClient, "NEYNAR_API_KEY")
	res, err := client.Feed.RetrieveTrendingCasts(context.Background(), RetrieveTrendingCastsParams{

		// FeedType:    FeedType("following"),
		// FilterType:  FilterType("fids"),
		// Fid:         123,
		// Fids:        "2",
		// ParentUrl:   "https://example.com/parent",
		// ChannelId:   "channel123",
		// EmbedUrl:    "https://example.com/embed",
		// OmitRecasts: true,
		// Limit:       1,
		// ViewerFid: 1,
		// Cursor:      "eyJ0aW1lc3RhbXAiOiIyMDI0LTA3LTI1IDEwOjExOjAzLjAwMDAwMDAifQ%3D%3D",
		// ViewerFid:   456,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Casts length: %v \n\n\n", len(res.Casts))
		fmt.Printf("%+v\n\n", res)
	}

}
