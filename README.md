# lukasmalkmus/diceware
> Diceware passphrases in go. - by **[Lukas Malkmus](https://github.com/lukasmalkmus)**

[![Travis Status][travis_badge]][travis]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![GoDoc][docs_badge]][docs]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]

---

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Usage](#usage)
4. [Contributing](#contributing)
5. [License](#License)

### Introduction
Package **diceware** is a simple implementation of the
[diceware passphrase generation method](http://world.std.com/~reinhold/diceware.html).
Instead of using the normal wordlists, it uses the computer-optimized
[diceword8k list](http://world.std.com/%7Ereinhold/dicewarefaq.html#diceware8k).
Furhtermore it utilizes go's `crypto/rand` library to generate true random
passphrases.

Be advised, that the prefered way of generating diceware passphrases is to do it
the old-school way by actually throwing real dices by hand. This is the only
100% secure way.

### Features
- [x] Simple API
- [x] Only go standard library
- [x] Passphrases with choosable length
- [x] Diceware extras for stronger passphrases
- [x] Verify passphrases

#### Todo
- [ ] Multiple word lists in multiple languages
- [ ] Read word list from file/buffer (`io.Reader`)

### Usage
#### Installation
Please use a dependency manager like [glide](http://glide.sh) to make sure you
use a tagged release.

Install using `go get`:
```bash
go get -u -v github.com/lukasmalkmus/diceware
```

#### Creation
Create a passphrase with default values (6 words, no extra):
```go
p, err := diceware.NewPassphrase()
if err != nil {
    // ...
}
fmt.Println(p)
```

It is also possible to create a passphrase with more or less words. Please note
that 6 words is a sensitiv default and less isn't recommended!
```go
p, err := diceware.NewPassphrase(
    diceware.Words(7), // Passphrase with 7 words
)
if err != nil {
    // ...
}
fmt.Println(p)
```

**Note!** If you want to use less than 6 words, be sure to set the `Validate` option
to `false`! Otherwise _validation will fail_!

#### Regeneration
All passphrases can be _regenerated_. This means the options you applied in the
`NewPassphrase()` function are reused for the passphrase generation.
```go
p.Regenerate()
fmt.Println(p)
```

#### Tips & Tricks
- Passphrase implements the Stringer interface thus it can be passed to every
function accepting this interface. For example `fmt.Println()`.
- The `String()` method isn't very "human friendly". Use the `Humanize()` method
to print the passphrase with whitspace seperated words.
- Passphrase strength can be improved by adding an extra. Do this by setting the
Extra option: `Extra(true)`

### Contributing
Feel free to submit PRs or to fill Issues. Every kind of help is appreciated.

### License
© Lukas Malkmus, 2017

Distributed under MIT License (`The MIT License`).

See [LICENSE](LICENSE) for more information.


[travis]: https://travis-ci.org/lukasmalkmus/diceware
[travis_badge]: https://travis-ci.org/lukasmalkmus/diceware.svg
[coverage]: https://coveralls.io/github/lukasmalkmus/diceware?branch=master
[coverage_badge]: https://coveralls.io/repos/github/lukasmalkmus/diceware/badge.svg?branch=master
[report]: https://goreportcard.com/report/github.com/lukasmalkmus/diceware
[report_badge]: https://goreportcard.com/badge/github.com/lukasmalkmus/diceware
[docs]: https://godoc.org/github.com/lukasmalkmus/diceware
[docs_badge]: https://godoc.org/github.com/lukasmalkmus/diceware?status.svg
[release]: https://github.com/lukasmalkmus/diceware/releases
[release_badge]: https://img.shields.io/github/release/lukasmalkmus/diceware.svg
[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/badge/license-MIT-blue.svg
