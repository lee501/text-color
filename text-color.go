package text_color

import (
	"fmt"
	"strings"
)

var _reset string = "\x1b[0m"

// Foreground color functions

func Default(text string) string {
	return _reset + text
}

func Red(text string) string {
	return "\x1b[31m" + text + _reset
}

func Blue(text string) string {
	return "\x1b[34m" + text + _reset
}

func Yellow(text string) string {
	return "\x1b[33m" + text + _reset
}

func Green(text string) string {
	return "\x1b[32m" + text + _reset
}

func Purple(text string) string {
	return "\x1b[35m" + text + _reset
}

func Cyan(text string) string {
	return "\x1b[36m" + text + _reset
}

func White(text string) string {
	return "\x1b[37m" + text + _reset
}

func Black(text string) string {
	return "\x1b[30m" + text + _reset
}

// Background color functions

func BgRed(text string) string {
	return "\x1b[41m" + text + _reset
}

func BgBlue(text string) string {
	return "\x1b[44m" + text + _reset
}

func BgYellow(text string) string {
	return "\x1b[43m" + text + _reset
}

func BgGreen(text string) string {
	return "\x1b[42m" + text + _reset
}

func BgPurple(text string) string {
	return "\x1b[45m" + text + _reset
}

func BgCyan(text string) string {
	return "\x1b[46m" + text + _reset
}

func BgWhite(text string) string {
	return "\x1b[47m" + text + _reset
}

func BgBlack(text string) string {
	return "\x1b[40m" + text + _reset
}

// Text style functions

func Bold(text string) string {
	return "\x1b[1m" + text + _reset
}

func Italic(text string) string {
	return "\x1b[3m" + text + _reset
}

func Underline(text string) string {
	return "\x1b[4m" + text + _reset
}

func Strikethrough(text string) string {
	return "\x1b[9m" + text + _reset
}

// Style provides a chainable builder to combine multiple text attributes.
type Style struct {
	codes []string
}

// NewStyle creates a new Style builder.
func NewStyle() *Style {
	return &Style{}
}

func (s *Style) addCode(code string) *Style {
	s.codes = append(s.codes, code)
	return s
}

// Foreground color methods

func (s *Style) Red() *Style         { return s.addCode("31") }
func (s *Style) Blue() *Style        { return s.addCode("34") }
func (s *Style) Yellow() *Style      { return s.addCode("33") }
func (s *Style) Green() *Style       { return s.addCode("32") }
func (s *Style) Purple() *Style      { return s.addCode("35") }
func (s *Style) Cyan() *Style        { return s.addCode("36") }
func (s *Style) White() *Style       { return s.addCode("37") }
func (s *Style) Black() *Style       { return s.addCode("30") }

// Background color methods

func (s *Style) BgRed() *Style       { return s.addCode("41") }
func (s *Style) BgBlue() *Style      { return s.addCode("44") }
func (s *Style) BgYellow() *Style    { return s.addCode("43") }
func (s *Style) BgGreen() *Style     { return s.addCode("42") }
func (s *Style) BgPurple() *Style    { return s.addCode("45") }
func (s *Style) BgCyan() *Style      { return s.addCode("46") }
func (s *Style) BgWhite() *Style     { return s.addCode("47") }
func (s *Style) BgBlack() *Style     { return s.addCode("40") }

// Text style methods

func (s *Style) Bold() *Style          { return s.addCode("1") }
func (s *Style) Italic() *Style        { return s.addCode("3") }
func (s *Style) Underline() *Style     { return s.addCode("4") }
func (s *Style) Strikethrough() *Style { return s.addCode("9") }

// Apply applies all accumulated styles to the given text.
func (s *Style) Apply(text string) string {
	if len(s.codes) == 0 {
		return text
	}
	return fmt.Sprintf("\x1b[%sm%s%s", strings.Join(s.codes, ";"), text, _reset)
}

// Sprintf formats the string and applies all accumulated styles.
func (s *Style) Sprintf(format string, a ...interface{}) string {
	return s.Apply(fmt.Sprintf(format, a...))
}
