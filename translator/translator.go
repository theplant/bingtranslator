package bingtranslator

import (
	"errors"
	"fmt"
	"github.com/fvbock/gorequests"
	"html"
	"net/http"
	"strings"
	"time"
)

const (
	SCOPEURL  = "http://api.microsofttranslator.com"
	GRANTTYPE = "client_credentials"

	AUTH_ENDPOINT            = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13/"
	TRANSLATE_ENDPOINT       = "http://api.microsofttranslator.com/v2/Http.svc/Translate"
	MULTI_TRANSLATE_ENDPOINT = "http://api.microsofttranslator.com/v2/Http.svc/TranslateArray"

	INPUT_TEXT = "text/plain"
	INPUT_HTML = "text/html"
)

var (
	clientId     string
	clientSecret string

	GlobalTranslatorToken *AccessToken
)

func SetCredentials(cId string, cSecret string) {
	clientId = cId
	clientSecret = cSecret
}

type AccessToken struct {
	Token     string `json:"access_token"`
	Type      string `json:"tokenType"`
	Scope     string `json:"scope"`
	ExpiresAt time.Time
}

/*
getToken retreives an AccessToken for the fiven ClientId and ClientSecret
*/
func getToken() (tok *AccessToken, err error) {
	data := map[string]string{
		"grant_type":    GRANTTYPE,
		"scope":         SCOPEURL,
		"client_id":     clientId,
		"client_secret": clientSecret,
	}

	r := gorequests.PostForm(AUTH_ENDPOINT, nil, data, -1)
	if r.Error != nil {
		err = errors.New(fmt.Sprintf("Error retrieving AccessToken:", r.Error))
		return
	}

	err = r.UnmarshalJson(&tok)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error unmarshaling AccessToken JSON:", r.Error))
	}
	// API docs state that the token will expire after 600s. give it a buffer.
	tok.ExpiresAt = time.Now().Add(570 * time.Second)
	return
}

type Translation struct {
	Text     string `xml:"TranslatedText"`
	Language string `xml:"To"`
}

type MultiTranslation struct {
	Translations []*Translation `xml:"TranslateArrayResponse"`
}

func (t *Translation) String() string {
	return fmt.Sprintf("%s: %s", t.Language, t.Text)
}

/*
Translate passes one or more texts to the Translator API. These texts must all
be be the same source language. Either plain text or html can be passed on: Set
the inputType to INPUT_TEXT or INPUT_HTML accordingly.

The AccessToken needed is stored in a global var and re-requested when in case
it expired.
*/
func Translate(from string, to string, input interface{}, inputType string) (translations []*Translation, err error) {
	// check the language codes
	from = strings.ToLower(from)
	if _, notsupported := supportedLanguages[from]; !notsupported {
		// if the source lang code is not set language detection will be
		// attempted on the API side
		if from != "" {
			err = errors.New(fmt.Sprintf("Source Language Code %s is currently not supported.", from))
			return
		}
	}

	to = strings.ToLower(to)
	if _, notsupported := supportedLanguages[to]; !notsupported {
		err = errors.New(fmt.Sprintf("Target Language Code %s is currently not supported.", to))
		return
	}

	// check whether we have a token or the current one is outdated. request
	// a new one in those cases.
	if GlobalTranslatorToken == nil || time.Now().After(GlobalTranslatorToken.ExpiresAt) {
		GlobalTranslatorToken, err = getToken()
		if err != nil {
			return
		}
	}

	// set the auth header
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Bearer %s", GlobalTranslatorToken.Token))

	// single and multi text translations work differently (different calls,
	// different request and response formats)
	var r *gorequests.Response
	var data map[string][]string
	var isSingleRequest bool
	switch input.(type) {
	case string:
		isSingleRequest = true
		data, err = gorequests.NewQueryData(
			map[string]string{
				"text":        input.(string),
				"from":        from,
				"to":          to,
				"contentType": inputType,
			})
		r = gorequests.Get(TRANSLATE_ENDPOINT, headers, data, -1)

	case []string:
		isSingleRequest = false
		texts := input.([]string)
		if inputType == INPUT_HTML {
			for i, t := range texts {
				texts[i] = html.EscapeString(t)
			}
		}
		treq := requestXML(from, to, texts, inputType)
		headers.Add("Content-Type", "text/xml")
		r = gorequests.Post(MULTI_TRANSLATE_ENDPOINT, headers, treq, nil, -1)

	}
	if r.Error != nil {
		err = r.Error
		return
	}

	if isSingleRequest {
		translation := &Translation{
			Language: to,
		}
		err = r.UnmarshalXML(&translation.Text)
		if err != nil {
			return
		}
		translations = append(translations, translation)
	} else {
		var mt MultiTranslation
		err = r.UnmarshalXML(&mt)
		if err != nil {
			return
		}
		for _, t := range mt.Translations {
			t.Language = to
		}
		translations = mt.Translations
	}

	return
}

/*
requestXML formats the MultiText translation request. just copied from the C#
example on the translator API site.
*/
func requestXML(from string, to string, input []string, inputType string) string {
	texts := func(in []string) string {
		var out string
		for _, t := range in {
			out += fmt.Sprintf("<string xmlns=\"http://schemas.microsoft.com/2003/10/Serialization/Arrays\">%s</string>", t)
		}
		return out
	}(input)

	base := `<TranslateArrayRequest>
    <AppId />
    <From>%s</From>
    <Options>
        <Category xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" />
        <ContentType xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2">%s</ContentType>
        <ReservedFlags xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" />
        <State xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" />
        <Uri xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" />
        <User xmlns="http://schemas.datacontract.org/2004/07/Microsoft.MT.Web.Service.V2" />
    </Options>
    <Texts>
        %s
    </Texts>
    <To>%s</To>
</TranslateArrayRequest>`
	return fmt.Sprintf(base, from, inputType, texts, to)
}
