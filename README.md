# gosen
  go敏感词处理代码（使用正则匹配），将字符串中的敏感词替换为自定义的token，可以过滤任意字符
## 使用方法
  - 可以参照token_test.go

  - 1. 执行 go get github.com/bugfan/gosen
  - 2. 使用o:=gosen.New("*") 新建一个对象，传入替换的token
  - 3. o.Tokens=xxx 设置需要过滤的token数组 
  - 4. o.Fit=true/false 设置是否根据敏感词的个数返回对应个数的token
  - 5. o.Filter(string) 调用此函数过滤字符串并返回替换后的结果
  
## 案例
```
package gosen
import(
	"testing"	
	"log"
)
func TestSen(t *testing.T) {
	
	log.Println("-----测试-----")
	obj:=New("*")																// new一个对象(传入过滤之后的token)
	obj.Fit=false																// 是否适配个数
	obj.Tokens=[]string{"共产党","射精","阴茎","茎","阴道","VX","威信","裸","逼"}	// 设置敏感词库 
	log.Println(obj.Filter(`微信, 射精 asd9u按实际 !2VX`))					  // 过滤
	
}
```

```
2018/04/26 16:51:37 -----测试-----
2018/04/26 16:51:37 微信, * asd9u按实际 !2* <nil>
PASS
ok      gosen   0.013s
```
