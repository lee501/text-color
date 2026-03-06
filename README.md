# text-color

A lightweight Go library to display colored and styled text in the terminal using ANSI escape codes.

## Install

```bash
go get github.com/lee501/text-color
```

## Features

- **Foreground colors**: Red, Green, Blue, Yellow, Purple, Cyan, White, Black
- **Background colors**: BgRed, BgGreen, BgBlue, BgYellow, BgPurple, BgCyan, BgWhite, BgBlack
- **Text styles**: Bold, Italic, Underline, Strikethrough
- **Style builder**: Combine multiple styles with a chainable API

## Usage

### Basic colors

```go
package main

import (
	tc "github.com/lee501/text_color"
)

func main() {
	println(tc.Red("This is red"))
	println(tc.Green("This is green"))
	println(tc.Blue("This is blue"))
	println(tc.Yellow("This is yellow"))
}
```

### Background colors

```go
println(tc.BgRed("Red background"))
println(tc.BgBlue("Blue background"))
```

### Text styles

```go
println(tc.Bold("Bold text"))
println(tc.Underline("Underlined text"))
println(tc.Italic("Italic text"))
println(tc.Strikethrough("Strikethrough text"))
```

### Combining styles

Use `NewStyle()` to combine multiple colors and styles:

```go
// Bold red text on a white background
style := tc.NewStyle().Bold().Red().BgWhite()
println(style.Apply("Bold red on white"))

// Underlined cyan text
println(tc.NewStyle().Underline().Cyan().Apply("Underlined cyan"))

// Format with Sprintf
println(tc.NewStyle().Green().Bold().Sprintf("Count: %d", 42))
```

## API Reference

### Foreground Colors

| Function | Description |
|----------|-------------|
| `Red(text)` | Red text |
| `Green(text)` | Green text |
| `Blue(text)` | Blue text |
| `Yellow(text)` | Yellow text |
| `Purple(text)` | Purple text |
| `Cyan(text)` | Cyan text |
| `White(text)` | White text |
| `Black(text)` | Black text |
| `Default(text)` | Reset to default |

### Background Colors

| Function | Description |
|----------|-------------|
| `BgRed(text)` | Red background |
| `BgGreen(text)` | Green background |
| `BgBlue(text)` | Blue background |
| `BgYellow(text)` | Yellow background |
| `BgPurple(text)` | Purple background |
| `BgCyan(text)` | Cyan background |
| `BgWhite(text)` | White background |
| `BgBlack(text)` | Black background |

### Text Styles

| Function | Description |
|----------|-------------|
| `Bold(text)` | Bold text |
| `Italic(text)` | Italic text |
| `Underline(text)` | Underlined text |
| `Strikethrough(text)` | Strikethrough text |

### Style Builder

```go
style := NewStyle().Bold().Red().BgWhite()
style.Apply("text")            // Apply styles to text
style.Sprintf("num: %d", 42)  // Format and apply styles
```
