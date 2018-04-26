
package gosen
import(
	"testing"	
	"log"
)
func TestSen(t *testing.T) {
	// log.Println("-----加载文件方式-----")
	// LoadTokens()
	// log.Println(DealSensitiveString(`微信,射精asd9u按实际`))

	log.Println("-----测试-----")
	obj:=New("*-")
	obj.Tokens=[]string{"共产党","射精","阴茎","茎","阴道","VX","威信","裸","逼"}
	log.Println(obj.Screening(`微信, 射精 asd9u按实际 !2VX`))
	
	
}
