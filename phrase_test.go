package diceware_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/lukasmalkmus/diceware"
)

func TestNewPassphrase(t *testing.T) {
	tests := []struct {
		words       int
		expectedErr error
	}{
		{diceware.MinWords, nil},
		{diceware.MinWords + 1, nil},
		{diceware.DefaultWords - 1, nil},
		{diceware.DefaultWords, nil},
		{diceware.DefaultWords + 1, nil},
		{diceware.DefaultWords + 10, nil},
		{diceware.MinWords - 1, diceware.ErrInvalidWordCount},
	}

	for _, tt := range tests {
		phrase, err := diceware.NewPassphrase(
			diceware.Words(tt.words),
			diceware.Validate(false),
		)
		equals(t, tt.expectedErr, err)
		if err == nil {
			equals(t, tt.words, len(strings.Fields(phrase.Humanize())))
		}
	}
}

func TestPassphrase_Regenerate(t *testing.T) {
	phrase, err := diceware.NewPassphrase(
		diceware.Validate(false),
	)
	ok(t, err)
	phraseStr := phrase.String()
	phrase.Regenerate()
	assert(t, phrase.String() != phraseStr, "Expected Regenerate() to create new, unique passphrase.")
}

func BenchmarkPassphrase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		diceware.NewPassphrase()
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
