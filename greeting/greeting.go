package greeting

// Greeter is the interface that wraps the basic greet method.
type Greeter interface {
	// Greet in a language.
	Greet() string
}

// New returns a greeter in the requested language, if the language is unknown
// it returns english greeter as default.
func New(lang string) Greeter {
	switch lang {
	case "english":
		return englishMan()
	case "chinese":
		return chinaMan()
	case "telugu":
		return babu()
	case "tamil":
		return ponnu()
	case "kannada":
		return maga()
	case "sanskrit":
		return purusha()
	case "hindi":
		return gori()
	default:
		return englishMan()
	}
}
