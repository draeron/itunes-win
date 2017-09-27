# itunes-win

[![Build status](https://ci.appveyor.com/api/projects/status/t9wbgwo0de3b2o0n?svg=true)]
[![GoDoc](https://godoc.org/github.com/tdsh/itunes-win/itunes?status.svg)](https://godoc.org/github.com/tdsh/itunes-win/itunes)

Package itunes-win enables you to handle iTunes on Windows in Go. This is
a thin wrapper of iTunes COM interface.

Note this library doesn't relect all of the methods and the properties of
iTunes COM interface.

## Installation

    go get github.com/tdsh/itunes-win/itunes

## License

Under MIT License: https://tdsh.mit-license.org/

## Documentation

- [Reference](https://godoc.org/github.com/tdsh/itunes-win/itunes)
- Examples
  - [Control iTunes](https://github.com/tdsh/itunes-win/blob/master/examples/itunes_example.go)
  - [Lists tracks released in 1990s](https://github.com/tdsh/itunes-win/blob/master/examples/list_90s_music.go)
  - [Lists tracks whose files are no longer available](https://github.com/tdsh/itunes-win/blob/master/examples/list_dead_tracks.go)
  - [Control playlist](https://github.com/tdsh/itunes-win/blob/master/examples/playlist_example.go)
  - [Lists file location of specified tracks](https://github.com/tdsh/itunes-win/blob/master/examples/search_location.go)
