package deepl

type Lang int

const (
	_ Lang = iota
	LangGerman
	LangEnglish
	LangFrench
	LangItaian
	LangJapanese
	LangSpanish
	LangDutch
	LangPolish
	LangPortuguese
	LangBrazPortuguese
	LangPortugueseAndBraz
	LangRussian
	LangChinese
)

// String return string for params
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
func (l Lang) String() string {
	switch l {
	case LangGerman:
		return "DE"
	case LangEnglish:
		return "EN"
	case LangItaian:
		return "IT"
	case LangJapanese:
		return "JA"
	case LangFrench:
		return "FR"
	case LangSpanish:
		return "ES"
	case LangDutch:
		return "NL"
	case LangPolish:
		return "PL"
	case LangPortuguese:
		return "PT-PT"
	case LangBrazPortuguese:
		return "PT-BR"
	case LangPortugueseAndBraz:
		return "PT"
	case LangRussian:
		return "RU"
	case LangChinese:
		return "ZH"
	}

	return "Undefined"
}
