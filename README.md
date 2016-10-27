# lukasma/diceware
> Diceware passphrases in go. - by **[Lukas Malkmus](https://github.com/LukasMa)**

[![Build Status][build_badge]][build]
[![Coverage Status][coverage_badge]][coverage]
[![Go Report][report_badge]][report]
[![GoDoc][docs_badge]][docs]
[![Latest Release][release_badge]][release]
[![License][license_badge]][license]

---

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Contributing](#contributing)
6. [License](#License)

### Introduction
Package **diceware** is a simple implementation of the [diceware passphrase generation method](http://world.std.com/~reinhold/diceware.html). Instead of using the normal wordlists, it uses the computer-optimized [diceword8k list](http://world.std.com/%7Ereinhold/dicewarefaq.html#diceware8k). Furhtermore it utilizes go's `crypto/rand` library to generate true random passphrases.

Be advised, that the prefered way of generating diceware passphrases is to do it the old-school way by actually throwing real dices by hand. This is the only 100% secure way.

### Features
- [x] Simple API
- [x] Only go standard library
- [x] Passphrases with choosable length
- [ ] Diceware extras for stronger passphrases
- [ ] Verify passphrases
- [ ] Multiple word lists in multiple languages
- [ ] Read word list from file/buffer (`io.Reader`)

### Installation
The easiest way to install this package is to use `go get`:
```go
go get -u -v github.com/LukasMa/diceware
```
Since this will pull the master branch, you should use a dependency manager like [glide](http://glide,sh) to be on the safe site.

### Usage

#### Creating passphrases
Create a passphrase with default values (6 words, no extra):
```go
p, err := diceware.NewPassphrase()
if err != nil {
    // ...
}
fmt.Println(p)
```

It is also possible to create a passphrase with more or less words. Please note that 6 words is a sensitiv default and less isn't recommended!
```go
p, err := diceware.NewPassphrase(
    diceware.Words(7), // Passphrase with 7 words
)
if err != nil {
    // ...
}
fmt.Println(p)
```

#### Regeneration
All passphrases can be _regenerated_. This means the options you applied in the `NewPassphrase()` function are reused for the passphrase generation.
```go
p.Regenerate()
fmt.Println(p)
```

### Contributing
Please feel free to submit Pull Requests or Issues.

### License
Copyright (c) 2016 Lukas Malkmus

Distributed under MIT License (`The MIT License`). See [LICENSE](LICENSE) for more information.

[license]: https://opensource.org/licenses/MIT
[license_badge]: https://img.shields.io/badge/liecense-MIT-blue.svg
[docs]: https://godoc.org/github.com/LukasMa/diceware
[docs_badge]: https://godoc.org/github.com/LukasMa/diceware?status.svg
[release]: https://github.com/LukasMa/diceware/releases
[release_badge]: https://img.shields.io/github/release/LukasMa/diceware.svg
[report]: https://goreportcard.com/report/github.com/LukasMa/diceware
[report_badge]: https://goreportcard.com/badge/github.com/LukasMa/diceware
[build]: https://travis-ci.org/LukasMa/diceware
[build_badge]: https://travis-ci.org/LukasMa/diceware
[coverage]: https://coveralls.io/github/LukasMa/diceware
[coverage_badge]: https://coveralls.io/repos/github/LukasMa/diceware/badge
