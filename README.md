# gch

[![goreleaser](https://github.com/jtprogru/gch/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/jtprogru/gch/actions/workflows/goreleaser.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/jtprogru/gch)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jtprogru/gch)

Go CLI Helper (gch) this is a simple CLI utility that helps make my life easier and will be gradually supplemented with various functionality.

Now gch is not able to do so much, but I use it every day.

## Installation

For installation you need to load latest version from [Release](https://github.com/jtprogru/gch/releases) page and download version for you platform.

Another way is usage `go install`:

```shell
# Get latest version from CLI
VERSION=`curl -sSL https://api.github.com/repos/jtprogru/gch/releases/latest -s | jq .name -r`
go install github.com/jtprogru/gch@$VERSION
```

## Project Status

![Alt](https://repobeats.axiom.co/api/embed/90f398a2bc0fb93e055987ed40743d2f318e2ebc.svg "Repobeats analytics image")

## License

MIT
