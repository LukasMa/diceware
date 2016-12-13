package diceware

import (
	"strings"
	"testing"
)

func TestNewPassphrase(t *testing.T) {
	// Should create passphrase with default settings.
	phrase, err := NewPassphrase()
	ok(t, err)
	equals(t, DefaultExtra, phrase.extra)
	equals(t, DefaultWords, phrase.wordCount)
	equals(t, DefaultWords, len(phrase.words))
}

func TestNewPassphraseExtra(t *testing.T) {
	testCases := []bool{
		DefaultExtra,
		true,
		false,
	}

	for _, test := range testCases {
		phrase, err := NewPassphrase(
			Extra(test),
		)
		ok(t, err)
		equals(t, test, phrase.extra)
	}
}

func TestNewPassphraseWordstestalid(t *testing.T) {
	testCases := []int{
		MinWords,
		MinWords + 1,
		DefaultWords - 1,
		DefaultWords,
		DefaultWords + 1,
		DefaultWords + 10,
	}

	for _, wordCount := range testCases {
		phrase, err := NewPassphrase(
			Words(wordCount),
		)
		ok(t, err)
		equals(t, DefaultExtra, phrase.extra)
		equals(t, wordCount, phrase.wordCount)
		equals(t, wordCount, len(phrase.words))
	}
}

func TestNewPassphraseWordsInvalid(t *testing.T) {
	testCases := []int{
		MinWords - 1,
	}

	for _, test := range testCases {
		phrase, err := NewPassphrase(
			Words(test),
		)
		assert(t, phrase == nil && err == ErrInvalidWordCount, "Expected function to error!")
	}
}

func TestPassphraseHumanize(t *testing.T) {
	testCases := []*Passphrase{}
	for i := 0; i < 100; i++ {
		phrase, err := NewPassphrase()
		ok(t, err)
		testCases = append(testCases, phrase)
	}

	for _, p := range testCases {
		str := ""
		for _, w := range p.words {
			str += string(w) + " "
		}
		str = strings.TrimSpace(str)
		equals(t, str, p.Humanize())
	}
}

func TestPassphraseString(t *testing.T) {
	testCases := []*Passphrase{}
	for i := 0; i < 100; i++ {
		phrase, err := NewPassphrase()
		ok(t, err)
		testCases = append(testCases, phrase)
	}

	for _, p := range testCases {
		str := ""
		for _, w := range p.words {
			str += string(w)
		}
		equals(t, str, p.String())
	}
}

func TestPassphraseRegenerate(t *testing.T) {
	phrase, err := NewPassphrase()
	ok(t, err)
	str := phrase.String()
	phrase.Regenerate()
	assert(t, phrase.String() != str, "Expected Regenerate() to create new, unique passphrase.")
}
