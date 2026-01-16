package processor

func Up(words []string) {
    for i, word := range words {
        r := []rune(word)       // convert string to rune slice
        for j, char := range r {
            if char >= 'a' && char <= 'z' { // only lowercase ASCII
                r[j] = char - 32  // convert to uppercase
            }
        }
        words[i] = string(r) // save modified string back into slice
    }
}

func Low(words []string){
	for i, word := range words{
		r := []rune(word)
		for j, char := range word{
			if char >= 'A' && char <= 'Z'{
				r[j] = char + 32
			}
		}
		words[i] = string(r)
	}
}

func Cap(words []string){
	for i, word := range words{
		r := []rune(word)
		r[0] = r[0] - 32
		words[i] = string(r)
	}
}