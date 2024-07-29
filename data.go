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


type WebhookSecret struct {
    UID       string    `json:"uid"`
    Value     string    `json:"value"`
    ExpiresAt time.Time `json:"expires_at"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"deleted_at"`
}
type Webhook struct {
    Object             string               `json:"object"`
    WebhookID          string               `json:"webhook_id"`
    DeveloperUUID      string               `json:"developer_uuid"`
    TargetURL          string               `json:"target_url"`
    Title              string               `json:"title"`
    Secrets            []WebhookSecret      `json:"secrets"`
    Description        string               `json:"description"`
    HTTPTimeout        string               `json:"http_timeout"`
    RateLimit          int                  `json:"rate_limit"`
    Active             bool                 `json:"active"`
    RateLimitDuration  string               `json:"rate_limit_duration"`
    CreatedAt          time.Time            `json:"created_at"`
    UpdatedAt          time.Time            `json:"updated_at"`
    DeletedAt          time.Time            `json:"deleted_at"`
    Subscription       WebhookSubscription  `json:"subscription,omitempty"`
}

type WebhookSubscription struct {
	Object         string                       `json:"object"`
	SubscriptionID string                       `json:"subscription_id"`
	Filters        WebhookSubscriptionFilters   `json:"filters"`
	CreatedAt      time.Time                    `json:"created_at"`
	UpdatedAt      time.Time                    `json:"updated_at"`
}


type SubscribedTo struct {
	Subscription
	ExpiresAt    string          `json:"expires_at"`
	SubscribedAt string          `json:"subscribed_at"`
	Tier         SubscriptionTier `json:"tier"`
	Creator      User            `json:"creator"`
}


type SubscriptionTier struct {
	ID    int     `json:"id"`
	Price Price   `json:"price"`
}


type Subscriber struct {
	Object       string            `json:"object"`
	User         User              `json:"user"`
	SubscribedTo SubscribedToObject `json:"subscribed_to"`
}

type SubscribedToObject struct {
	Object          string `json:"object"`
	ProviderName    string `json:"provider_name"`
	ContractAddress string `json:"contract_address"`
	Chain           int    `json:"chain"`
	ExpiresAt       string `json:"expires_at"`
	SubscribedAt    string `json:"subscribed_at"`
	TierID          string `json:"tier_id"`
}


type Metadata struct {
	Title  string `json:"title"`
	Symbol string `json:"symbol"`
	ArtURL string `json:"art_url"`
}

type Price struct {
	PeriodDurationSeconds int    `json:"period_duration_seconds"`
	TokensPerPeriod       string `json:"tokens_per_period"`
	InitialMintPrice      string `json:"initial_mint_price"`
}

type Tier struct {
	ID    int   `json:"id"`
	Price Price `json:"price"`
}

type Token struct {
	Symbol   string  `json:"symbol"`
	Address  *string `json:"address"`
	Decimals int     `json:"decimals"`
	ERC20    bool    `json:"erc20"`
}



type Subscription struct {
	Object          string   `json:"object"`
	ProviderName    string   `json:"provider_name"`
	ContractAddress string   `json:"contract_address"`
	Chain           int      `json:"chain"`
	Metadata        Metadata `json:"metadata"`
	OwnerAddress    string   `json:"owner_address"`
	Price           Price    `json:"price"`
	Tiers           []Tier   `json:"tiers"`
	ProtocolVersion int      `json:"protocol_version"`
	Token           Token    `json:"token"`
	ExpiresAt       string   `json:"expires_at"`
	SubscribedAt    string   `json:"subscribed_at"`
	Tier            Tier     `json:"tier"`
	Creator         User  `json:"creator"`
}
type SubscriptionsCreated struct {
	Object                string         `json:"object"`
	SubscriptionsCreated []Subscription `json:"subscriptions_created"`
}


type StorageUsageResponse struct {
	Object           string           `json:"object"`
	User             UserDehydrated   `json:"user"`
	Casts            StorageObject    `json:"casts"`
	Reactions        StorageObject    `json:"reactions"`
	Links            StorageObject    `json:"links"`
	VerifiedAddresses StorageObject   `json:"verified_addresses"`
	UsernameProofs   StorageObject    `json:"username_proofs"`
	Signers          StorageObject    `json:"signers"`
	TotalActiveUnits int              `json:"total_active_units"`
    ErrorResponse
}

