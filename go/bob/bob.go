// Package bob has method to bob's replies
package bob

import(
	"strings"
//	"unicode"
)

// Hey returns Bob's remark
func Hey(remark string) string {
	
	remark = strings.TrimSpace(remark)
	if remark ==""{
		return "Fine. Be that way!" 
	}
	yell := isYelling(remark)
	question := isQuestion(remark)

	if(question){
		if(yell){
			return  "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	
	if(yell){
		return "Whoa, chill out!"
	}
	
	return "Whatever."
}

func isUpper(s string) bool {
	alphaFlag := false
	
    	for _, r := range s {
	
        	if r <= 'z' && r >= 'a'{
            		return false
        	} else if r <= 'Z' && r >= 'A'{
			alphaFlag = true
		}
    	}
	if alphaFlag {
    		return true
	}
	return false
}

func isYelling(remark string) bool{
	if isUpper(remark){ 
		return true;
	}
	return false;
}

func isQuestion(remark string) bool{
	if(remark[len(remark) -1 ] == '?'){
		return true;
	}
	return false;
}
