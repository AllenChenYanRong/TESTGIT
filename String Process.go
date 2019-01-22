package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {
	str := "My name is "
	str1 := "aello"
	c := []byte(str1)
	c[0] = 'h'
	str2 := string(c)
	fmt.Println("字串合併")
	fmt.Println(str + "Evan") // My name is Evan
	fmt.Println("字串改字元")
	fmt.Println(str2) // hello
	fmt.Println("字串按字典順序比較結果為-1,0,1")
	fmt.Println(strings.Compare("a", "c")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("c", "a")) // 1
	fmt.Println("看前面字串有沒有包含後面字串")
	fmt.Println(strings.Contains("hello world", "llo w")) // true
	fmt.Println(strings.Contains("hello world", "321"))   // false
	fmt.Println(strings.Contains("hello world", "wor"))   // true
	fmt.Println("看前面字串有沒有包含後面字串的任意字元")
	fmt.Println(strings.ContainsAny("team", "i"))      // false
	fmt.Println(strings.ContainsAny("failure", "u i")) // true
	fmt.Println(strings.ContainsAny("failure", "a b")) // true
	fmt.Println("看字串有沒有包含後面的ASCII或rune類型的r字元")
	fmt.Println(strings.ContainsRune("abc", 97))        // true
	fmt.Println(strings.ContainsRune("b", 98))          // true
	fmt.Println(strings.ContainsRune("abc", rune('b'))) // true
	fmt.Println("看後面字串在前一字串有重複幾次")
	fmt.Println(strings.Count("Evannnn", "n"))      // 4
	fmt.Println(strings.Count("Evanvanvan", "van")) // 3
	fmt.Println("查看前後字串不分大小寫之後是否相等")
	fmt.Println(strings.EqualFold("Evan", "evan"))               // true
	fmt.Println(strings.EqualFold("Evan", "evann"))              // false
	fmt.Println(strings.EqualFold("Hello world", "hello World")) // true
	fmt.Println("使用space切字串token")
	fmt.Printf("Fields are: %q", strings.Fields("  abc1  evan  cam   ")) // Fields are: ["abc" "evan" "cam"]
	fmt.Println("使用function判斷哪些字元要切成token")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	} // Fields are: ["abc1" "evan2" "cam3" "xyz" "z2" "123"]
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  abc1@evan2,cam3*xyz%z2^123...", f))
	fmt.Println("判斷是否為後面字串當開頭")
	fmt.Println(strings.HasPrefix("Evan", "Ev"))  // true
	fmt.Println(strings.HasPrefix("Evan", "van")) // false
	fmt.Println(strings.HasPrefix("Evan", ""))    // true
	fmt.Println("判斷是否為後面字串當結尾")
	fmt.Println(strings.HasSuffix("Evan", "an")) // true
	fmt.Println(strings.HasSuffix("Evan", "Ev")) // false
	fmt.Println(strings.HasSuffix("Evan", ""))   // true
	fmt.Println("返回substring在字串中最前面的索引，不存在返回-1")
	fmt.Println(strings.Index("Evanevan", "an")) // 2
	fmt.Println(strings.Index("Evan", "don"))    // -1
	fmt.Println("返回後面字串任一字元在前字串中最前面的索引，皆不存在返回-1")
	fmt.Println(strings.IndexAny("Evan", "orangE")) // 0
	fmt.Println(strings.IndexAny("Evan", "orange")) // 2
	fmt.Println(strings.IndexAny("Evan", "pig"))    // -1
	fmt.Println("同上不過只判斷一個字元")
	fmt.Println(strings.IndexByte("Evan", 'e')) // -1
	fmt.Println(strings.IndexByte("Evan", 'a')) // 2
	fmt.Println(strings.IndexByte("Evan", 'A')) // -1
	fmt.Println("使用function抓取字串裡中文最前面的索引，無則返回-1")
	f1 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f1))   // 7
	fmt.Println(strings.IndexFunc("哈囉,世界", f1))       // 0
	fmt.Println(strings.IndexFunc("Hello,world", f1)) // -1
	fmt.Println("在各個字串中間加入想要的分隔標準")
	str3 := []string{"abc", "evan", "don"}
	fmt.Println(strings.Join(str3, ", ")) // abc, evan, don
	fmt.Println(strings.Join(str3, "||")) // abc||evan||don
	fmt.Println("後面字串在前字串排在最後的索引")
	fmt.Println(strings.LastIndex("evan evan evan", "an")) // 12
	fmt.Println(strings.LastIndex("evan evan", "ban"))     // -1
	fmt.Println("後面字串任一字元在前字串排在最後的索引")
	fmt.Println(strings.LastIndexAny("evan evan evan", "aban")) // 13
	fmt.Println(strings.LastIndexAny("evan evan", "b"))         // -1
	fmt.Println("後面字元在前字串排在最後的索引")
	fmt.Println(strings.LastIndexByte("evan evan evan", 'a')) // 12
	fmt.Println(strings.LastIndexByte("evan evan", 'b'))      // -1
	fmt.Println("數字字元在字串中最後面的索引")
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) // 5
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber)) // 2
	fmt.Println("重複印出字串")
	fmt.Println(strings.Repeat("evan ", 3)) // evan evan evan
	fmt.Println("替換字串中的substring，最後為記數，小於0表示全部替換")
	fmt.Println(strings.Replace("evann evann evann", "nn", "n", 2))     // evan evan evann
	fmt.Println(strings.Replace("evannn evannn evannn", "nn", "n", -1)) // evann evann evann
	fmt.Println("使用split函數後面的字串切token")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Println("使用split函數後面的字串切token，位於token後面")
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ",")) // ["a," "b," "c"]
	fmt.Println("使用split函數後面的字串切token看要切成幾分")
	fmt.Printf("%q\n", strings.SplitN("a,b,c,d", ",", 3)) // ["a" "b" "c,d"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
	fmt.Println("每個token字首自動換成大寫")
	fmt.Println(strings.Title("my localhost web")) // Her Royal Highness
	fmt.Println("每個token自動全部換成大寫或小寫(用於德語、芬蘭語...等國家)")
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş")) // önnek iş
	fmt.Println("每個token全部自動換成大寫或小寫")
	fmt.Println(strings.ToTitle("хлеб"))             // ХЛЕБ
	fmt.Println(strings.ToLower("My LoCalhost Web")) // my localhost web
	// ToUpper全部自動換成大寫
	fmt.Println("字串前後去掉後面字串包含的文字")
	fmt.Println(strings.Trim("¡¡¡Hello, !!!Evan!!!", "!¡")) // Hello, !!!Evan

	fmt.Println("時間字串")
	currentTime := time.Now()
	fmt.Println(currentTime)
	y := time.Now().Year()
	m := time.Now().Month()
	d := time.Now().Day()
	h := time.Now().Hour()
	min := time.Now().Minute()
	s := time.Now().Second()
	fmt.Printf("%d-%d-%d %d-%d-%d\n", y, m, d, h, min, s)
	fmt.Println(y, m, d, h, min, s) //123
}
