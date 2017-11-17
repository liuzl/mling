package mling

type Tokenizer interface {
	Tokenize(text string) []string
}
