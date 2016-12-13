package diceware

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
)

const (
	// DefaultExtra is the default value for the extra character. An extra can
	// be added to a passphrase to increase security without adding another
	// word. It isn't required by default.
	DefaultExtra = false

	// DefaultWords is the default amount of words used to build a passphrase.
	// This is set to a sensitive default.
	// Ref: https://diceware.blogspot.de/2014/03/time-to-add-word.html
	DefaultWords = 6

	// MinPhraseLength is the smallest amount of characters allowed in a
	// passphrase. Since generation is random, there is a very small chance of
	// getting a passphrase which has less than 17 characters in total which IS
	// NOT considered save.
	// Ref: http://world.std.com/~reinhold/dicewarefaq.html#14characters
	MinPhraseLength = 17

	// MinWords is the required amount of words used to build a passphrase. This
	// values exists just for convenience and it IS NOT SAFE to use a one word
	// passphrase!
	MinWords = 1
)

var (
	// ErrInvalidWordCount is raised when the specified amount of words drops
	// below the MinWords constant.
	ErrInvalidWordCount = errors.New("The amount of words is invalid.")
)

// An Option serves as a functional parameter which can be used to costumize the
// generation of the passphrase.
type Option func(p *Passphrase) error

// Words is an Option that defines the amount of words that should be picked for
// the new Passphrase.
func Words(words int) Option {
	return func(p *Passphrase) error { return p.setWords(words) }
}
func (p *Passphrase) setWords(words int) error {
	if words < MinWords {
		return ErrInvalidWordCount
	}
	p.wordCount = words
	return nil
}

// Extra is an Option that specifies whaether an extra will be added to the
// passphrase or not.
func Extra(extra bool) Option {
	return func(p *Passphrase) error { return p.setExtra(extra) }
}
func (p *Passphrase) setExtra(extra bool) error {
	p.extra = extra
	return nil
}

// A Passphrase is a diceware passphrase. It is a build from a handful of words
// that are randomly picked from a list of words.
// Ref: http://world.std.com/~reinhold/diceware.html
type Passphrase struct {
	extra     bool
	wordCount int
	words     []Word
}

// NewPassphrase defines, generates, validates and returns a new diceware
// passphrase.
func NewPassphrase(options ...Option) (*Passphrase, error) {
	// Create passphrase with default settings.
	p := &Passphrase{
		extra:     DefaultExtra,
		wordCount: DefaultWords,
		words:     nil,
	}

	// Apply supplied options.
	for _, option := range options {
		if err := option(p); err != nil {
			return nil, err
		}
	}

	// Generate passphrase.
	p.generate()

	// Return passphrase.
	return p, nil
}

// String implements the Stringer interface.
func (p Passphrase) String() string {
	str := ""
	for _, word := range p.words {
		str += word.String()
	}
	return str
}

// Regenerate will generate the passphrase from scratch but keep the originally
// provided parameters.
func (p *Passphrase) Regenerate() {
	p.generate()
}

func (p *Passphrase) generate() {
	p.words = nil
	for i := 0; i < p.wordCount; i++ {
		id := generateID()
		p.words = append(p.words, GetWord(id))
	}
}

func generateID() int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
