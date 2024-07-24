package neynarsdk



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
}

type ParentAuthor struct {
    Fid interface{} `json:"fid"`
}

type Author struct {
    Object            string            `json:"object"`
    Fid               int               `json:"fid"`
    CustodyAddress    string            `json:"custody_address"`
    Username          string            `json:"username"`
    DisplayName       string            `json:"display_name"`
    PfpURL            string            `json:"pfp_url"`
    Profile           Profile           `json:"profile"`
    FollowerCount     int               `json:"follower_count"`
    FollowingCount    int               `json:"following_count"`
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
    LikesCount   int     `json:"likes_count"`
    RecastsCount int     `json:"recasts_count"`
    Likes        []User  `json:"likes"`
    Recasts      []User  `json:"recasts"`
}

type User struct {
    Fid   int    `json:"fid"`
    Fname string `json:"fname"`
}

type Replies struct {
    Count int `json:"count"`
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