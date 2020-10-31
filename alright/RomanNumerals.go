package main

import "strings"

func ConvertToArabic(roman string) (total int) {
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allRomanNumerals.ValueOf(symbols...)
	}
	return
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return true
		}
	}
	return false
}

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
			i++
		} else {
			symbols = append(symbols, []byte{byte(symbol)})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

//type RomanNumeral struct {
//	Value int
//	Symbol string
//}
//
//type RomanNumerals []RomanNumeral
//
//
//var allRomanNumerals = []RomanNumeral {
//	{1000,"M"},
//	{900,"CM"},
//	{500,"D"},
//	{400,"CD"},
//	{100,"C"},
//	{90,"XC"},
//	{50,"L"},
//	{40,"XL"},
//	{10,"X"},
//	{9, "IX"},
//	{5,"V"},
//	{4,"IV"},
//	{1,"I"},
//}
//
//
//
//func ConvertToRoman(num int) string {
//	var result strings.Builder
//
//	for _, numeral := range allRomanNumerals {
//		for num >= numeral.Value {
//			result.WriteString(numeral.Symbol)
//			num -= numeral.Value
//		}
//	}
//
//	return result.String()
//}
//
//
//
//
//
//func (r RomanNumerals) ValueOf(symbol string) int {
//	for _, s := range r {
//		if s.Symbol == symbol {
//			return s.Value
//		}
//
//	}
//
//	return 0
//}
//
//
//func ConvertingToArabic(num string) int {
//	total := 0
//	for i := 0; i<len(num); i++ {
//		symbol := num[i]
//
//		if couldBeSubtractive(i,symbol,num) {
//			nextSymbol := num[i+1]
//
//			potentialNumber := string([]byte{symbol,nextSymbol})
//
//			if value := allRomanNumerals.ValueOf(potentialNumber); value !=0 {
//				total += value
//				i++
//			} else {
//				total++
//			}
//		} else {
//			total+= allRomanNumerals.ValueOf(string([]byte{symbol}))
//		}
//	}
//	return total
//}
//
//func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
//	return index+1 < len(roman) && currentSymbol=='I'
//}
//
//
//func main() {
//	print(ConvertToRoman(40))
//}

//for num > 0 {
//switch {
//case num > 9:
//result.WriteString("X")
//num -= 10
//case num > 8:
//result.WriteString("IX")
//num -= 9
//case num > 4:
//result.WriteString("V")
//num -= 5
//case num > 3:
//result.WriteString("IV")
//num -= 4
//default:
//result.WriteString("I")
//num --
//
//}
//}
