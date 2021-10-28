package tools

import (
	"log"
	"regexp"
	"strings"
	"shadygoat.eu/MatchingAPI/profile"
)

var ecCompiled *regexp.Regexp
var socCompiled *regexp.Regexp

func Prepare() {
	compiled, err := regexp.Compile(`ec=\d+?.\d+?`)
	if err != nil {
		log.Fatal(err)
	}
	ecCompiled = &compiled
	
	compiled, err := regexp.Compile(`soc=\d+?.\d+?`)
	if err != nil {
		log.Fatal(err)
	}
	socCompiled = &compiled
	
}

func GetBDSMResults() {

}

func GetPoliticalResults(url string) profile.Political {
	// ec=10&soc=10
	query := strings.Replace(url, "https://www.politicalcompass.org/analysis2?", "", 1)
	for _, k := range query {1
		if (ecCompiled.FindString(query)) {
			// TODO: FIXME:
		}
	}
}