package diceware

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestNewPassphrase(t *testing.T) {
	// Should create passphrase with default settings.
	phrase, err := NewPassphrase(
		Validate(false),
	)
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
			Validate(false),
		)
		ok(t, err)
		equals(t, test, phrase.extra)
	}
}

func TestNewPassphraseWords(t *testing.T) {
	testCases := []struct {
		observed    int
		expectedErr error
	}{
		{MinWords, nil},
		{MinWords + 1, nil},
		{DefaultWords - 1, nil},
		{DefaultWords, nil},
		{DefaultWords + 1, nil},
		{DefaultWords + 10, nil},
		{MinWords - 1, ErrInvalidWordCount},
	}

	for _, test := range testCases {
		phrase, err := NewPassphrase(
			Words(test.observed),
			Validate(false),
		)
		equals(t, test.expectedErr, err)
		if err == nil {
			equals(t, DefaultExtra, phrase.extra)
			equals(t, test.observed, phrase.wordCount)
			equals(t, test.observed, len(phrase.words))
		}
	}
}

func TestPassphraseHumanize(t *testing.T) {
	testCases := []*Passphrase{}
	for i := 0; i < 100; i++ {
		phrase, err := NewPassphrase(
			Validate(false),
		)
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
		phrase, err := NewPassphrase(
			Validate(false),
		)
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
	phrase, err := NewPassphrase(
		Validate(false),
	)
	ok(t, err)
	phraseStr := phrase.String()
	phrase.Regenerate()
	assert(t, phrase.String() != phraseStr, "Expected Regenerate() to create new, unique passphrase.")
}

func TestPassphraseValidate(t *testing.T) {
	testCases := []struct {
		phrase      *Passphrase
		expectedRes bool
	}{
		// Ok.
		{
			phrase:      &Passphrase{DefaultExtra, DefaultValidate, 6, []string{"11111", "22222", "33333", "44444", "55555", "66666"}},
			expectedRes: true,
		},
		// To few words.
		{
			phrase:      &Passphrase{DefaultExtra, DefaultValidate, 5, []string{"11111", "22222", "33333", "44444", "55555"}},
			expectedRes: false,
		},
		// To few characters.
		{
			phrase:      &Passphrase{DefaultExtra, DefaultValidate, 6, []string{"1", "2", "3", "4", "5", "6"}},
			expectedRes: false,
		},
	}

	for _, test := range testCases {
		equals(t, test.expectedRes, test.phrase.Validate())
	}
}

func TestGetWord(t *testing.T) {
	tc := map[int]string{
		8192:     "a",
		8192 * 2: "a",
	}

	for id, word := range tc {
		w := getWord(int64(id))
		equals(t, word, w)
	}
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
