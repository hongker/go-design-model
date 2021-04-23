## 使用
```go
import (
    "github.com/hongker/go-design-model/singleton"
    "fmt"
)

func main() {
    instance := New(someConstructor)
    // 执行多次，也只会初始化一次
    for i := 0; i<10;i++ {
        fmt.Println(instance.Get())
    }   
    
}
// someConstructor 某个对象的构造函数
func someConstructor() interface{}{
	fmt.Println("init")
    return "hello"
}
```

output:
```
init
hello
hello
hello
hello
hello
hello
hello
hello
hello
hello
```