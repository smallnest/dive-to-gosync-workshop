## 架构

查看:

```sh
go tool dist list
```

安装指定架构的标准库

```
./buildpkg.sh $GOOS $GOARCH
```

具体实现: https://github.com/golang/go/tree/master/src/runtime/internal/atomic

## darwin/amd64

add:
```asm
LOCK
XADDQ	CX, (AX)
```

load:
```asm
MOVQ	(AX), AX
```asm

cas:
```asm
LOCK
CMPXCHGQ	DX, (CX)
```

store:
```asm
XCHGQ	CX, (AX)
```

swap:
```asm
XCHGQ	CX, (AX)
```


## windows/386

add:
```asm
sync/atomic.AddInt64(SB)
```

load:
```asm
sync/atomic.LoadInt64(SB)
```asm

cas:
```asm
sync/atomic.CompareAndSwapInt64(SB)
```

store:
```asm
sync/atomic.StoreInt64(SB)
```

swap:
```asm
sync/atomic.SwapInt64(SB)
```

https://github.com/golang/go/blob/master/src/runtime/internal/atomic/asm_386.s

## windows arm

add:
```asm
sync/atomic.AddInt64(SB)
```

load:
```asm
sync/atomic.LoadInt64(SB)
```asm

cas:
```asm
sync/atomic.CompareAndSwapInt64(SB)
```

store:
```asm
sync/atomic.StoreInt64(SB)
```

swap:
```asm
sync/atomic.SwapInt64(SB)
```


## linux arm64
`LDXR/STXR`、`LDAXR/STLXR`

## linux mips

add:
```asm
sync/atomic.AddInt64(SB)
```

load:
```asm
sync/atomic.LoadInt64(SB)
```asm

cas:
```asm
sync/atomic.CompareAndSwapInt64(SB)
```

store:
```asm
sync/atomic.StoreInt64(SB)
```

swap:
```asm
sync/atomic.SwapInt64(SB)
```

## linux mips64

`SYNC`
