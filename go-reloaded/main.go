package main
import "fmt"
import "os"
import "io/fs"
import "strings"
import "strconv"
import "go-reloaded/processor"
//import "go-reloaded/utils"

func main ()  {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <inputfile> <outputfile>")
        return
	}
	
	inputfile := os.Args[1]
	//outputfile := os.Args[2]
	myFs := os.DirFS(".")

	data, err := fs.ReadFile(myFs, inputfile)
	if err != nil{
		fmt.Println(err)
	}
	
	//process 
	text := []rune(string(data))
	fmt.Println(len(text))
	inbracket := false
	inquotes := false
	
	var modifier string
	res := "" 
	//fmt.Println(string(text))
	for i := 0; i < len(text); {
		v := text[i]
		if v == ' '{
			if i < len(text){
				switch text[i+1]{
					case ',', '?', ';', ':', '.', ' ':
						processor.FixPunctuation(text[i:])
						text = text[:len(text)-1]
						continue
					default:
						_ = 4 + 5
				}
			}
		}
		if v == '(' {
			inbracket = true
			i++
			continue
		}
		if v == ')'{
			inbracket = false
			i++
			continue
		}
		if !inbracket && modifier != ""{
			modifv := strings.Split(modifier, ", ")

			//fmt.Println(modifv)
			nres := strings.Fields(res)
			if len(modifv) > 1{
				//fmt.Println(res)
				index, _ := strconv.Atoi(modifv[1])
				
				//fmt.Println(nres)
				switch modifv[0]{
					case "up":
						processor.Up(nres[len(nres)-index: ])
					case "low":
						processor.Low(nres[len(nres)-index: ])
					case "cap":
						processor.Cap(nres[len(nres)-index: ])
				}
			}else{
				//fmt.Println(res)
				switch modifv[0]{
					case "up":
						processor.Up(nres[len(nres)-1: ])
					case "low":
						processor.Low(nres[len(nres)-1: ])
					case "cap":
						processor.Cap(nres[len(nres)-1: ])
					case "hex":
						processor.HextoDec(nres[len(nres)-1: ])
					case "bin":
						processor.BintoDec(nres[len(nres)-1: ])
				}
				res = strings.Join(nres, " ")
			}
			modifier = ""
		}
		if inbracket {
			modifier = modifier + string(v)
			i++
			continue
		}
		//Handle punctuation
		if v == ',' || v == '.' || v == '?' || v == ';' || v == ':' {
			res += string(v)
			
			j := i + 1

			for j < len(text) && text[j] == ' '{
				j++
			}
			res += " "
			i = j
			continue
		}

		// handle quotes
		if v == '\''{
			inquotes = !inquotes
			
			continue
			
		}
		
		if inquotes{
			

		}
		res += string(v)
		i++
	}
	fmt.Println(len(text))
	fmt.Println(res)
}