package main

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	// going to used table based testing

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertCorrect(t, got, test.Roman)
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases[:4] {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertingToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func assertCorrect(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}

var cases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	{"1 gets converted to I", 1, "I"},
	{"2 gets converted to II", 2, "II"},
	{"3 gets converted to III", 3, "III"},
	{"4 gets converted to IV", 4, "IV"},
	{"5 gets converted to V", 5, "V"},
	{"6 gets converted to VI", 6, "VI"},
	{"7 gets converted to VII", 7, "VII"},
	{"8 gets converted to VIII", 8, "VIII"},
	{"9 gets converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"11 gets converted to XI", 11, "XI"},
	{"12 gets converted to XII", 12, "XII"},
	{"13 gets converted to XIII", 13, "XIII"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"15 gets converted to XV", 15, "XV"},
	{"16 gets converted to XVI", 16, "XVI"},
	{"17 gets converted to XVII", 17, "XVII"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"19 gets converted to XIX", 19, "XIX"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"2014 gets converted to MMXIV", 2014, "MMXIV"},
	{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
}

//t.Run("1 gets converted to I", func(t *testing.T) {
//	got := ConvertToRoman(1)
//	want := "I"
//
//	assertCorrect(t,got,want)
//})
//
//t.Run("2 gets converted to II", func(t *testing.T) {
//	got := ConvertToRoman(2)
//	want := "II"
//
//	assertCorrect(t,got,want)
//})
