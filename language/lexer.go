package language

// TODO
const (
	SOF          = "<SOF>"
	EOF          = "<EOF>"
	Bang         = "!"
	Dollar       = "$"
	ParenLeft    = "("
	ParenRight   = ")"
	Spread       = "..."
	Colon        = ":"
	Equals       = "="
	At           = "@"
	BracketLeft  = "["
	BracketRight = "]"
	BraceLeft    = "{"
	Pipe         = "|"
	BraceRight   = "}"
	Name         = "Name"
	Int          = "Int"
	Float        = "Float"
	String       = "String"
	Comment      = "Comment"
)

// Token ...
type Token struct {
	Kind   string
	Start  int
	End    int
	Line   int
	Column int
	Value  string
}

// Lexer ...
type Lexer struct {
	Source    *Source
	Options   *map[string]string
	Token     *Token
	PrevToken *Token
	NextToken *Token
	Line      int
	LineStart int
}

// NewLexer ...
func NewLexer(source *Source, options *map[string]string) *Lexer {
	lexer := &Lexer{Source: source, Options: options}

	SOFToken := &Token{Kind: SOF}

	lexer.Token = SOFToken
	lexer.PrevToken = SOFToken

	return lexer
}

// Advance ...
func (l *Lexer) Advance() *Token {
	l.LastToken = l.Token

	token := l.readToken()
	for token.Kind == Comment {
		token := l.readToken()
	}

	return l.Token
}

func (l *Lexer) readToken() *Token {
	bodyLength := len(l.Source.Body)

	position := l.positionAfterWhitespace()
	line := l.Line
	col := position - l.LineStart

	if position >= bodyLength {
		return &Token{Kind: EOF, Start: bodyLength, End: bodyLength, Line: line, Column: col}
	}

	return nil
}

func (l *Lexer) positionAfterWhitespace() int {
}

func (l *Lexer) readComment() *Token {
	return nil
}

func (l *Lexer) readNumber() *Token {
	return nil
}

func (l *Lexer) readDigits() *Token {
	return nil
}

func (l *Lexer) readString() *Token {
	return nil
}

func (l *Lexer) readName() *Token {
	return nil
}
