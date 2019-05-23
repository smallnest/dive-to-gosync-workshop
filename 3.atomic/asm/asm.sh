# go tool dist list

# sudo ./buildpkg.sh darwin amd64
# sudo ./buildpkg.sh windows 386
# sudo ./buildpkg.sh windows arm
# sudo ./buildpkg.sh linux arm64
# sudo ./buildpkg.sh linux mips
# sudo ./buildpkg.sh linux mips64


指令后缀的意义：

```go
// 后缀 https://en.wikibooks.org/wiki/X86_Assembly/GAS_Syntax#Operation_Suffixes
//
// b = byte (8 bit)
// s = single (32-bit floating point)
// w = word (16 bit)
// l = long (32 bit integer or 64-bit floating point)
// q = quad (64 bit)
// t = ten bytes (80-bit floating point)
```

## darwin/amd64
GOOS=darwin GOARCH=amd64 go tool compile -S  main.go

## windows/386
GOOS=windows GOARCH=386 go tool compile -S  main.go

## windows arm
GOOS=windows GOARCH=arm go tool compile -S  main.go

## linux arm64
GOOS=linux GOARCH=arm64 go tool compile -S  main.go

## linux mips
GOOS=linux GOARCH=mips go tool compile -S  main.go

## linux mips64
GOOS=linux GOARCH=mips64 go tool compile -S  main.go

