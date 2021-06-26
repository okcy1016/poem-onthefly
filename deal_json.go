package main

import (
	"encoding/json"
)

type ConvCfg struct {

	Interval string
	Font string
	Fontsize string
	Imgloc string
	Color string
	XOffset string
	YOffset string
	IfDae bool
}

func JsonStrToStru(convCfgStr string) ConvCfg {

	var convCfg ConvCfg
	
	convCfgByte := []byte(convCfgStr)

	err := json.Unmarshal(convCfgByte, &convCfg)
	ChkErr(err)

	return convCfg
}
