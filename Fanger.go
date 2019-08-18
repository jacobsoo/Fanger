package Fanger

import(
    "fmt"
    "strings"
)

func Defang(szURL string){
	replacer := strings.NewReplacer("http", "hxxp", ".", "[.]", "ftp", "fxp")
	
    output := replacer.Replace(szURL)
    fmt.Printf("[*] Original URL : %v\n", szURL)
	fmt.Printf("[*] Defanged URL : %v\n", output)
}

func Refang(szURL string){
    replacer := strings.NewReplacer("hxxp", "http", "[.]", ".", "\\.", ".", "[.", ".", "fxp", "ftp")
	
    output := replacer.Replace(szURL)
    fmt.Printf("[*] Original URL : %v\n", szURL)
	fmt.Printf("[*] Refanged URL : %v\n", output)
}
