## 反射的实现原理

reflect.Type 是一个接口

```go
type Type interface {
        Align() int
        FieldAlign() int
        Method(int) Method
        MethodByName(string) (Method, bool)
        NumMethod() int
        ...
        Implements(u Type) bool
        ...
}
```

reflect.Value 是一个结构体

```go
type Value struct {
        // 包含过滤的或者未导出的字段
}

```

反射包里的方法基本上就是围绕着这 reflect.Type 和 reflect.Value 来设计的
reflect.ValueOf()和 reflect.TypeOf()
这两个函数能够获取变量的类型和变量的值
而这两个函数的入参都是 interface{}类型
interface{}在语言内部通过反射包的 reflect.emptyInterface 结构体表示

```go
type emptyInterface struct{
typ  *rtype  // 用来表示变量的类型
word unsafe.Pointer // 用来指向变量内部封装的数据
}
```

reflect.TypeOf()会将传入的类型隐式转换成 reflect.emptyInterface 类型，并且获取其中存储的类型信息
具体是怎么拿到的呢？

```go
func TypeOf(i interface{}) Type {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	return toType(eface.typ)
}

func toType(t *rtype) Type {
	if t == nil {
		return nil
	}
	return t
}
```

也就是说，TypeOf()将传递进来的空接口转换成 emptyInterface,emptyInterface 是一个结构体，其中的字段 rtype 代表变脸的类型信息,也就是 reflect.rtype

拿到这个 reflect.rtype 之后，可以使用 reflect.rtype.String()方法，可以获取类型的名字
这个 rtype 实现了 reflect.Type

那么反射是如何获取变量的值的呢？

reflect.ValueOf()

```go
func ValueOf(i interface{}) Value {
	if i == nil {
		return Value{}
	}

	escapes(i)

	return unpackEface(i)
}

func unpackEface(i interface{}) Value {
	e := (*emptyInterface)(unsafe.Pointer(&i))
	t := e.typ
	if t == nil {
		return Value{}
	}
	f := flag(t.Kind())
	if ifaceIndir(t) {
		f |= flagIndir
	}
	return Value{t, e.word, f}
}
```

简单说，当调用了 refelct.ValueOf()，它会先判断是不是 nil
然后调用 escapes()方法

这个方法会使得变量的值逃逸到堆上，然后通过 reflect.unpackEface
这个方法也会首先将 interface{}转换成 reflect.emptyface
然后将具体的类型和指针包装成 reflect.Value 结构返回

获取值的操作都是基于这个结构体的一些方法
