package jackanalyzer

import (
	"encoding/xml"
	"fmt"
)

type Engine struct {
	*Tockenizer
	outputFile string
}

type class struct {
	Name       xml.Name `xml:"class"`
	Keyword    string   `xml:"keyword"`
	Identifier string   `xml:"identifier"`
	Symbol     string   `xml:"symbol"`
}

func NewEngine(inputFile, outputFile string) (*Engine, error) {
	tockenizer, err := NewTockenizer(inputFile)
	if err != nil {
		return nil, err
	}

	return &Engine{
		Tockenizer: tockenizer,
		outputFile: outputFile,
	}, nil
}

func (e *Engine) CompileClass() error {
	c := class{Keyword: "class"}

	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	token := e.Tockenizer.CurrentToken()
	if token.TokenType != IDENTIFIER {
		return fmt.Errorf("%s was not found", IDENTIFIER.String())
	}
	c.Identifier = token.Character

	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}

	token = e.Tockenizer.CurrentToken()
	if token.TokenType != SYMBOL {
		return fmt.Errorf("%s was not found", SYMBOL.String())
	}
	c.Symbol = token.Character

	return nil
}
