## 单例模式
无论获取多少次实例，保证实例的初始化只会调用一次。通过`sync.Once`实现，
## 使用
```go
package main
import (
    "github.com/hongker/go-design-model/singleton"
    "fmt"
)

func main() {
    single := singleton.New(someConstructor)
    // 执行多次，也只会初始化一次
    for i := 0; i<10;i++ {
        fmt.Println(single.Instance())
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