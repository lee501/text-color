package text_color

import (
	"testing"
)

var reset = "\x1b[0m"

// Tests for foreground color functions

func TestBlack(t *testing.T) {
	got := Black("black")
	want := "\x1b[30mblack" + reset
	if got != want {
		t.Errorf("Black() = %q, want %q", got, want)
	}
}

func TestRed(t *testing.T) {
	got := Red("red")
	want := "\x1b[31mred" + reset
	if got != want {
		t.Errorf("Red() = %q, want %q", got, want)
	}
}

func TestGreen(t *testing.T) {
	got := Green("green")
	want := "\x1b[32mgreen" + reset
	if got != want {
		t.Errorf("Green() = %q, want %q", got, want)
	}
}

func TestYellow(t *testing.T) {
	got := Yellow("yellow")
	want := "\x1b[33myellow" + reset
	if got != want {
		t.Errorf("Yellow() = %q, want %q", got, want)
	}
}

func TestBlue(t *testing.T) {
	got := Blue("blue")
	want := "\x1b[34mblue" + reset
	if got != want {
		t.Errorf("Blue() = %q, want %q", got, want)
	}
}

func TestPurple(t *testing.T) {
	got := Purple("purple")
	want := "\x1b[35mpurple" + reset
	if got != want {
		t.Errorf("Purple() = %q, want %q", got, want)
	}
}

func TestCyan(t *testing.T) {
	got := Cyan("cyan")
	want := "\x1b[36mcyan" + reset
	if got != want {
		t.Errorf("Cyan() = %q, want %q", got, want)
	}
}

func TestWhite(t *testing.T) {
	got := White("white")
	want := "\x1b[37mwhite" + reset
	if got != want {
		t.Errorf("White() = %q, want %q", got, want)
	}
}

func TestDefault(t *testing.T) {
	got := Default("text")
	want := reset + "text"
	if got != want {
		t.Errorf("Default() = %q, want %q", got, want)
	}
}

// Tests for background color functions

func TestBgRed(t *testing.T) {
	got := BgRed("text")
	want := "\x1b[41mtext" + reset
	if got != want {
		t.Errorf("BgRed() = %q, want %q", got, want)
	}
}

func TestBgBlue(t *testing.T) {
	got := BgBlue("text")
	want := "\x1b[44mtext" + reset
	if got != want {
		t.Errorf("BgBlue() = %q, want %q", got, want)
	}
}

func TestBgYellow(t *testing.T) {
	got := BgYellow("text")
	want := "\x1b[43mtext" + reset
	if got != want {
		t.Errorf("BgYellow() = %q, want %q", got, want)
	}
}

func TestBgGreen(t *testing.T) {
	got := BgGreen("text")
	want := "\x1b[42mtext" + reset
	if got != want {
		t.Errorf("BgGreen() = %q, want %q", got, want)
	}
}

func TestBgPurple(t *testing.T) {
	got := BgPurple("text")
	want := "\x1b[45mtext" + reset
	if got != want {
		t.Errorf("BgPurple() = %q, want %q", got, want)
	}
}

func TestBgCyan(t *testing.T) {
	got := BgCyan("text")
	want := "\x1b[46mtext" + reset
	if got != want {
		t.Errorf("BgCyan() = %q, want %q", got, want)
	}
}

func TestBgWhite(t *testing.T) {
	got := BgWhite("text")
	want := "\x1b[47mtext" + reset
	if got != want {
		t.Errorf("BgWhite() = %q, want %q", got, want)
	}
}

func TestBgBlack(t *testing.T) {
	got := BgBlack("text")
	want := "\x1b[40mtext" + reset
	if got != want {
		t.Errorf("BgBlack() = %q, want %q", got, want)
	}
}

// Tests for text style functions

func TestBold(t *testing.T) {
	got := Bold("text")
	want := "\x1b[1mtext" + reset
	if got != want {
		t.Errorf("Bold() = %q, want %q", got, want)
	}
}

func TestItalic(t *testing.T) {
	got := Italic("text")
	want := "\x1b[3mtext" + reset
	if got != want {
		t.Errorf("Italic() = %q, want %q", got, want)
	}
}

func TestUnderline(t *testing.T) {
	got := Underline("text")
	want := "\x1b[4mtext" + reset
	if got != want {
		t.Errorf("Underline() = %q, want %q", got, want)
	}
}

func TestStrikethrough(t *testing.T) {
	got := Strikethrough("text")
	want := "\x1b[9mtext" + reset
	if got != want {
		t.Errorf("Strikethrough() = %q, want %q", got, want)
	}
}

// Tests for Style builder

func TestStyleSingleColor(t *testing.T) {
	got := NewStyle().Red().Apply("hello")
	want := "\x1b[31mhello" + reset
	if got != want {
		t.Errorf("Style.Red().Apply() = %q, want %q", got, want)
	}
}

func TestStyleCombined(t *testing.T) {
	got := NewStyle().Bold().Red().BgWhite().Apply("hello")
	want := "\x1b[1;31;47mhello" + reset
	if got != want {
		t.Errorf("Style combined = %q, want %q", got, want)
	}
}

func TestStyleEmpty(t *testing.T) {
	got := NewStyle().Apply("hello")
	want := "hello"
	if got != want {
		t.Errorf("Style empty = %q, want %q", got, want)
	}
}

func TestStyleSprintf(t *testing.T) {
	got := NewStyle().Green().Bold().Sprintf("count: %d", 42)
	want := "\x1b[32;1mcount: 42" + reset
	if got != want {
		t.Errorf("Style.Sprintf() = %q, want %q", got, want)
	}
}

func TestStyleUnderlineCyan(t *testing.T) {
	got := NewStyle().Underline().Cyan().Apply("link")
	want := "\x1b[4;36mlink" + reset
	if got != want {
		t.Errorf("Style underline+cyan = %q, want %q", got, want)
	}
}

// Test empty string

func TestEmptyString(t *testing.T) {
	got := Red("")
	want := "\x1b[31m" + reset
	if got != want {
		t.Errorf("Red(\"\") = %q, want %q", got, want)
	}
}