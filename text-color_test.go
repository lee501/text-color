package text_color

import (
	"testing"
)

func TestBlack(t *testing.T) {
	str := "black"
	t.Run("test black", func(t *testing.T) {
		println(Black(str))
	})
}

func TestRed(t *testing.T) {
	str := "red"
	t.Run("test red", func(t *testing.T) {
		println(Red(str))
	})
}