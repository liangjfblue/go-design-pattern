# 包装函数

源码例子：go1.13.7 Go/src/io/io.go:430

包装函数主要是对某个函数封装一层，把被包装函数传入包装函数，使包装函数具有被包装函数的功能外，
还可以额外做处理来适应新的需求。

应用场景：
- 调用第三方API。可能不完全符合自己的业务，比如第三方API返回的是A， 自己的业务需要在调用API时加入自己的业务信息
- 使用第三方API。对方修改了某参数类型，但是自己的业务系统大量引用，不想改变那么可以包装一层API，适应自己的系统
- 拓展函数的能力。对某个接口已有的功能进行拓展
- ...

所以包装函数是非常实用的，可能在实际的开发中你应用很多，但是可能一时没有注意到。

在go源码中，特别是io库，就有包装函数的应用。


    func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }
    
    type LimitedReader struct {
        R Reader // underlying reader
        N int64  // max bytes remaining
    }
    
    func (l *LimitedReader) Read(p []byte) (n int, err error) {
        if l.N <= 0 {
            return 0, EOF
        }
        if int64(len(p)) > l.N {
            p = p[0:l.N]
        }
        n, err = l.R.Read(p)
        l.N -= int64(n)
        return
    }

Reader 是一个接口。

    type Reader interface {
        Read(p []byte) (n int, err error)
    }

LimitedReader实现的包装函数的作用是封装一层代码，拓展Reader的Read接口能力，满足任意实现Read接口的Reader实现类使用

原理主要是定义一个LimitedReader结构，引用了Reader的Read接口，对外部传入的实现Read接口的Reader实现类做截取指定n个字符的逻辑

go实现包扎黄函数的重点：
- 定义一个结构XXX，拥有需要包装的函数的接口
- 定义一个NewXXX函数，提供参数，用于外部传入被包装函数的接口
- 实现接口

