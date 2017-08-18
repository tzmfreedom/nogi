# Nogi

Nogizaka 46 Command Line Interface

## Install

```
$ go get github.com/tzmfreedom/nogi
```

## Usage

Get Members
```
$ nogi members
```

Get Songs
```
$ nogi songs
```

## Contribute

Just send pull request if needed or fill an issue!

### How to add songs and members

`nogi` stores songs and members data with go-bindata.

If you want to modify songs or members, please modify toml file and recreate go-bindata.

1. Add songs/members to data/songs.toml or data/members.toml.
```toml
[[songs]]
name = "逃げ水"
no = 19
```

2. Execute [go-bindata](https://github.com/jteeuwen/go-bindata) to create bindata.go.
```bash
$ go-bindata data/ # or make build
```

## License

The MIT License See [LICENSE](https://github.com/tzmfreedom/nogi/blob/master/LICENSE) file.
