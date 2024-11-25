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
	c := class{}

	// Parse keyword class
	token := e.Tockenizer.CurrentToken()
	if err := e.checkTokenType(token, KEYWORD); err != nil {
		return err
	}
	if token.Character != "class" {
		return fmt.Errorf("missing class declaration")
	}
	c.Elements = []interface{}{Keyword{Value: token.Character}}

	// Parse class identifier
	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	token = e.Tockenizer.CurrentToken()
	if err := e.checkTokenType(token, IDENTIFIER); err != nil {
		return err
	}
	c.Elements = append(c.Elements, Identifier{Value: token.Character})

	// Parse { opening class
	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	if err := e.checkTokenType(token, SYMBOL); err != nil {
		return err
	}
	token = e.Tockenizer.CurrentToken()
	if token.Character != "{" {
		return fmt.Errorf("missing { symbol")
	}
	c.Elements = append(c.Elements, Symbol{Value: token.Character})

	// Parse classVarDec
	if err := e.Tockenizer.Advance(); err != nil {
		return err
	}
	for {
		if !valueInTable(token.Character, classVarDecTypes) {
			break
		}
		classVarDec, err := e.CompileClassVarDec()
		if err != nil {
			return err
		}
		c.Elements = append(c.Elements, classVarDec)
	}

	return nil
}

func (e *Engine) CompileClassVarDec() (classVarDec, error) {
	c := classVarDec{}

	// Parse static | field
	token := e.Tockenizer.CurrentToken()
	if err := e.checkTokenType(token, KEYWORD); err != nil {
		return classVarDec{}, err
	}
	if !valueInTable(token.Character, classVarDecTypes) {
		return classVarDec{}, fmt.Errorf("class variables are not static or field")
	}
	c.Elements = []interface{}{
		Keyword{Value: token.Character},
	}

	// Parse type
	if err := e.Tockenizer.Advance(); err != nil {
		return classVarDec{}, err
	}
	c.Elements = append(c.Elements, Element{XMLName: xml.Name{Local: token.TokenType.String()}, Value: token.Character})

	// Parse var names
	if err := e.Tockenizer.Advance(); err != nil {
		return classVarDec{}, err
	}
	for {
		if err := e.checkTokenType(token, IDENTIFIER); err != nil {
			return classVarDec{}, err
		}
		c.Elements = append(c.Elements, Identifier{Value: token.Character})

		if err := e.Tockenizer.Advance(); err != nil {
			return classVarDec{}, err
		}

		// check if multiple variables are being declared
		if token.Character == "," {
			c.Elements = append(c.Elements, Symbol{Value: token.Character})
			if err := e.Tockenizer.Advance(); err != nil {
				return classVarDec{}, err
			}
			continue
		}

		// check if the classVarDec has ended
		if token.Character == ";" {
			c.Elements = append(c.Elements, Symbol{Value: token.Character})
			break
		}

		return classVarDec{}, fmt.Errorf("class var declaration without proper end")
	}

	return c, nil
}

func (e *Engine) checkTokenType(token Token, tokenType TokenType) error {
	if token.TokenType != tokenType {
		return fmt.Errorf("%s was not found", tokenType.String())
	}
	return nil
}
