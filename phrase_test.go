package diceware

import "testing"

func TestNewPassphrase(t *testing.T) {
	// Should create passphrase with default settings.
	phrase, err := NewPassphrase()
	ok(t, err)
	equals(t, DefaultExtra, phrase.extra)
	equals(t, DefaultWords, phrase.wordCount)
	equals(t, DefaultWords, len(phrase.words))
}

func TestNewPassphraseExtra(t *testing.T) {
	tc := []bool{
		DefaultExtra,
		true,
		false,
	}

	for _, v := range tc {
		phrase, err := NewPassphrase(
			Extra(v),
		)
		ok(t, err)
		equals(t, v, phrase.extra)
	}
}

func TestNewPassphraseWordsValid(t *testing.T) {
	tc := []int{
		MinWords,
		MinWords + 1,
		DefaultWords - 1,
		DefaultWords,
		DefaultWords + 1,
		DefaultWords + 10,
	}

	for _, wordCount := range tc {
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
	tc := []int{
		MinWords - 1,
	}

	for _, wordCount := range tc {
		phrase, err := NewPassphrase(
			Words(wordCount),
		)
		assert(t, phrase == nil && err == ErrInvalidWordCount, "Expected function to error!")
	}
}

func TestPassphraseRegenerate(t *testing.T) {
	phrase, err := NewPassphrase()
	ok(t, err)
	str := phrase.String()
	phrase.Regenerate()
	assert(t, phrase.String() != str, "Expected Regenerate() to create new, unique passphrase.")
}

func TestPassphraseString(t *testing.T) {
	tc := []*Passphrase{}
	for i := 0; i < 100; i++ {
		phrase, err := NewPassphrase()
		ok(t, err)
		tc = append(tc, phrase)
	}

	for _, p := range tc {
		str := ""
		for _, w := range p.words {
			str += string(w)
		}
		equals(t, str, p.String())
	}
}
