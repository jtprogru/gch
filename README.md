# gch

[![goreleaser](https://github.com/jtprogru/gch/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/jtprogru/gch/actions/workflows/goreleaser.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/jtprogru/gch)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jtprogru/gch)

Go CLI Helper (gch) this is a simple CLI utility that helps make my life easier and will be gradually supplemented with various functionality.

Now gch is not able to do so much, but I use it every day.

## Usage

```shell
gch

Go CLI Helper this is a simple CLI utility that helps
make my life easier and will be gradually supplemented with various functionality.

Now gch is not able to do so much, but I use it every day.

Complete documentation is available at https://github.com/jtprogru/gch

Usage:
  gch [command]

Available Commands:
  breif       Generate a short description for your long URL
  cbrf        Get currency exchange rates for RUB/USD and RUB/EUR
  completion  Generate the autocompletion script for the specified shell
  dupl        Show all duplicates JPG and PNG in folder
  help        Help about any command
  passwd      Generate random password
  short       Make short link from URL
  sretask     Create template with SRE task

Flags:
      --config string   config file (default is $HOME/.gch.yaml)
  -h, --help            help for gch

Use "gch [command] --help" for more information about a command.

```

## Installation

For installation you need to load latest version from [Release](https://github.com/jtprogru/gch/releases) page and download version for you platform.

Another way is usage `go install` – for more details see project [Wiki](https://github.com/jtprogru/gch/wiki#installation).

## Project Status

![Alt](https://repobeats.axiom.co/api/embed/90f398a2bc0fb93e055987ed40743d2f318e2ebc.svg "Repobeats analytics image")

## License

[MIT](LICENSE)
