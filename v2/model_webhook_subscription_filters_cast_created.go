/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type WebhookSubscriptionFiltersCastCreated struct {
	// Exclude casts that matches these authors. **Note:** This is applied as an AND operation against rest of the filters. Rest of the filters are bundled as an OR operation. 
	ExcludeAuthorFids []int32 `json:"exclude_author_fids,omitempty"`
	AuthorFids []int32 `json:"author_fids,omitempty"`
	MentionedFids []int32 `json:"mentioned_fids,omitempty"`
	ParentUrls []string `json:"parent_urls,omitempty"`
	RootParentUrls []string `json:"root_parent_urls,omitempty"`
	ParentHashes []string `json:"parent_hashes,omitempty"`
	ParentAuthorFids []int32 `json:"parent_author_fids,omitempty"`
	// Regex pattern to match the text key of the cast. **Note:**  1) Regex must be parsed by Go's RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: (?i)\\\\$degen should be written as (?i)\\\\\\\\$degen 
	Text string `json:"text,omitempty"`
	// Regex pattern to match the embeded_url (key embeds) of the cast. **Note:**  1) Regex must be parsed by Go's RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: \\\\b(farcaster|neynar)\\\\b should be written as \\\\\\\\b(farcaster|neynar)\\\\\\\\b 
	Embeds string `json:"embeds,omitempty"`
}
