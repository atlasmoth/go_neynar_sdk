package neynarsdk

import "time"



type Cast struct {
    Object       string         `json:"object"`
    Hash         string         `json:"hash"`
    ThreadHash   string         `json:"thread_hash"`
    ParentHash   interface{}    `json:"parent_hash"`
    ParentURL    string         `json:"parent_url"`
    RootParentURL string        `json:"root_parent_url"`
    ParentAuthor ParentAuthor   `json:"parent_author"`
    Author       Author         `json:"author"`
    Text         string         `json:"text"`
    Timestamp    string         `json:"timestamp"`
    Embeds       []Embed        `json:"embeds"`
    Reactions    Reactions      `json:"reactions"`
    Replies      Replies        `json:"replies"`
    Channel      Channel        `json:"channel"`
    MentionedProfiles []Author  `json:"mentioned_profiles"`
    DirectReplies  []Cast       `json:"direct_replies"`
}

type ParentAuthor struct {
    Fid interface{} `json:"fid"`
}

type Author struct {
    Object            string            `json:"object"`
    Fid               uint32               `json:"fid"`
    CustodyAddress    string            `json:"custody_address"`
    Username          string            `json:"username"`
    DisplayName       string            `json:"display_name"`
    PfpURL            string            `json:"pfp_url"`
    Profile           Profile           `json:"profile"`
    FollowerCount     uint32               `json:"follower_count"`
    FollowingCount    uint32               `json:"following_count"`
    Verifications     []string          `json:"verifications"`
    VerifiedAddresses VerifiedAddresses `json:"verified_addresses"`
    ActiveStatus      string            `json:"active_status"`
    PowerBadge        bool              `json:"power_badge"`
}

type Profile struct {
    Bio Bio `json:"bio"`
}

type Bio struct {
    Text string `json:"text"`
}

type VerifiedAddresses struct {
    EthAddresses []string `json:"eth_addresses"`
    SolAddresses []string `json:"sol_addresses"`
}

type Embed struct {
    URL string `json:"url"`
}

type Reactions struct {
    LikesCount   uint32     `json:"likes_count"`
    RecastsCount uint32     `json:"recasts_count"`
    Likes        []User  `json:"likes"`
    Recasts      []User  `json:"recasts"`
}



type Replies struct {
    Count uint32 `json:"count"`
}

type Channel struct {
    Object   string `json:"object"`
    ID       string `json:"id"`
    Name     string `json:"name"`
    ImageURL string `json:"image_url"`
}

type Next struct {
    Cursor string `json:"cursor"`
}



type UserProfile struct {
	Bio struct {
		Text             string   `json:"text"`
		MentionedProfiles []string `json:"mentioned_profiles"`
	} `json:"bio"`
}

type ViewerContext struct {
	Following  bool `json:"following"`
	FollowedBy bool `json:"followed_by"`
}



type User struct {
	Object          string           `json:"object"`
	Fid             int              `json:"fid"`
	Username        string           `json:"username"`
	DisplayName     string           `json:"display_name"`
	CustodyAddress  string           `json:"custody_address"`
	PfpURL          string           `json:"pfp_url"`
	Profile         UserProfile      `json:"profile"`
	FollowerCount   int              `json:"follower_count"`
	FollowingCount  int              `json:"following_count"`
	Verifications   []string         `json:"verifications"`
	VerifiedAddresses VerifiedAddresses `json:"verified_addresses"`
	ActiveStatus    string           `json:"active_status"`
	PowerBadge      bool             `json:"power_badge"`
	ViewerContext   ViewerContext    `json:"viewer_context"`
    Fname           string           `json:"fname"`
}

type Follow struct {
	Object string `json:"object"`
	User   User   `json:"user"`
}



type Reaction struct {
	Object string `json:"object"`
	Cast   Cast   `json:"cast"`
	User   User   `json:"user"`
}

type Notification struct {
	Object              string    `json:"object"`
	MostRecentTimestamp time.Time `json:"most_recent_timestamp"`
	Type                string    `json:"type"`
	Follows             []Follow  `json:"follows"`
	Reactions           []Reaction `json:"reactions"`
}
