# pramana-model

Shared contract for [Pramana](https://github.com/pramana-io/pramana) - the
core types every Pramana tool is built on.

This module defines the vocabulary of the system: what a resource is, what an
actor is, what an action is, and how authority and spend are represented. It
contains **types and interfaces only** - no logic, no I/O, no dependencies.
Everything else in Pramana fills these types rather than redefining them.

## Why it's separate

A cloud resource, a bot, and a factory machine are all `Actor`s performing
`Action`s under an `Authorization`, consuming `Spend`. Keeping that contract in
one dependency-free module means every other repo - the record, the scanner,
the language - agrees on the same shapes without depending on each other.

## The types

| Type            | Meaning                                  |
| --------------- | ---------------------------------------- |
| `Resource`      | A thing that exists                      |
| `Actor`         | A thing that acts                        |
| `Action`        | What was done                            |
| `Authorization` | What was granted, versus what was used   |
| `Spend`         | Budget consumption                       |
| `Snapshot`      | State at a moment in time                |
| `Gap`           | What is absent from the record, and why  |

## Install

```bash
go get github.com/pramana-io/pramana-model
```

## Status

Pre-v1. Types are being added incrementally; the contract is not yet stable.

## License

Apache-2.0