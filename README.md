# go-design-model
用golang实现的常用设计模式库

## 安装
```
go get github.com/hongker/go-design-model
```
## 单例模式
- [单例模式](https://github.com/hongker/go-design-model/tree/main/singleton)
- 使用
```go
import (
    "github.com/hongker/go-design-model/singleton"
    "fmt"
)

func main() {
    instance := New(someConstructor)
    fmt.Println(instance.Get())
}
func someConstructor() interface{}{
    return "hello"
}
```