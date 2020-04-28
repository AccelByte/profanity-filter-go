// Copyright (c) 2020 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package profanityfilter

import (
	"encoding/json"
	"io/ioutil"

	swearfilter "github.com/JoshuaDoes/gofuckyourself"
)

const (
	ArabicProfanityList     = "ar.json"
	ChineseProfanityList    = "zh.json"
	CzechProfanityList      = "cs.json"
	DanishProfanityList     = "da.json"
	DutchProfanityList      = "nl.json"
	EnglishProfanityList    = "en.json"
	EsperantoProfanityList  = "eo.json"
	FilipinoProfanityList   = "fil.json"
	FinnishProfanityList    = "fi.json"
	FrenchProfanityList     = "fr.json"
	FrenchCAProfanityList   = "fr-CA-u-sd-caqc.json"
	GermanProfanityList     = "de.json"
	HindiProfanityList      = "hi.json"
	HungarianProfanityList  = "hu.json"
	ItalianProfanityList    = "it.json"
	JapaneseProfanityList   = "ja.json"
	KabyleProfanityList     = "kab.json"
	KlingonProfanityList    = "tlh.json"
	KoreanProfanityList     = "ko.json"
	NorwegianProfanityList  = "no.json"
	PersianProfanityList    = "fa.json"
	PolishProfanityList     = "pl.json"
	PortugueseProfanityList = "pt.json"
	RussianProfanityList    = "ru.json"
	SpanishProfanityList    = "es.json"
	SwedishProfanityList    = "sv.json"
	ThaiProfanityList       = "th.json"
	TurkishProfanityList    = "tr.json"

	profanityDirectory = "./profanitylist/naughty-words-js-master/"
)

func readProfanityFromJsonFile(fileName string) ([]string, error) {
	dataJson, err := ioutil.ReadFile(profanityDirectory + fileName)
	if err != nil {
		return nil, err
	}

	var profanityList []string
	err = json.Unmarshal(dataJson, &profanityList)
	if err != nil {
		return nil, err
	}
	return profanityList, nil
}

// ReadAllCountryProfanityList returns list of bad words from all countries
func ReadAllCountryProfanityList() ([]string, error) {
	listProfanityCountry := []string{ArabicProfanityList, ChineseProfanityList, CzechProfanityList, DanishProfanityList, DutchProfanityList, EnglishProfanityList, EsperantoProfanityList, FilipinoProfanityList, FinnishProfanityList, FrenchProfanityList, FrenchCAProfanityList, GermanProfanityList, HindiProfanityList, HungarianProfanityList, ItalianProfanityList, JapaneseProfanityList, KabyleProfanityList, KlingonProfanityList, KoreanProfanityList, NorwegianProfanityList, PersianProfanityList, PolishProfanityList, PortugueseProfanityList, RussianProfanityList, SpanishProfanityList, SwedishProfanityList, ThaiProfanityList, TurkishProfanityList}
	var profanityListAllCountry []string
	for _, country := range listProfanityCountry {
		profanityList, err := readProfanityFromJsonFile(country)
		if err != nil {
			return nil, err
		}
		profanityListAllCountry = append(profanityListAllCountry, profanityList...)
	}
	return profanityListAllCountry, nil
}

// ProfanityCheck checks if a message contains blacklisted words
func ProfanityCheck(word string) (bool, []string, error) {
	swears, err := ReadAllCountryProfanityList()
	if err != nil {
		return false, nil, err
	}
	filter := swearfilter.New(true, false, false, false, false, swears...)
	isSwearFound, swearsFound, err := filter.Check(word)

	return isSwearFound, swearsFound, err
}
