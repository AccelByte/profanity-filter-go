// Copyright (c) 2020 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var list = map[string]string{
	"arabic":     "ar.json",
	"chinese":    "zh.json",
	"czech":      "cs.json",
	"danish":     "da.json",
	"dutch":      "nl.json",
	"english":    "en.json",
	"esperanto":  "eo.json",
	"filipino":   "fil.json",
	"finnish":    "fi.json",
	"french":     "fr.json",
	"frenchCA":   "fr-CA-u-sd-caqc.json",
	"german":     "de.json",
	"hindi":      "hi.json",
	"hungarian":  "hu.json",
	"italian":    "it.json",
	"japanese":   "ja.json",
	"kabyle":     "kab.json",
	"klingon":    "tlh.json",
	"korean":     "ko.json",
	"norwegian":  "no.json",
	"persian":    "fa.json",
	"polish":     "pl.json",
	"portuguese": "pt.json",
	"russian":    "ru.json",
	"spanish":    "es.json",
	"swedish":    "sv.json",
	"thai":       "th.json",
	"turkish":    "tr.json",
}

const (
	profanityDirectory = "./profanitylist/naughty-words-js-master/"
)

func readProfanityFromJSONFile(fileName string) ([]string, error) {
	dataJSON, err := ioutil.ReadFile(profanityDirectory + fileName)
	if err != nil {
		return nil, err
	}

	var profanityList []string
	err = json.Unmarshal(dataJSON, &profanityList)

	if err != nil {
		return nil, err
	}

	return profanityList, nil
}

func main() {
	for k, v := range list {
		profanityList, err := readProfanityFromJSONFile(v)
		if err != nil {
			fmt.Println(err.Error())
		}

		s := fmt.Sprintf(`
		var %sProfanityList = profanityList{
			Country:     "%s",
			BadWordList: []string{"%s"},
		}`, k, k, strings.Join(profanityList, `", "`))

		fmt.Println(s)
	}
}
