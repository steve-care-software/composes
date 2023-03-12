package composes

import "github.com/steve-care-software/libs/cryptography/hash"

// Builder represents a compose builder
type Builder interface {
	Create() Builder
	WithTokens(tokens []Token) Builder
	WithVariables(variables uint) Builder
	Now() (Compose, error)
}

// Compose represents a compose
type Compose interface {
	Tokens() []Token
	Variables() uint
}

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithList(list []Element) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	List() []Element
}

// ElementBuilder represents an element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithAmount(amount uint) ElementBuilder
	WithToken(token Token) ElementBuilder
	WithValue(value []byte) ElementBuilder
	WithVariable(variable uint) ElementBuilder
	Now() (Element, error)
}

// Element represents a compose element
type Element interface {
	Hash() hash.Hash
	Amount() uint
	Content() Content
}

// Content represents an element's content
type Content interface {
	Hash() hash.Hash
	IsToken() bool
	Token() Token
	IsValue() bool
	Value() []byte
	IsVariable() bool
	Variable() *uint
}
