# typer-go

[![](https://goreportcard.com/badge/github.com/shilangyu/typer-go)](https://goreportcard.com/report/github.com/shilangyu/typer-go)
[![Build Status](https://travis-ci.com/shilangyu/typer-go.svg?branch=master)](https://travis-ci.com/shilangyu/typer-go)

Typer [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) game implemented in golang!

### usage

To build from source (needs [golang](https://golang.org/dl/) installed (v1.13+)):

```sh
go get github.com/shilangyu/typer-go
```

or get an executable from the [release tab](https://github.com/shilangyu/typer-go/releases) and put into PATH

Now just run `typer-go` in your terminal and the TUI will start. Full screen terminal is recommended.

### navigation

The whole TUI has mouse support!

| key               | description          |
| ----------------- | -------------------- |
| <kbd>↑</kbd>      | menu navigation up   |
| <kbd>↓</kbd>      | menu navigation down |
| <kbd>enter</kbd>  | confirm              |
| <kbd>tab</kbd>    | change view focus    |
| <kbd>ctrl+q</kbd> | back                 |
| <kbd>ctrl+c</kbd> | exit                 |

### Alternatives

- [typer-rs](https://github.com/krawieck/typer-rs)
