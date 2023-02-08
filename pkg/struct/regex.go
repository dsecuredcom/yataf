package _struct

type RegExConfig struct {
	Elements []RegEx
}

type RegEx struct {
	Name    string
	Matcher string
	Type    string
}
