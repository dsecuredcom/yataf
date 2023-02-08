package config

import (
	_struct "github.com/Damian89/FileAnalyzer/pkg/struct"
	"strings"
)

func GetRegExConfig(Types []string) _struct.RegExConfig {
	Config := _struct.RegExConfig{
		Elements: []_struct.RegEx{
			{Type: "Credentials", Name: "User/Pass Case#1", Matcher: "(?im)['|\"](user|username|mail|email|access|ident|login)['|\"](\\s*?):(\\s*?)['|\"][\\S\\s]+[\\S]+['|\"](\\s*?),(\\s*?)['|\"](pass|password|passwordpassive|authkey|credentials|access_key|apikey|secret|access)['|\"](\\s*?):(\\s*?)['|\"][\\S\\s]+[\\S]+['|\"]"},
			{Type: "Credentials", Name: "Slack Token", Matcher: `(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})`},
			{Type: "Credentials", Name: "RSA private key", Matcher: `-----BEGIN RSA PRIVATE KEY-----`},
			{Type: "Credentials", Name: "SSH (DSA) private key", Matcher: `-----BEGIN DSA PRIVATE KEY-----`},
			{Type: "Credentials", Name: "SSH (EC) private key", Matcher: `-----BEGIN EC PRIVATE KEY-----`},
			{Type: "Credentials", Name: "PGP private key block", Matcher: `-----BEGIN PGP PRIVATE KEY BLOCK-----`},
			{Type: "Credentials", Name: "Open SSH Private Key", Matcher: `-----BEGIN OPENSSH PRIVATE KEY-----`},
			{Type: "Credentials", Name: "Amazon Related Key ID", Matcher: `AKIA[0-9A-Z]{16}`},
			{Type: "Credentials", Name: "Amazon MWS Auth Token", Matcher: `amzn\\.mws\\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`},
			{Type: "Credentials", Name: "Amazon Secret Ident", Matcher: `(?i)aws_secret_access_key|aws_secret_key|aws\.secret_key|aws\.secret_access_key`},
			{Type: "Credentials", Name: "Amazon Secret Key", Matcher: `(?im)['|\"]aws_secret_access_key['|\"](\\s*?):(\\s*?)['|\"][\\S\\s]+[\\S]+['|\"]`},
			{Type: "Credentials", Name: "Facebook Access Token", Matcher: `EAACEdEose0cBA[0-9A-Za-z]+`},
			{Type: "Credentials", Name: "Facebook OAuth", Matcher: `[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K]['|\"](\s*?):(\s*?)['|\"][0-9a-f]{32}['|\"]`},
			{Type: "Credentials", Name: "GitHub", Matcher: `[g|G][i|I][t|T][h|H][u|U][b|B]['|\"](\s*?):(\s*?)['|\"][0-9a-zA-Z]{35,40}['|\"]`},
			{Type: "Credentials", Name: "Generic API Key", Matcher: `[a|A][p|P][i|I][_]?[k|K][e|E][y|Y]['|\"](\s*?):(\s*?)['|\"][0-9a-zA-Z]{32,45}['|\"]`},
			{Type: "Credentials", Name: "Generic Secret #1", Matcher: `[s|S][e|E][c|C][r|R][e|E][t|T]['|\"](\s*?):(\s*?)['|\"][0-9a-zA-Z-=_.]{16,45}['|\"]`},
			{Type: "Credentials", Name: "Generic Secret #2", Matcher: `[s|S][e|E][c|C][r|R][e|E][t|T](\s*?)=(\s*?)['|\"][0-9a-zA-Z-=_.]{16,45}['|\"]`},
			{Type: "Credentials", Name: "Google Related Key", Matcher: `AIza[0-9A-Za-z\\-_]{35}`},
			{Type: "Credentials", Name: "Google OAuth", Matcher: `[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`},
			{Type: "Credentials", Name: "Google (GCP) Service-account", Matcher: `\"type\": \"service_account\"`},
			{Type: "Credentials", Name: "Google OAuth Access Token", Matcher: `ya29\\.[0-9A-Za-z\\-_]+`},
			{Type: "Credentials", Name: "Heroku API Key", Matcher: `[h|H][e|E][r|R][o|O][k|K][u|U].*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`},
			{Type: "Credentials", Name: "MailChimp API Key", Matcher: `[0-9a-f]{32}-us[0-9]{1,2}`},
			{Type: "Credentials", Name: "Mailgun API Key", Matcher: `key-[0-9a-zA-Z]{32}`},
			{Type: "Credentials", Name: "Password in URL", Matcher: `[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@.{1,100}[\"'\\s]`},
			{Type: "Credentials", Name: "PayPal Braintree Access Token", Matcher: `access_token\\$production\\$[0-9a-z]{16}\\$[0-9a-f]{32}`},
			{Type: "Credentials", Name: "Picatic API Key", Matcher: `sk_live_[0-9a-z]{32}`},
			{Type: "Credentials", Name: "Slack Webhook", Matcher: `https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`},
			{Type: "Credentials", Name: "Potential Slack related #1", Matcher: `slack[_|.|-](access|api|secret|token)[_|.|-](token|key|id)`},
			{Type: "Credentials", Name: "Stripe API Key", Matcher: `sk_live_[0-9a-zA-Z]{24}`},
			{Type: "Credentials", Name: "Stripe Restricted API Key", Matcher: `rk_live_[0-9a-zA-Z]{24}`},
			{Type: "Credentials", Name: "Square Access Token", Matcher: `sq0atp-[0-9A-Za-z\\-_]{22}`},
			{Type: "Credentials", Name: "Square OAuth Secret", Matcher: `sq0csp-[0-9A-Za-z\\-_]{43}`},
			{Type: "Credentials", Name: "Twilio API Key", Matcher: `SK[0-9a-fA-F]{32}`},
			{Type: "Credentials", Name: "Twitter Access Token", Matcher: `[t|T][w|W][i|I][t|T][t|T][e|E][r|R]['|\"](\s*?):(\s*?)[1-9][0-9]+-[0-9a-zA-Z]{40}`},
			{Type: "Credentials", Name: "Twitter OAuth", Matcher: `[t|T][w|W][i|I][t|T][t|T][e|E][r|R]['|\"](\s*?):(\s*?)['|\"][0-9a-zA-Z]{35,44}['|\"]`},
			{Type: "Credentials", Name: "Shared Access Key", Matcher: `(?i)SharedAccessKey\=`},
			{Type: "Credentials", Name: "Authorization Header #1", Matcher: `(?i)["|'](Basic|Bearer|Digest|HOBA|Mutual|Negotiate|OAuth|SCRAM-SHA-1|SCRAM-SHA-256|vapid) [a-zA-Z0-9\\\-\._~\+\/]+=*["|']`},
			{Type: "Credentials", Name: "Authorization Header #2", Matcher: `setHeader\(["|']Authorization["|']`},
			{Type: "Credentials", Name: "Authorization Header #3", Matcher: `(?i)Authorization:\s*(Basic|Bearer|Digest|HOBA|Mutual|Negotiate|OAuth|SCRAM-SHA-1|SCRAM-SHA-256|vapid) [a-zA-Z0-9\\\-\._~\+\/=]+`},
			{Type: "Credentials", Name: "Seen in the past tokens", Matcher: `(?i)['|"](DISCOVERY_IAM_APIKEY|appPassword|slackToken|slack_signing_secret|watson_assistant_api_key|pythonPassword)['|"]`},
			{Type: "Credentials", Name: "Secret indicator with .", Matcher: `(?i)['|"][a-zA-Z-\.]+\.(access-key|password|apikey|secret|access_key|secret-key|pwd|passwd|appsecret|app_secret)['|"](\s*?):(\s*?)['|"][\S\s]+[\S]+['|"](\s*?)`},
			{Type: "Credentials", Name: "Secret indicator with _", Matcher: `(?i)['|"][a-zA-Z-\.]+_(access-key|password|apikey|secret|access_key|secret-key|pwd|passwd|appsecret|app_secret)['|"](\s*?):(\s*?)['|"][\S\s]+[\S]+['|"](\s*?)`},
			{Type: "Credentials", Name: "Credentials in URL", Matcher: `(?i)['|"](http|https|ftp|postgresql|smtp|ssh|mongodb|redis|mysql|mssql|jdbc|odbc)://[a-zA-Z0-9\_-]+:.*@[a-z0-9]+`},

			{Type: "Urls-Paths", Name: "Cloudinary", Matcher: `cloudinary://.*`},
			{Type: "Urls-Paths", Name: "Firebase URL", Matcher: `[a-zA-Z0-9-]+\.firebaseio\.com`},
			{Type: "Urls-Paths", Name: "Urls and paths #1", Matcher: `(?i)[a-zA-Z]\.(get|post|fetch|patch|delete|option|put|ajax)\((\s*?)['|"][a-zA-Z0-9-_\/:\?]['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #2", Matcher: `(?i)<(a|script).*(href|src)=['|"]\/[a-zA-Z0-9]+.*['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #3", Matcher: `(?i)<(a|script).*href=['|"]http[s]?:\/\/[a-zA-Z0-9\.\-\/?=]+['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #4", Matcher: `(?i)<(a|script).*href=['|"]\/\/[a-zA-Z0-9\.\-\/?=]+['|"].*>.*<\/(a|script)>`},
			{Type: "Urls-Paths", Name: "Urls and paths #5", Matcher: `(?i)<base.*href=['|"]\/.*\/>`},
			{Type: "Urls-Paths", Name: "Urls and paths #6", Matcher: `(?i)(url|path|file)['|"](\s*?):(\s*?)['|"].*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #8", Matcher: `(?im)url: ['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #9", Matcher: `(?im)url: [a-zA-Z0-9]+(\s*?)\+(\s*?)['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #10", Matcher: `(?im)(axios|jquery|\$).(get|post|delete|put|ajax)\(['|"]\/.*?['|"]`},
			{Type: "Urls-Paths", Name: "Urls and paths #11", Matcher: `(?im)(axios|jquery|\$).(get|post|delete|put|ajax)\([a-zA-Z0-9._]+,`},
			{Type: "Urls-Paths", Name: "Urls and paths #12", Matcher: `(?i)fetch\(['|"]\/.*?['|"],`},
			{Type: "Urls-Paths", Name: "Urls and paths #13", Matcher: `(?im)['|"][a-zA-Z0-9_\.]+['|"](\s*?):(\s*?)['|"](http[s]?:|//|\./).*?['|"]`},
		},
	}

	if IsTypeAllowed(Types, "all") {
		return Config
	}

	FilteredConfig := _struct.RegExConfig{
		Elements: []_struct.RegEx{},
	}

	for _, Element := range Config.Elements {
		DesiredType := strings.ToLower(Element.Type)
		if IsTypeAllowed(Types, DesiredType) {
			FilteredConfig.Elements = append(FilteredConfig.Elements, Element)
		}
	}

	return FilteredConfig
}

func IsTypeAllowed(Types []string, DesiredType string) bool {

	for _, Type := range Types {
		Type = strings.ToLower(Type)
		DesiredType = strings.TrimSpace(DesiredType)
		if Type == DesiredType {
			return true
		}
	}
	return false
}
