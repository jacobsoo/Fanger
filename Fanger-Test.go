package main
 
import (
	"fmt"
    "Fanger"
    "flag"
)

func main(){
    DeFangPtr := flag.String("defang", "URL", "A url to be defanged")
    ReFangPtr := flag.String("refang", "fanged URL", "A url to be refanged")
    flag.Parse()

    if *DeFangPtr != "URL" {
        fmt.Println("[+] Defanging in process.")
        Fanger.Defang(*DeFangPtr)
    }else if *ReFangPtr != "fanged URL" {
        fmt.Println("[+] Refanging in process.")
        Fanger.Refang(*ReFangPtr)
    }
}
