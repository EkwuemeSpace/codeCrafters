package main
import "fmt"
import "bufio"
import "os"
import "strings"


func formatter(){
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("WELCOME! THIS IS A STRING FORMATTER CLI")
	fmt.Println()
	initial:
	fmt.Println("Our Operators are: \n-lower <text>\n-cap <text>\n-title <text>\n-snake <text>\n-reverse <text>\n-exit ")
	fmt.Println()
	fmt.Println("choose an operator!")
	scanner.Scan()
	input:=scanner.Text()
	input=strings.TrimSpace(input)
	parts:=strings.SplitN(input, " ", 2)
	operation:=parts[0]
	var text string
	if len(parts)==2{
		text=parts[1]
	}
	if len(parts)!=2{
		fmt.Println("✗ No text provided. Usage: "+operation+ " <text>")
		fmt.Println()
		goto initial
		
	}
	
	switch operation{
	case "cap":
		fmt.Println(ToUpperCase(text))
	case "lower":
		fmt.Println(ToLowerCase(text))
	case "title":
		fmt.Println(capsFirstLetter(text))
	case "snake":
		fmt.Println(snakeCase(text))
	case "reverse":
		fmt.Println(reverseSentence(text))
	case "exit":
		fmt.Println("shutting down string formatter!")
		return
	default:
		fmt.Println(`✗ Unknown command: `+fmt.Sprintf("%q",operation))


	}
	

}

func main(){
	formatter()
}
