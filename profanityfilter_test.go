// Copyright (c) 2020 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package profanityfilter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	profanityfilter "accelbyte.net/profanity-filter-go"
)

func Test_ProfanityCheck(t *testing.T) {
	inputs := map[string]string{
		"ArabicBadWord":        "سكس",
		"ChineseBadWord":       "三级片",
		"CzechBadWord":         "buzna",
		"DanishBadWord":        "bøsserøv",
		"DutchBadWord":         "afberen",
		"EnglishBadWord":       "asshole",
		"EsperantoBadWord":     "bugri",
		"FilipinoBadWord":      "jakol",
		"FinnishBadWord":       "hevonvitunperse",
		"FrenchBadWord":        "baiser",
		"FrenchCABadWord":      "noune",
		"GermanBadWord":        "arschlecker",
		"HindiBadWord":         "bhangi",
		"HungarianBadWord":     "baszik",
		"ItalianBadWord":       "battona",
		"JapaneseBadWord":      "エロティズム",
		"KabyleBadWord":        "iwellaqen",
		"KlingonBadWord":       "QI'yaH",
		"KoreanBadWord":        "계집년",
		"NorwegianBadWord":     "forbanna",
		"PersianBadWord":       "پورنو",
		"PolishBadWord":        "jajko",
		"PortugueseBadWord":    "chochota",
		"RussianBadWord":       "perdet",
		"SpanishBadWord":       "Cabrón",
		"SwedishBadWord":       "olla",
		"ThaiBadWord":          "เหี้ย",
		"TurkishProfanBadWord": "Çingeneye"}

	for i, v := range inputs {
		t.Run("Test_"+i, func(t *testing.T) {
			isSwearWord, swearFound, err := profanityfilter.Filter.ProfanityCheck(v)
			if err != nil {
				t.Fatal(err.Error())
			}
			assert.True(t, isSwearWord, i)
			assert.NotNil(t, swearFound, i)
			assert.Nil(t, err, i)
		})
	}
}
