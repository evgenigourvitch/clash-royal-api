package main

import (
	"fmt"
	"github.com/evgenigourvitch/clash-royal-api/crawler"
	"github.com/evgenigourvitch/clash-royal-api/login"
	"github.com/evgenigourvitch/clash-royal-api/objects"
)

const (
	clanTag = "%2329VRJ9Q" /* СССР за 30 лет */
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
//	res, err := cr.Request(cr.GetSwaggerUrl()+"clans/" + clanTag + "/members", objects.EResponseTypePlayersList, 0)
	res, err := cr.Request(cr.GetSwaggerUrl()+"clans/" + clanTag + "/warlog", objects.EResponseTypeClansWarLog, 0)
	if err != nil {
		fmt.Printf("got error: %+v\n", err)
		return
	}
	wars, ok := res.(*objects.WarsResponse)
	if !ok {
		fmt.Printf("failed to cast to *objects.WarsResponse\n")
		return
	}
	fmt.Printf("%+v\n", wars)
//	cr.GetBattles(clanMembers)
	/*
	return
	res, err := cr.Request(cr.GetSwaggerUrl()+"locations", objects.EResponseTypeLocations, 0)
	if err != nil {
		fmt.Printf("got error: %+v\n", err)
		return
	}
	locations := res.(*objects.LocationsResponse)
	cr.GetClansByLocations(locations)
	*/
}
