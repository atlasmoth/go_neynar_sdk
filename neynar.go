package neynar

import (
	hub "github.com/atlasmoth/go_neynar_sdk/hub"
	stp "github.com/atlasmoth/go_neynar_sdk/stp"
	v1 "github.com/atlasmoth/go_neynar_sdk/v1"
	v2 "github.com/atlasmoth/go_neynar_sdk/v2"
)

func NewV1Client(cfg v1.Configuration) v1.APIClient{
	return *v1.NewAPIClient(&cfg)
}

func NewV2Client(cfg v2.Configuration) v2.APIClient{
	return *v2.NewAPIClient(&cfg)
}

func NewHubClient(cfg hub.Configuration) hub.APIClient{
	return *hub.NewAPIClient(&cfg)
}

func NewSTPClient(cfg stp.Configuration) stp.APIClient{
	return *stp.NewAPIClient(&cfg)
}


