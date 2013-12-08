# Reversal

Simple reverse proxy, written in Go

## Usage

    reversal [local-port] [destination]

For example, to forward `localhost:8000` to `192.168.1.10:5000`

    reversal :8000 http://192.168.1.10:5000

## Downloads

- [Windows]()
- [OSX]()
- [Linux]()

More platform-specific builds available in [the latest release]().

Builds powered by [gox](https://github.com/mitchellh/gox)
