

package main

// brainfuck
import ("fmt"
         "os")

var p, pc int
var a [30000]byte

//const prog = "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.!"



func scan(dir int, prog string) {
	for nest := dir; dir*nest > 0; pc += dir {
		switch prog[pc+dir] {
		case ']':
			nest--
		case '[':
			nest++
		}
	}
}

func main() {
  checkArgs()
var prog string = os.Args[1]
	r := ""
	for {
		switch prog[pc] {
		case '>':
			p++
		case '<':
			p--
		case '+':
			a[p]++
		case '-':
			a[p]--
		case '.':
			r += string(a[p])
		case '[':
			if a[p] == 0 {
				scan(1,prog)
			}
		case ']':
			if a[p] != 0 {
				scan(-1,prog)
			}
		default:
			fmt.Println("working..")
			fmt.Println(r)
			return
		}
		pc++
	}
}
func checkArgs(){
  if len(os.Args) != 2 {
    os.Exit(0)
  }
}
