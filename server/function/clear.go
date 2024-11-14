package main

import "strings"

// Fonction pour enlever les accents
func enleverAccents(mot string) string {
	accents := []string{"é", "è", "ê", "ë", "ï", "î", "ô", "ö", "ù", "ü", "û", "à", "â", "ä", "ç", "ÿ", "œ", "æ", "ᵫ", "ꭣ", "ꭡ"}
	nonAccents := []string{"e", "e", "e", "e", "i", "i", "o", "o", "u", "u", "u", "a", "a", "a", "c", "y", "oe", "ae", "ue", "uo", "ie"}
	for i, accent := range accents {
		mot = strings.Replace(mot, accent, nonAccents[i], -1)
	}
	return mot
}

func enleverMajuscules(mot string) string {
	majuscules := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "Á", "À", "Â", "Ã", "Å", "Ä", "Ç", "È", "É", "Ê", "Ë", "Î", "Ì", "Ï", "Ñ", "Ô", "Õ", "Û", "Ù", "Ü", "Ÿ", "Æ", "Ꜵ", "Œ"}
	nonMajuscules := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "a", "a", "a", "a", "a", "a", "c", "e", "e", "e", "e", "i", "i", "i", "n", "o", "o", "u", "u", "u", "y", "ae", "ao", "oe"}
	for i, accent := range majuscules {
		mot = strings.Replace(mot, accent, nonMajuscules[i], -1)
	}
	return mot
}
