package main

import (
	"io/ioutil"
	"strings"
)


func saveCfgToDisk(convCfgStr string) {

	// Add some keywords make it as javascript syntax
	convCfgStr = "cfg = " + convCfgStr + ";"
	
	convCfgByte := []byte(convCfgStr)

	err := ioutil.WriteFile("webview_ui/js/config.js", convCfgByte, 0644)
	ChkErr(err)
}

func loadCfgFromDisk() string {

	convCfgByte, err := ioutil.ReadFile("webview_ui/js/config.js")
	ChkErr(err)

	convCfgStr := string(convCfgByte)

	// Remove the javascript keywords
	convCfgStr = convCfgStr[6:]

	// Whether there is a end buffer(Emacs) in the config.js, the ";" should be cut
	convCfgStr = strings.TrimSuffix(convCfgStr, ";\n")
	convCfgStr = strings.TrimSuffix(convCfgStr, ";")
	
	return convCfgStr
}