type StorageAllocation struct {
	Object    string         `json:"object"`
	User      UserDehydrated `json:"user"`
	Units     int            `json:"units"`
	Expiry    string         `json:"expiry"`
	Timestamp string         `json:"timestamp"`
}

type StorageObject struct {
	Object   string `json:"object"`
	Used     int    `json:"used"`
	Capacity int    `json:"capacity"`
}

type UserDehydrated struct {
	Object string `json:"object"`
	Fid    int32  `json:"fid"`
}

type MuteList struct {
	Object   string `json:"object"`
	Muted    User   `json:"muted"`
	MutedAt  int64  `json:"muted_at"`
}
type  ReactionWithCastInfo struct {
    ReactionType      string    `json:"reaction_type"`
    ReactionTimestamp time.Time `json:"reaction_timestamp"`
    Object            string    `json:"object"`
    User             User `json:"user"`
}
type SignerUUID string

// CastHash represents the cast hash
type CastHash string

// FrameInput represents the input for a frame
type FrameInput struct {
    Text string `json:"text,omitempty"`
}

// FrameState represents the state for a frame
type FrameState struct {
    Serialized string `json:"serialized,omitempty"`
}

// FrameTransaction represents the transaction for a frame
type FrameTransaction struct {
    Hash string `json:"hash,omitempty"`
}

// FrameAddress represents the address for a frame
type FrameAddress struct {
    Address string `json:"address,omitempty"`
}

// FrameActionButton represents a button within a frame
type FrameActionButton struct {
    Title      string `json:"title,omitempty"`
    Index      int    `json:"index"`
    ActionType string `json:"action_type"`
    Target     string `json:"target,omitempty"`
    PostURL    string `json:"post_url,omitempty"`
}

// FrameAction represents an action to be taken on a frame
type FrameAction struct {
    Version     string           `json:"version,omitempty"`
    Title       string           `json:"title,omitempty"`
    Image       string           `json:"image,omitempty"`
    Button      FrameActionButton `json:"button"`
    Input       FrameInput       `json:"input,omitempty"`
    State       FrameState       `json:"state,omitempty"`
    Transaction FrameTransaction `json:"transaction,omitempty"`
    Address     FrameAddress     `json:"address,omitempty"`
    FramesURL   string           `json:"frames_url"`
    PostURL     string           `json:"post_url"`
}




type Frame struct {
    Version            string               `json:"version"`
    Image              string               `json:"image"`
    Buttons            []FrameActionButton  `json:"buttons"`
    PostURL            string               `json:"post_url"`
    FramesURL          string               `json:"frames_url"`
    Title              string               `json:"title,omitempty"`
    ImageAspectRatio   string               `json:"image_aspect_ratio,omitempty"`
    Input              map[string]string    `json:"input,omitempty"`
    State              map[string]string    `json:"state,omitempty"`
}

type ValidatedFrameAction struct {
	Object        string              `json:"object"`
	URL           string              `json:"url"`
	Interactor    User                `json:"interactor"`
	TappedButton  map[string]int      `json:"tapped_button"`
	Input         FrameInput          `json:"input"`
	State         FrameState          `json:"state"`
	Timestamp     time.Time           `json:"timestamp"`
	Signer        struct {
		Client User `json:"client"`
	} `json:"signer"`
	Transaction FrameTransaction `json:"transaction"`
	Address     FrameAddress     `json:"address"`
}

type FrameValidateAnalyticsInteractors struct {
	Interactors []Interactor `json:"interactors"`
}

type Interactor struct {
	FID               uint32  `json:"fid"` 
	Username          string  `json:"username"`
	InteractionCount  float64 `json:"interaction_count"`
}


type FrameValidateAnalyticsTotalInteractors struct {
	TotalInteractors float64 `json:"total_interactors"`
}


type FrameValidateAnalyticsInteractionsPerCast struct {
	InteractionsPerCast []InteractionPerCast `json:"interactions_per_cast"`
}

type InteractionPerCast struct {
	Start            time.Time `json:"start"`
	Stop             time.Time `json:"stop"`
	Time             time.Time `json:"time"`
	InteractionCount float64   `json:"interaction_count"`
	CastURL          string    `json:"cast_url"`
}


type FrameValidateAnalyticsInputText struct {
	InputTexts []FrameInputText `json:"input_texts"`
}

type FrameInputText struct {
	FID        string `json:"fid"` 
	Username   string `json:"username"`
	InputText  string `json:"input_text"`
}