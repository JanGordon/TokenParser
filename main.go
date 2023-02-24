package main

import "fmt"

type Token struct {
	OpenLabel       string
	CloseLabel      string
	Children        []*Token
	OriginalContent string
	StartIndex      int
}

func (t *Token) GetTokenlessContent() string {
	return ""
}

type Parser struct {
	Tokens []*Token
}

func (p *Parser) testForLabel(label string, index int, text string) bool {
	if index+len(label) < len(text) {
		fragment := text[index : index+len(label)]
		if fragment == label {
			return true
		}
	}
	return false
}

func (p *Parser) Parse(text string) []*Token {
	var outTokens []*Token
	var openTokens []*Token
	for index, _ := range text {
		// find if tokens label matches the current index
		for _, token := range p.Tokens {
			// test both ope nand close labels
			if p.testForLabel(token.OpenLabel, index, text) {
				fmt.Println("adding opne")
				newToken := *token
				newToken.StartIndex = index

				openTokens = append(openTokens, &newToken)

				continue
			}
			if p.testForLabel(token.CloseLabel, index, text) {
				lastOpenToken := openTokens[len(openTokens)-1]
				lastOpenToken.OriginalContent = text[lastOpenToken.StartIndex:index]
				fmt.Println(lastOpenToken.StartIndex, index)
				if len(openTokens)-2 < 0 {
					outTokens = append(outTokens, lastOpenToken)
				} else {
					openTokens[len(openTokens)-2].Children = append(openTokens[len(openTokens)-2].Children, lastOpenToken)
				}
				openTokens = openTokens[:len(openTokens)-1] // pop the last element
			}
		}
	}
	return outTokens
}

func main() {
	p := &Parser{[]*Token{
		{
			"!!",
			">!",
			nil,
			"",
			0,
		},
	}}
	for _, r := range p.Parse("hello !! actually bye planet !! a child >! >!  world") {
		fmt.Println(r, r.Children[0])
	}
}
