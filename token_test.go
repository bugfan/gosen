package gosen

import (
	"log"
	"testing"
)

func TestSen(t *testing.T) {

	log.Println("-----测试-----")
	obj := New("*")                                                           // new一个对象(传入过滤之后的token)
	obj.Fit = false                                                           // 是否适配个数 (根据需要传，一般不用设置)
	obj.Tokens = []string{"共产党", "射精", "阴茎", "茎", "阴道", "VX", "威信", "裸", "逼"} // 设置敏感词库
	log.Println(obj.Filter(`微信, 射精 asd9u按实际 !2VX`))                           // 过滤

}
