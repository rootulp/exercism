# Exercism

[![License](https://img.shields.io/:license-mit-blue.svg)](https://rootulp.mit-license.org)

[Exercism](http://exercism.io/rootulp) solutions

## Getting Started

If you are new to exercism, you can get started with this [intro](http://exercism.io/how-it-works/newbie).

## Tests

![Go](https://github.com/rootulp/exercism/workflows/Go/badge.svg)
![Ruby](https://github.com/rootulp/exercism/workflows/Ruby/badge.svg?branch=master)
![Python](https://github.com/rootulp/exercism/workflows/Python/badge.svg?branch=master)
![Rust](https://github.com/rootulp/exercism/workflows/Rust/badge.svg?branch=master)

## Code Style

| language     | code style & linter                                                | command                                                                          |
| ------------ | ------------------------------------------------------------------ | -------------------------------------------------------------------------------- |
| coffeescript | [coffeelint](https://github.com/clutchski/coffeelint)              | `yarn lint`                                                                      |
| ecmascript   | [airbnb style guide](https://github.com/airbnb/javascript)         | `cd <exercise> && yarn lint`                                                     |
| elm          | [elm-format](https://github.com/avh4/elm-format)                   | `elm-format .`                                                                   |
| go           | [gofmt](https://golang.org/cmd/gofmt/)                             | `go fmt .`                                                                       |
| java         | [google-java-format](https://github.com/google/google-java-format) | `java -jar google-java-format-1.8-all-deps.jar  -r ./**/src/main/java/*.java -r` |
| javascript   | [prettier](https://github.com/prettier/prettier)                   | `yarn lint`                                                                      |
| python       | [pep 8](https://www.python.org/dev/peps/pep-0008/)                 | `pycodestyle .` or `autopep8 --in-place --aggressive --aggressive --recursive .` |
| ruby         | [rubocop](https://github.com/bbatsov/rubocop)                      | `rubocop .`                                                                      |
| rust         | [rustfmt](https://github.com/rust-lang/rustfmt)                    | `cargo fmt --all`                                                                |
| typescript   | [tslint](https://github.com/palantir/tslint)                       | `cd <exercise> && yarn lint`                                                     |

## Contribute

I'd appreciate any feedback via [issues](https://github.com/rootulp/exercism/issues/new).
