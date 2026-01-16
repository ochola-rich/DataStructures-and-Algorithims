package processor

func FixPunctuation(texts []rune){
	copy(texts[0:], texts[1:])
}

// func FixSpacing(texts [] rune){
// 	texts = texts[:len(texts) + 1]
// 	for i := 0; i < len(texts) - 1; i ++{
// 		if texts[i] == ',' && texts[i+1] != ' ' {
// 			copy(texts[i+1:], texts[i:])
// 			texts[i] = ' '
// 		}
		
// 	}
// }