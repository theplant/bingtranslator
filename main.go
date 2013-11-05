package main

import (
	"flag"
	"fmt"
	btr "github.com/theplant/bingtranslator/translator"
	"os"
)

var (
	clientId     = flag.String("i", "", "ClientId as registered on Azure")
	clientSecret = flag.String("s", "", "The ClientSecret you got when registering your app")
)

func main() {
	flag.Parse()
	if *clientId == "" || *clientSecret == "" {
		fmt.Println("Please set the ClientID and Secret:\ngo run main.go -i <YOUR CLIENT ID> -s <YOUR CLIENT SECRET>")
		os.Exit(1)
	}
	fmt.Println("Translate")
	btr.SetCredentials(*clientId, *clientSecret)

	// single text
	var text string = "Let's see how you are translating this."
	translations, err := btr.Translate("en", "de", text, btr.INPUT_TEXT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(translations)

	// single html
	var html string = `<div class="stream-content-body markdown-style">
					<p><span class="j-show-object-content" data-object-id="foo" data-org-id="bar">@<span>Name</span></span>How hard would it be to extend the search system to spit out the high value terms in a piece of text?</p>
				</div>`

	translations, err = btr.Translate("en", "de", html, btr.INPUT_HTML)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(translations)

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

	// supported languages
	fmt.Println(btr.SupportedLanguages())
}
