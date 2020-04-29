// Copyright (c) 2020 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package profanityfilter

import (
	swearfilter "github.com/JoshuaDoes/gofuckyourself"
)

// Filter variable that used for profanity checking
var Filter *profanityFilter

type profanityFilter struct {
	SwearFilter   *swearfilter.SwearFilter
	ProfanityList []string
}

func init() {
	if Filter == nil {
		badWords := getListOfBadwords()
		swearFilter := swearfilter.New(true, false, false, false, false, badWords...)

		Filter = &profanityFilter{
			SwearFilter:   swearFilter,
			ProfanityList: badWords,
		}
	}
}

// ProfanityCheck checks if a message contains blacklisted words
func (filter *profanityFilter) ProfanityCheck(word string) (bool, []string, error) {
	isSwearFound, swearsFound, err := filter.SwearFilter.Check(word)

	return isSwearFound, swearsFound, err
}
