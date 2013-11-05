bingtranslator
==============

golang wrapper for the Translate and TranslateArray methods of the bing Translator API

Usage:
------

Import the package

```
import (
	"fmt"
	btr "github.com/theplant/bingtranslator/translator"
)
```

Send two strings for translation from english to german:

```
// first you need to set your ClientId and ClientSecret
btr.SetCredentials(clientId, clientSecret)

// multi html
texts := []string{
	"<li>Start writing an entry<br></li>",
	"<li>Qortex continuously update the X most “interesting” words as you write<br></li>",
}
translations, err = btr.Translate("en", "DE", texts, btr.INPUT_HTML)
if err != nil {
	fmt.Println(err)
}
fmt.Println(translations)
```

You can send single or multiple texts - as plain `btr.INPUT_TEXT` text or as html `btr.INPUT_HTML`.
Check [main.go](https://github.com/theplant/bingtranslator/blob/master/main.go)

`btr.SupportedLanguages()` returns a mapping of language codes and names of all supported languages:

```
ar:       Arabic
bg:       Bulgarian
ca:       Catalan
zh-CHS:   Chinese Simplified
zh-CHT:   Chinese Traditional
cs:       Czech
da:       Danish
nl:       Dutch
en:       English
et:       Estonian
fi:       Finnish
fr:       French
de:       German
el:       Greek
ht:       Haitian Creole
he:       Hebrew
hi:       Hindi
mww:      Hmong Daw
hu:       Hungarian
id:       Indonesian
it:       Italian
ja:       Japanese
tlh:      Klingon
tlh-Qaak: Klingon (pIqaD)
ko:       Korean
lv:       Latvian
lt:       Lithuanian
ms:       Malay
mt:       Maltese
no:       Norwegian
fa:       Persian
pl:       Polish
pt:       Portuguese
ro:       Romanian
ru:       Russian
sk:       Slovak
sl:       Slovenian
es:       Spanish
sv:       Swedish
th:       Thai
tr:       Turkish
uk:       Ukrainian
ur:       Urdu
vi:       Vietnamese
```
