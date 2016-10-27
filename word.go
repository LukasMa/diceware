package diceware

// 0X1fff = 8191
const normalizer = 0X1fff

// A Word is a word from a list of words. A handful of these words build a
// diceware passphrase.
type Word string

// String implements the Stringer interface.
func (w Word) String() string {
	return string(w)
}

// GetWord retrieves the requested word from the standard word list. The
// function also handles integers that are overflowing the max id of 8191.
func GetWord(id int64) Word {
	// Normalize index.
	n := id & normalizer
	return diceware8k[n]
}
