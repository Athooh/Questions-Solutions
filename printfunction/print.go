package printfunction

import (
	"github.com/01-edu/z01"
)

func Print(s any) {
if data, ok :=s.(string); ok{
	foo(data)
}else if n, ok:=s.(int);ok{
	if n==0{
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	s := ""
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n = n / 10

	}
	foo(s)
}
}

// func Printn(n int) {
// 	if n==0{
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	}
// 	if n < 0 {
// 		z01.PrintRune('-')
// 		n *= -1
// 	}
// 	s := ""
// 	for n > 0 {
// 		s = string(rune(n%10+'0')) + s
// 		n = n / 10

// 	}
// 	print(s)
// }

func foo(s string){
	for _, c := range s {
		z01.PrintRune(rune(c))
	}
	z01.PrintRune('\n')
}
// func main (){
// 	Print(20)
// }