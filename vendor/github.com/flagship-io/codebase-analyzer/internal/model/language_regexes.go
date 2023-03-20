package model

import (
	"encoding/json"
	"log"
)

type LanguageRegex struct {
	FileExtension string   `json:"file_extension"`
	Regexes       []string `json:"regexes"`
}

var LanguageRegexes = []LanguageRegex{
	{
		FileExtension: `\.[jt]sx?$`,
		Regexes: []string{
			`useFsFlag[(](?:(?:\s*['"](.*)['"]\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                   // SDK React V3
			`['"]?key['"]?\s*\:\s*['"](.+?)['"](?:.*\s*)['"]?defaultValue['"]?\s*\:\s*(['"].*['"]|[^\r\n\t\f\v,}]+).*[},]?`, // SDK JS V2 && SDK React V2
			`getFlag[(](?:(?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)])?`,                                     // SDK JS V3
		},
	},

	{
		FileExtension: `\.go$`,
		Regexes: []string{
			`\.GetModification(?:String|Number|Bool|Object|Array)\s*(?:\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\))?`, // SDK GO V2
		},
	},
	{
		FileExtension: `\.py$`,
		Regexes: []string{
			`\.get_modification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|True|False|\d+|"[^"]*"))?\s*\)`, // SDK PYTHON V2
		},
	},
	{
		FileExtension: `\.java$`,
		Regexes: []string{
			`\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK JAVA V2
			`\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK JAVA V3
		},
	},
	{
		FileExtension: `\.php$`,
		Regexes: []string{
			`\-\>getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK PHP V1 && SDK PHP V2
			`\-\>getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK PHP V3
		},
	},
	{
		FileExtension: `\.kt$`,
		Regexes: []string{
			`\.getModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK ANDROID V2
			`\.getFlag[(](?:\s*["'](.*)["']\s*,\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                             // SDK ANDROID V3

		},
	},
	{
		FileExtension: `\.swift$`,
		Regexes: []string{
			`\.getModification\(\s*["'](\w+)['"]\s*,\s*default(?:String|Double|Float|Int|Bool|Json|Array)\s*:\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)\s*(?:,\s*activate\s*:\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK iOS V2
			`\.getFlag[(]\s*key\s*:\s*(?:\s*["'](.*)["']\s*,\s*defaultValue\s*:\s*(["'].*\s*[^"]*["']|[^)]*))\s*[)]`,                                                                                                                 // SDK iOS V3
		},
	},
	{
		FileExtension: `\.m$`,
		Regexes: []string{
			`getModification\s*:\s*@\s*['"](.+?)['"](?:\s*)default(?:String|Double|Bool|Float|Int|Json|Array):\@?\s*(['"].+?['"]|YES|NO|TRUE|FALSE|true|false|[+-]?(?:\d*[.])?\d+)?`, // SDK iOS V2
			`getFlagWithKey\s*:\s*\@['"](.+?)['"](?:\s*)['"]?defaultValue['"]?\s*\:\s*\@?\s*(.+?)\s*[\]]`,                                                                            // SDK iOS V3
		},
	},
	{
		FileExtension: `\.[fc]s$`,
		Regexes: []string{
			`\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			`\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,         // SDK .NET V3
		},
	},
	{
		FileExtension: `\.vb$`,
		Regexes: []string{
			`\.GetModification\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|True|false|False|\d+|"[^"]*"))?\s*\)`, // SDK .NET V1
			`\.GetFlag\(\s*["']([\w\-]+)['"]\s*,\s*(["'][^"]*['"]|[+-]?(?:\d*[.])?\d+|true|false|False|True)(?:\s*,\s*(?:true|false|\d+|"[^"]*"))?\s*\)`,                    // SDK .NET V3
		},
	},
}

func AddCustomRegexes(customRegexJSON string) {
	regexes := []LanguageRegex{}
	err := json.Unmarshal([]byte(customRegexJSON), &regexes)

	if err != nil {
		log.Printf("Error when parsing custom regexes : %v", err)
		return
	}

	LanguageRegexes = append(LanguageRegexes, regexes...)
}
