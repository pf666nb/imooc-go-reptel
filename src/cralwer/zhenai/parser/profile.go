package parser

import (
	"cralwer/engine"
	"cralwer/model"
	"fmt"
	"github.com/Andrew-M-C/go.jsonvalue"
	"regexp"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">([^<]+)</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">([^<]+)座</div>`)
var HeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">([^<]+)cm</div>`)
var WeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">([^<]+)kg</div>`)
var OccupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c="">工作地:([^<]+)</div>`)
var testRe = regexp.MustCompile(`window\.__INITIAL_STATE__=(.*?);`)

func ParseProfile(contents []byte,
	name string) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = name

	json_str := extractString(contents, testRe)

	j := jsonvalue.MustUnmarshal(json_str)

	itemArray, err := j.GetArray("objectInfo", "basicInfo")
	if err != nil {
		fmt.Println(err)
	} else {
		profile.Info = itemArray.String()
	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) []byte {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return match[1]
	} else {
		return []byte{}
	}
}
