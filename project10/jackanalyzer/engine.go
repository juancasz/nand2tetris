package jackanalyzer

import (
	"encoding/xml"
	"fmt"
)

type Engine struct {
	*Tockenizer
	outputFile string
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
	if token.Character != "{" {
		return fmt.Errorf("missing { symbol")
	}
	c.Symbol = token.Character

	return nil
}

func (e *Engine) CompileClassVarDec() error {
	c := classVarDec{}

	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	token := e.Tockenizer.CurrentToken()
	if err := e.checkTokenType(token, KEYWORD); err != nil {
		return err
	}
	if !valueInTable(token.Character, classVarDecTypes) {
		return fmt.Errorf("class variables are not static or field")
	}
	c.Elements = []interface{}{
		Keyword{Value: token.Character},
	}

	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	c.Elements = append(c.Elements, Element{XMLName: xml.Name{Local: token.TokenType.String()}, Value: token.Character})

	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}

	for {
		if err := e.checkTokenType(token, IDENTIFIER); err != nil {
			return err
		}
		c.Elements = append(c.Elements, Identifier{Value: token.Character})

		if err := e.Tockenizer.Advance(); err != nil {
			return err
		}

		if token.Character == "," {
			c.Elements = append(c.Elements, Symbol{Value: token.Character})
			if err := e.Tockenizer.Advance(); err != nil {
				return err
			}
			continue
		}

		if token.Character == ";" {
			c.Elements = append(c.Elements, Symbol{Value: token.Character})
			break
		}

		return fmt.Errorf("class var declaration without proper end")
	}

	return nil
}

func (e *Engine) checkTokenType(token Token, tokenType TokenType) error {
	if token.TokenType != tokenType {
		return fmt.Errorf("%s was not found", tokenType.String())
	}
	return nil
}
