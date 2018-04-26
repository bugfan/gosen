package gosen

import (
	. "fmt"
	"regexp"
)
// 传入过滤之后的token
func New(t string)Sen{
	return Sen{Token:t}
}
type Sen struct{
	Tokens 	[]string
	Token 	string
	Fit 	bool		//是否根据敏感词个数返回对应个数的过滤后的token 默认为返回单个过滤后的token
}
// 屏蔽
func (s * Sen) Screening (str string)(string,error){
	//Println("敏感词过滤: source{%s} tokens{%+v}", str, tokens)
	//处理汉字
	for j := 0; j < len(s.Tokens); j++ {
		//`[阴( )+|(\\t)+道]`
		//Printf("NO.%d %s (%d)\n", j, tokens[j], len(tokens[j]))
		reg2 := "("
		kl := len(s.Tokens[j])
		if kl > 3 {
			for k := 0; k < kl; k += 3 {
				reg2 += Sprintf(`%s`, s.Tokens[j][k:k+3])
				reg2 += " *\\t*"
			}
			reg2 = reg2[0 : len(reg2)-5] //去掉末尾规则
		} else {
			if s.Tokens[j][0:1] <= "z" {
				//处理字母
				wl := len(s.Tokens[j])
				for q := 0; q < wl; q++ {
					reg2 += Sprintf(`%s`, s.Tokens[j][q:q+1])
					reg2 += " *\\t*"
				}
				reg2 = reg2[0 : len(reg2)-5]
			} else {
				//处理一个汉
				reg2 += Sprintf(` *\t*%s *\t*`, s.Tokens[j][0:3])
			}
		}
		reg2 = reg2 + ")"
		reg := regexp.MustCompile(reg2)
		str = reg.ReplaceAllString(str,s.get8(len(s.Tokens[j])))
	}
	//Println("敏感词过滤结果: %s", str)
	return str, nil
}
// 根据i递归取i个token
func (s * Sen) get8 (i int) string {
	if !s.Fit{
		return s.Token
	}
	s8 := s.Token
	if i > 1 {
		s8 = s.get8(i-1) + s.Token
		return s8
	}
	return s8
}
