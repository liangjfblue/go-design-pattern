# 适配器

在设计模式中也有适配器模式。

> 定义为：适配器模式（Adapter Pattern）是作为两个不兼容的接口之间的桥梁。将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能在一起工作的那些类可以一起工作

在这里思想也是适用的。其中最能代表的是go中http的例子

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
    
    type HandlerFunc func(ResponseWriter, *Request)
    
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }


在用标准库的http实现web时最终都会调用，来启动监听web服务

    func ListenAndServe(addr string, handler Handler) error

其中在加入路由函数时有两种途径：

第一种：

    func TestAdapterFunction1(t *testing.T) {
        http.HandleFunc("/tets", func(w http.ResponseWriter, r *http.Request) {
            w.Write([]byte("hello\n"))
        })
    
        http.ListenAndServe(":8080", nil)
    }


第二种，自定义handler

    type H struct{}
    
    func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("this is version 2"))
    }
    
    func hello(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello\n"))
    }
    
    func TestAdapterFunction2(t *testing.T) {
        mux := http.NewServeMux()
        mux.Handle("/", &H{})
        mux.HandleFunc("/bye", hello)
        http.ListenAndServe(":8080", mux)
    }

看起来，第一种是非常简单的，单纯的传入一个 `func(ResponseWriter, *Request)` 类型的函数就ok了

第二种要自己实现 handler的接口，即

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }


追踪源码下去，都是调用`Handler接口`的`ServeHTTP`函数。因此其实两种方法最终实现的目的都是一样的。
`serverHandler{c.server}.ServeHTTP(w, w.req)`

第一种是用默认的handler，第二种是自定义的，可以自己实现对应的路由管理。比如添加和查询等，gin就是自己实现了一个handler

    g := gin.Default()
    g.GET("/test", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello\n"))
    })
    http.ListenAndServe(":8080", g)


**正题来了，为什么第一种方法这么简便呢？**

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }
    
    type HandlerFunc func(ResponseWriter, *Request)
    
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }
    
    func ListenAndServe(addr string, handler Handler) error
    
再看回这段代码，这里是定义了`HandlerFunc类型`，只要是实现了`Handler接口`的对象，通过下面的方式传入

    http.HandleFunc("/tets", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello\n"))
    })

内部都会转换成符合`http.Handler接口`的类型，这样在最终调用的时候，
`serverHandler{c.server}.ServeHTTP(w, w.req)`，都会调用真正传入的处理函数

这种通过接口类型来转换实现某接口类型的对象的方法，就是适配类型 `adapter function type`

## adapter function type原理
将符合接口的函数定义为类型，然后对这个类型实现接口中的函数，实现的时候就直接调用自身。使用的时候只需要将自定义的函数（原型相同）做类型转换就完成了






