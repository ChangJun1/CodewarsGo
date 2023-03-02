package Fundamental

/*
Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/

func Disemvowel(comment string) string {
	commentBytes := []byte(comment)
	var res []byte
	j := 0
	for i := 0; i < len(commentBytes); i++ {
		switch commentBytes[i] {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			fallthrough
		case 'A':
			fallthrough
		case 'E':
			fallthrough
		case 'I':
			fallthrough
		case 'O':
			fallthrough
		case 'U':
			res = append(res, commentBytes[j:i]...)
			j = i + 1
		}
	}
	res = append(res, commentBytes[j:]...)
	return string(res)
}

// most popular
//func Disemvowel(comment string) string {
//	for _, c := range "aiueoAIUEO" {
//		comment = strings.ReplaceAll(comment, string(c), "")
//	}
//	return comment
//}
