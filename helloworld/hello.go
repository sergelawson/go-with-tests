package helloworld

var HelloMap = map[string]string{
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
	"German":  "Hallo, ",
	"English": "Hello, ",
}

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	helloPrefix := HelloMap["English"]

	if language != "" {
		helloPrefix = HelloMap[language]
	}
	return helloPrefix + name
}
