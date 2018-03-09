Golang sizeof tips
------------------

**Web tool for interactive playing with Golang struct sizes.**


## Aim
Provide comfortable tool to see how fields in struct are aligned,
to compare different structs and as the result - to understand
and remember alignment rules.

## Installing
To install correct versions of dependencies
[Dep dependency manager](https://github.com/golang/dep) should be used.
```bash
brew install dep
go get github.com/l3x/sizeof-struct
cd github.com/l3x/sizeof-struct
dep ensure
go build -o $GOBIN/server
```
You may also install via simple `go get` by your own risk.


## Usage
```bash
./server -http=:7777 start
./server stop
./server restart
```

## Platform support
Tested on Linux and OS X x64 platforms, but should work properly and on other
*nix-like platforms.
Windows is not supported due to daemonization.

## License
[Apache License 2.0](LICENSE)
