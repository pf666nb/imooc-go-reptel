package parser

import (
	"cralwer/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		user := string(m[2])
		result.Items = append(result.Items, "User "+user)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(
				c []byte) engine.ParseResult {
				return ParseProfile(c, user)
			},
		})

	}
	return result

}
