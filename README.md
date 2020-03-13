# typer-go

[![](https://goreportcard.com/badge/github.com/shilangyu/typer-go)](https://goreportcard.com/report/github.com/shilangyu/typer-go)
[![](https://github.com/shilangyu/typer-go/workflows/ci/badge.svg)](https://github.com/shilangyu/typer-go/actions)

Test your typing speed in a Typer [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) game!

- [typer-go](#typer-go)
	- [features](#features)
	- [install](#install)
	- [usage](#usage)
	- [settings](#settings)
	- [navigation](#navigation)

## features

- collect statistics to track your progress
  - Words per minute
  - Timestamps
  - Amount of mistakes for each word
  - Amount of time for each word
- play with your friends in local multiplayer (global coming soon)
- customizable
- comfort of your sweet sweet terminal

## install

Grab an executable from the [release tab](https://github.com/shilangyu/typer-go/releases)

... or if you're an Gopher build from source (requires Go v1.13+):

```sh
go get github.com/shilangyu/typer-go
```

## usage

Just run `typer-go` in your terminal and the TUI will start. Full screen terminal is recommended. There are no CLI commands and flags.

## settings

| name          | values                     | description                                                                                                                                  |
| ------------- | -------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| highlight     | `'background'` or `'text'` | The way the already typed text should be highlighted. `'background'` to fill the background, `'text'` to just fill the text.                 |
| error display | `'typed'` or `'text'`      | what should be shown when incorrect character is inputted. `'typed'` will show the typed char, `'text'` will show what should've been typed. |
| texts path    | any string                 | path to your custom typer texts where each text is separated by a new line. If path is empty, preloaded texts will be loaded.                |

## navigation

[~~The whole TUI has mouse support!~~](https://github.com/shilangyu/typer-go/issues/9)

| key               | description          |
| ----------------- | -------------------- |
| <kbd>↑</kbd>      | menu navigation up   |
| <kbd>↓</kbd>      | menu navigation down |
| <kbd>enter</kbd>  | confirm              |
| <kbd>esc</kbd>    | back                 |
| <kbd>tab</kbd>    | switch focus         |
| <kbd>ctrl+c</kbd> | exit                 |
