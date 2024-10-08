/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ChannelActivity struct {
	Object string `json:"object"`
	CastCount1d string `json:"cast_count_1d"`
	CastCount7d string `json:"cast_count_7d"`
	CastCount30d string `json:"cast_count_30d"`
	Channel *Channel `json:"channel"`
}
