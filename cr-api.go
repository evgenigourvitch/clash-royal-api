package main

import (
	"fmt"
	"github.com/evgenigourvitch/clash-royal-api/crawler"
	"github.com/evgenigourvitch/clash-royal-api/login"
	"github.com/evgenigourvitch/clash-royal-api/objects"
)

func main() {
	credentials, err := objects.NewCredentials()
	if err != nil {
		panic(err)
	}
	loginService := login.NewLoginService(credentials)
	loginObject, err := loginService.Login()
	if err != nil {
		panic(err)
	}
	cr, err := crawler.NewCralwer(loginObject.SwaggerUrl, loginObject.TemporaryAPIToken)
	if err != nil {
		return
	}
/*
	res, err := cr.Request(cr.GetSwaggerUrl()+"clans/%23VGV0RP8/members", objects.EResponseTypePlayersList, 0)
	if err != nil {
		fmt.Printf("got error: %+v\n", err)
		return
	}
	clanMembers := res.(*objects.PlayersResponse)
	cr.GetBattles(clanMembers)
*/
	res, err := cr.Request(cr.GetSwaggerUrl()+"locations", objects.EResponseTypeLocations, 0)
	if err != nil {
		fmt.Printf("got error: %+v\n", err)
		return
	}
	locations := res.(*objects.LocationsResponse)
	cr.GetClansByLocations(locations)
}
