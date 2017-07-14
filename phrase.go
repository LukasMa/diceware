package diceware

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"strings"
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

	// DefaultValidate is the default value for the validation step.
	DefaultValidate = true

	// MinPhraseLength is the smallest amount of characters allowed in a
	// passphrase to pass the validation. Since generation is random, there is a
	// very small chance of getting a passphrase which has less than 17
	// characters in total which IS NOT considered save.
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
	ErrInvalidWordCount = errors.New("diceware: amount of words is invalid")

	// ErrValidationFailed is raised when the generated passphrase doesn't met
	// the default security standards.
	ErrValidationFailed = errors.New("diceware: invalid passphrase was generated")
)

// An Option serves as a functional parameter which can be used to costumize the
// generation of the passphrase.
type Option func(p *Passphrase) error

// Extra is an Option that specifies whaether an extra will be added to the
// passphrase or not.
func Extra(extra bool) Option {
	return func(p *Passphrase) error { return p.setExtra(extra) }
}
func (p *Passphrase) setExtra(extra bool) error {
	p.extra = extra
	return nil
}

// Validate is an Option that specifies whaether passphrase validation will be
// performed or not.
func Validate(validate bool) Option {
	return func(p *Passphrase) error { return p.setValidate(validate) }
}
func (p *Passphrase) setValidate(validate bool) error {
	p.validate = validate
	return nil
}

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

// A Passphrase is a diceware passphrase. It is build from a handful of words
// that are randomly picked from a list of words.
// Ref: http://world.std.com/~reinhold/diceware.html
type Passphrase struct {
	extra     bool
	validate  bool
	wordCount int
	words     []string
}

// NewPassphrase defines, generates, validates and returns a new diceware
// passphrase.
func NewPassphrase(options ...Option) (*Passphrase, error) {
	// Create passphrase with default settings.
	p := &Passphrase{
		extra:     DefaultExtra,
		validate:  DefaultValidate,
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
	if err := p.Regenerate(); err != nil {
		return nil, err
	}

	// Return passphrase.
	return p, nil
}

// Humanize will return a human readable string which has a whitspace between
// each word.
func (p Passphrase) Humanize() string {
	str := ""
	for _, word := range p.words {
		str += word + " "
	}
	return strings.TrimSpace(str)
}

// String implements the Stringer interface.
func (p Passphrase) String() string {
	str := ""
	for _, word := range p.words {
		str += word
	}
	return str
}

// Regenerate will generate the passphrase from scratch but keep the originally
// provided parameters.
func (p *Passphrase) Regenerate() error {
	// Re(generate) passphrase.
	if err := p.generate(); err != nil {
		return err
	}

	// Validate passphrase.
	if p.validate && !p.Validate() {
		return ErrValidationFailed
	}

	return nil
}

// Validate verifies that the passphrase mets certain standards like a secure
// length and word count.
func (p *Passphrase) Validate() bool {
	return MinPhraseLength <= len(p.String()) && DefaultWords <= p.wordCount
}

func (p *Passphrase) generate() error {
	p.words = make([]string, p.wordCount)
	for i := 0; i < p.wordCount; i++ {
		id, err := generateID(math.MaxInt64)
		if err != nil {
			return err
		}
		p.words[i] = getWord(id)
	}

	if p.extra {
		id, err := generateID(36)
		if err != nil {
			return err
		}
		wc, err := generateID(int64(len(p.words)))
		if err != nil {
			return err
		}
		p.words[wc] += extras[id]
	}

	return nil
}

func generateID(from int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(from))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// getWord retrieves the requested word from the standard word list. The
// function also handles integers that are overflowing the max id of 8191.
func getWord(id int64) string {
	return diceware8k[id&8191]
}
