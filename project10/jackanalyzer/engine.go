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
		if !valueInTable(e.CurrentToken().Character, classVarDecTypes) {
			break
		}
		classVarDec, err := e.CompileClassVarDec()
		if err != nil {
			return err
		}
		c.Elements = append(c.Elements, classVarDec)
	}

	// Parse subroutineDec
	for {
		if !valueInTable(e.CurrentToken().Character, subroutineDecType) {
			break
		}
		subroutineDec, err := e.CompileSubroutineDec()
		if err != nil {
			return err
		}
		c.Elements = append(c.Elements, subroutineDec)
	}

	// Parse } closing class
	if err := e.checkTokenType(token, SYMBOL); err != nil {
		return err
	}
	token = e.Tockenizer.CurrentToken()
	if token.Character != "}" {
		return fmt.Errorf("missing { symbol")
	}
	c.Elements = append(c.Elements, Symbol{Value: token.Character})

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
	token = e.CurrentToken()
	c.Elements = append(c.Elements, Element{XMLName: xml.Name{Local: token.TokenType.String()}, Value: token.Character})

	// Parse var names
	if err := e.Tockenizer.Advance(); err != nil {
		return classVarDec{}, err
	}
	for {
		if err := e.checkTokenType(e.CurrentToken(), IDENTIFIER); err != nil {
			return classVarDec{}, err
		}
		c.Elements = append(c.Elements, Identifier{Value: e.CurrentToken().Character})

		if err := e.Tockenizer.Advance(); err != nil {
			return classVarDec{}, err
		}

		// check if multiple variables are being declared
		if token.Character == "," {
			c.Elements = append(c.Elements, Symbol{Value: e.CurrentToken().Character})
			if err := e.Tockenizer.Advance(); err != nil {
				return classVarDec{}, err
			}
			continue
		}

		// check if the classVarDec has ended
		if token.Character == ";" {
			c.Elements = append(c.Elements, Symbol{Value: e.CurrentToken().Character})
			break
		}

		return classVarDec{}, fmt.Errorf("class var declaration without proper end")
	}

	return c, nil
}

func (e *Engine) CompileSubroutineDec() (subroutineDec, error) {
	s := subroutineDec{}

	// Parse constructor | function | method
	token := e.Tockenizer.CurrentToken()
	if err := e.checkTokenType(token, KEYWORD); err != nil {
		return subroutineDec{}, err
	}
	if !valueInTable(token.Character, subroutineDecType) {
		return subroutineDec{}, fmt.Errorf("subroutineDec does not start with constructor, function or method")
	}
	s.Elements = []interface{}{
		Keyword{Value: token.Character},
	}

	// Parse void | type
	if err := e.Tockenizer.Advance(); err != nil {
		return subroutineDec{}, err
	}
	s.Elements = append(s.Elements, Element{XMLName: xml.Name{Local: e.CurrentToken().TokenType.String()}, Value: e.CurrentToken().Character})

	// Parse subroutine name
	if err := e.Tockenizer.Advance(); err != nil {
		return subroutineDec{}, err
	}
	s.Elements = append(s.Elements, Identifier{Value: e.CurrentToken().Character})

	// Parse (
	if err := e.Tockenizer.Advance(); err != nil {
		return subroutineDec{}, err
	}
	if err := e.checkTokenType(e.CurrentToken(), SYMBOL); err != nil {
		return subroutineDec{}, err
	}
	if e.CurrentToken().Character != "(" {
		return subroutineDec{}, fmt.Errorf("missing ( in subroutine Dec")
	}

	// Parse parameter list
	if err := e.Tockenizer.Advance(); err != nil {
		return subroutineDec{}, err
	}
	parameterList, err := e.CompileParameterList()
	if err != nil {
		return subroutineDec{}, err
	}
	s.Elements = append(s.Elements, parameterList)

	return s, nil
}

