package config

func GetNegativeExtractions() map[string]string {
	return map[string]string{
		"url\":\"string\"":                   "equals",
		"\"Basic Information\"":              "equals",
		"\"Basic Navigation\"":               "equals",
		"'common.words.password':'Password'": "equals",
		">detectmobilebrowsers.com<":         "contains",
		"/mobile-detect.js":                  "contains",
		"http://mysite.com/":                 "contains",
	}

}
