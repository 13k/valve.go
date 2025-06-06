package api

import (
	"github.com/13k/valve.go/steamweb/schema"
)

// ISteamWebAPIUtilGetServerInfo can be used as result of method ISteamWebAPIUtil/GetServerInfo.
type ISteamWebAPIUtilGetServerInfo struct {
	ServerTime       uint64 `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

// ISteamWebAPIUtilGetSupportedAPIList can be used as result of method
// ISteamWebAPIUtil/GetSupportedAPIList.
type ISteamWebAPIUtilGetSupportedAPIList struct {
	API *schema.Schema `json:"apilist"`
}
