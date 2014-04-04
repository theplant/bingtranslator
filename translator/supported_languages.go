package bingtranslator

/*
SupportedLanguages returns a map of codes to names of all supported languages.
NB: this is a static list - does currently not use the translator API
*/
func SupportedLanguages() map[string]string {
	return supportedLanguages
}

var supportedLanguages = map[string]string{
	"ar":       "Arabic -- العربية",
	"bg":       "Bulgarian -- български",
	"ca":       "Catalan -- Català",
	"zh-chs":   "Chinese Simplified -- 简体中文",
	"zh-cht":   "Chinese Traditional -- 繁體中文",
	"cs":       "Czech -- čeština",
	"da":       "Danish -- dansk",
	"nl":       "Dutch -- Nederlands",
	"en":       "English",
	"et":       "Estonian -- eesti",
	"fi":       "Finnish -- suomalainen",
	"fr":       "French -- français",
	"de":       "German -- Deutsch",
	"el":       "Greek -- ελληνικά",
	"ht":       "Haitian Creole -- kreyòl ayisyen",
	"he":       "Hebrew -- עברית",
	"hi":       "Hindi -- हिंदी",
	"mww":      "Hmong Daw -- Hmong Daw",
	"hu":       "Hungarian -- magyar",
	"id":       "Indonesian -- Indonesia",
	"it":       "Italian -- italiano",
	"ja":       "Japanese -- 日本語",
	"tlh":      "Klingon",
	"tlh-qaak": "Klingon (pIqaD)",
	"ko":       "Korean -- 한국의",
	"lv":       "Latvian -- Latvijas",
	"lt":       "Lithuanian -- Lietuvos",
	"ms":       "Malay -- Melayu",
	"mt":       "Maltese -- Malti",
	"no":       "Norwegian -- norsk",
	"fa":       "Persian -- فارسی",
	"pl":       "Polish -- polski",
	"pt":       "Portuguese -- português",
	"ro":       "Romanian -- român",
	"ru":       "Russian -- русский",
	"sk":       "Slovak -- slovenský",
	"sl":       "Slovenian -- slovenščina",
	"es":       "Spanish -- español",
	"sv":       "Swedish -- Svenska",
	"th":       "Thai -- ภาษาไทย",
	"tr":       "Turkish -- Türk",
	"uk":       "Ukrainian -- Український",
	"ur":       "Urdu -- اردو",
	"vi":       "Vietnamese -- Việt",
}
