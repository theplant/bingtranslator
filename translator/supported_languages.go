package bingtranslator

/*
SupportedLanguages returns a map of codes to names of all supported languages.
NB: this is a static list - does currently not use the translator API
*/
func SupportedLanguages() map[string]string {
	return supportedLanguages
}

var supportedLanguages = map[string]string{
	"ar":       "Arabic",
	"bg":       "Bulgarian",
	"ca":       "Catalan",
	"zh-chs":   "Chinese Simplified",
	"zh-cht":   "Chinese Traditional",
	"cs":       "Czech",
	"da":       "Danish",
	"nl":       "Dutch",
	"en":       "English",
	"et":       "Estonian",
	"fi":       "Finnish",
	"fr":       "French",
	"de":       "German",
	"el":       "Greek",
	"ht":       "Haitian Creole",
	"he":       "Hebrew",
	"hi":       "Hindi",
	"mww":      "Hmong Daw",
	"hu":       "Hungarian",
	"id":       "Indonesian",
	"it":       "Italian",
	"ja":       "Japanese",
	"tlh":      "Klingon",
	"tlh-qaak": "Klingon (pIqaD)",
	"ko":       "Korean",
	"lv":       "Latvian",
	"lt":       "Lithuanian",
	"ms":       "Malay",
	"mt":       "Maltese",
	"no":       "Norwegian",
	"fa":       "Persian",
	"pl":       "Polish",
	"pt":       "Portuguese",
	"ro":       "Romanian",
	"ru":       "Russian",
	"sk":       "Slovak",
	"sl":       "Slovenian",
	"es":       "Spanish",
	"sv":       "Swedish",
	"th":       "Thai",
	"tr":       "Turkish",
	"uk":       "Ukrainian",
	"ur":       "Urdu",
	"vi":       "Vietnamese",
}
