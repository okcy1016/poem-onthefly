package main

import (
	"time"
	"fmt"
)


func StartSchedule(interval int, convCfg ConvCfg) {

	WelcomeMsg()
	
	for {
		// Check network connection firstly
		ChkNetConn()
		
		poeLoc, err := FetchPoetry(GenPoeUrl())
		ChkErr(err)

		poeStr := ParseJsonPoe(poeLoc)
		fmt.Println(poeStr)

		picLoc := GenPic(poeStr, convCfg)

		SetWallPaper(picLoc, interval)

		time.Sleep(time.Duration(interval) * time.Hour)
	}
}
