package translate

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const path = "/v2/translate"

type Response struct {
	Translations []Text `json:"translations"`
}

type Text struct {
	DetectedSourceLanguage string `json:"detected_source_language"`
	Text                   string `json:"text"`
}

type Request struct {
	AuthKey string   `url:"auth_key"`
	Text    []string `url:"text"`
	// TargetLang is require
	// "DE" - German
	// "EN" - English
	// "FR" - French
	// "IT" - Italian
	// "JA" - Japanese
	// "ES" - Spanish
	// "NL" - Dutch
	// "PL" - Polish
	// "PT-PT" - Portuguese (all Portuguese varieties excluding Brazilian Portuguese)
	// "PT-BR" - Portuguese (Brazilian)
	// "PT" - Portuguese (unspecified variant for backward compatibility; please select PT-PT or PT-BR instead)
	// "RU" - Russian
	// "ZH" - Chinese
	TargetLang string `url:"target_lang"`
	// SourceLang is optional
	// "DE" - German
	// "EN" - English
	// "FR" - French
	// "IT" - Italian
	// "JA" - Japanese
	// "ES" - Spanish
	// "NL" - Dutch
	// "PL" - Polish
	// "PT" - Portuguese (all Portuguese varieties mixed)
	// "RU" - Russian
	// "ZH" - Chinese
	SourceLang string `url:"source_lang,omitempty"`
	// SplitSentences is optional
	// "0" - no splitting at all, whole input is treated as one sentence
	// "1" (default) - splits on interpunction and on newlines
	// "nonewlines" - splits on interpunction only, ignoring newlines
	SplitSentences string `url:"split_sentences,omitempty"`
	// PreserveFormatting is optional
	// "0" (default)
	// "1"
	PreserveFormatting string `url:"preserve_formatting,omitempty"`
	// Formality is optional
	// "default" (default)
	// "more" - for a more formal language
	// "less" - for a more informal language
	Formality string `url:"formality,omitempty"`

	// AND MORE
	// FOR XML
	// TagHandling is optional for XML
	TagHandling string `url:"tag_handling,omitempty"`
	// OutlineDetection is optional for XML
	OutlineDetection string `url:"outline_detection,omitempty"`
	// SplittingTags is optional for XML
	SplittingTags string `url:"splitting_tags,omitempty"`
	// NonSplittingTags is optional for XML
	NonSplittingTags string `url:"non_splitting_tags,omitempty"`
	// IgnoreTags is optional for XML
	IgnoreTags string `url:"ignore_tags,omitempty"`
}

func (p *Request) SetAuth(key string) {
	p.AuthKey = key
}

func (p *Request) Path() string {
	return path
}

func (p *Request) Method() string {
	return http.MethodPost
}

func (p *Request) Query() string {
	v, _ := query.Values(p)
	return v.Encode()
}

func (p *Request) Param() *strings.Reader {
	v, _ := query.Values(p)
	return strings.NewReader(v.Encode())
}

func (p *Response) Unescape() {
	for i := range p.Translations {
		t, err := url.QueryUnescape(p.Translations[i].Text)
		if err != nil {
			continue
		}
		p.Translations[i].Text = t
	}
}
