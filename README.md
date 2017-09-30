[![Build Status](https://travis-ci.org/kcmerrill/kronk.svg?branch=master)](https://travis-ci.org/kcmerrill/kronk) [![Join the chat at https://gitter.im/kcmerrillkronk](https://badges.gitter.im/kronk.svg)](https://gitter.im/kcmerrillkronk?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![kronk](assets/kronk.jpg "kronk")

# kronk

Simple(r) text searching. A [Marvin](https://github.com/kcmerrill/marvin) companion app.


## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/kronk/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/kronk/linux/amd64)

via go:

`$ go get -u github.com/kcmerrill/kronk`

## Usage

`kronk` requires input via `stdin`. Many ways to get data via stdin, but in this example we'll simply `cat` or `curl` a file into `kronk`.

```
$> stdin | kronk <arguments> <name:regular-expression>...
```

A simple 1 match example:
```bash
$> curl https://api.github.com/users/kcmerrill/repos | kronk 'repo:"full_name": "(.*?)"'
```

Multiple matches(must all yeild the same number of results)
```bash
$> curl https://api.github.com/users/kcmerrill/repos | kronk 'repo:"full_name": "(.*?)"' 'issues:"open_issues": (\d+)'
```

Using [Marvin](https://github.com/kcmerrill/marvin)? Need dynamic inventory? Just remember to use the appropriate `del` based on your needs.
```bash
$> curl https://api.github.com/users/kcmerrill/repos | kronk --out inline 'repo:"full_name": "(.*?)"' 'issues:"open_issues": (\d+)'
```

## Screencast/Demo

[![asciicast](assets/demo.png)](https://asciinema.org/a/140001)