func (e *Engine) CompileSubroutineBody() (subroutineBody, error) {
	s := subroutineBody{}

	// Parse { opening
	token := e.CurrentToken()
	if err := e.checkTokenType(token, SYMBOL); err != nil {
		return subroutineBody{}, err
	}
	token = e.Tockenizer.CurrentToken()
	if token.Character != "{" {
		return subroutineBody{}, fmt.Errorf("missing { symbol")
	}
	s.Elements = []interface{}{Symbol{Value: token.Character}}

	// parse var dec
	varDec, err := e.CompileVarDec()
	if err != nil {
		return subroutineBody{}, err
	}
	s.Elements = append(s.Elements, varDec)

	// Parse } closing
	token = e.CurrentToken()
	if err := e.checkTokenType(token, SYMBOL); err != nil {
		return subroutineBody{}, err
	}
	if token.Character != "}" {
		return subroutineBody{}, fmt.Errorf("missing } symbol")
	}
	s.Elements = []interface{}{Symbol{Value: token.Character}}

	return s, nil
}

func (e *Engine) CompileVarDec() (varDec, error) {
	v := varDec{}

	// parse var
	token := e.CurrentToken()
	if err := e.checkTokenType(token, KEYWORD); err != nil {
		return varDec{}, err
	}
	v.Elements = []interface{}{Keyword{Value: token.Character}}

	// parse type
	if err := e.Tockenizer.Advance(); err != nil {
		return varDec{}, err
	}
	token = e.CurrentToken()
	v.Elements = append(v.Elements, Element{XMLName: xml.Name{Local: token.TokenType.String()}, Value: token.Character})

	for {
		// parse var name
		if err := e.Tockenizer.Advance(); err != nil {
			return varDec{}, err
		}
		if err := e.checkTokenType(e.CurrentToken(), IDENTIFIER); err != nil {
			return varDec{}, err
		}
		v.Elements = append(v.Elements, Identifier{Value: e.CurrentToken().Character})

		// parse possible ","
		if err := e.Tockenizer.Advance(); err != nil {
			return varDec{}, err
		}
		if e.CurrentToken().Character == "," {
			v.Elements = append(v.Elements, Symbol{Value: e.CurrentToken().Character})
		} else {
			break
		}
	}

	// parse ;
	if err := e.checkTokenType(e.CurrentToken(), SYMBOL); err != nil {
		return varDec{}, err
	}
	if e.CurrentToken().Character != ";" {
		return varDec{}, fmt.Errorf("missing ;")
	}
	v.Elements = append(v.Elements, Symbol{Value: e.CurrentToken().Character})

	return v, nil
}

func (e *Engine) CompileParameterList() (parameterList, error) {
	p := parameterList{}
	for {
		if e.CurrentToken().TokenType != KEYWORD && e.CurrentToken().TokenType != SYMBOL {
			break
		}

		// parse keyword
		if err := e.checkTokenType(e.CurrentToken(), KEYWORD); err != nil {
			return parameterList{}, err
		}
		if !valueInTable(e.CurrentToken().Character, varTypes) {
			return parameterList{}, fmt.Errorf("invalid types in parameter list")
		}
		p.Elements = append(p.Elements, Keyword{Value: e.CurrentToken().Character})

		// parse identifier
		if err := e.Tockenizer.Advance(); err != nil {
			return parameterList{}, err
		}
		if err := e.checkTokenType(e.CurrentToken(), IDENTIFIER); err != nil {
			return parameterList{}, err
		}
		p.Elements = append(p.Elements, Identifier{Value: e.CurrentToken().Character})

		// look for ,
		if err := e.Tockenizer.Advance(); err != nil {
			return parameterList{}, err
		}
		if e.CurrentToken().Character == "," {
			p.Elements = append(p.Elements, Symbol{Value: e.CurrentToken().Character})
			if err := e.Tockenizer.Advance(); err != nil {
				return parameterList{}, err
			}
		}

	}
	return p, nil
}

func (e *Engine) checkTokenType(token Token, tokenType TokenType) error {
	if token.TokenType != tokenType {
		return fmt.Errorf("%s was not found", tokenType.String())
	}
	return nil
}
