//Sum of sequence

package kata


func SequenceSum(start, end, step int) int {
  var result int
  if start==end{
    return start
  }else if start<end{
    for i:=start;i<=end;i=i+step{
      result+=i
    }
    return result
  }else{
  return 0
    }
}


//sort the vowels
package kata
import (
  "strings"
  )


func SortVowels(s string) string {
  var str string 
  for i:=0;i<len(s);i++{
    switch s[i]{
      case 'A','E','I','O','U','a','e','i','o','u':
      str+="|"+string(s[i])+"\n"
      default:
       str+=string(s[i])+"|\n"
    }
  }
	return strings.TrimSuffix(str,"\n")
}

//get middle character
package kata

func GetMiddle(s string) string {
  length:=len(s)
  mid:=length/2
    if length%2==0{
      return s[mid-1:mid+1]
    }else{
      return string(s[mid])
    }
  //Code goes here!
}


