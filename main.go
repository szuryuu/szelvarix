package main

import (
	"fmt"
	"strings"
	"szelvarix/greetings"
)

var name string

func main() {
	// ASCII-art Tree of Life
	fmt.Print(`
            .        +          .      .          .
     .            _        .                    .
  ,              /;-._,-.____        ,-----.__
 ((        .    (_:#::_.:::. '-._   /:, /-._, '._,
  '                 \   _|'"=:_::.' );  \ __/ /
                      ,    './  \:. '.   )==-'  .
    .      ., ,-=-.  ,\, +#./'   \:.  / /           .
.           \/:/'-' , ,\ ' + "" +    ):  \ /_  -o
       .    /:+- - + +- : :- + + -:'  /(o-) \)     .
  .      ,=':  \    ' '/ + "" +  ' , , ,:' '--".--"---._/7
   '.   (    \: \,-._' ' + '\, ,"   _,--._,---":.__/
              \:  '  X' _| _,\/'   .-'
.               ":._:'\____  /:'  /      .           .
                    \::.  :\/:'  /              +
   .                 '.:.  /:'  }      .
           .           ):_(:;   \           .
                      /:. _/ ,  |
    .                (|::.     , + "" +                   .
     .                |::.    {\                  .
                      |::.\  \ '.
                      |:::(\    |
              O       |:::/{ }  |                  (o
               )  ___/#\::'/ (O "==._____   O, (O  /'
          ~~~w/w~"~~,\ '/,-(~'"~~~~~~~~"~o~\~/~w|/~
dew   ~~~~~~~~~~~~~~~~~~~~~~~\\W~~~~~~~~~~~~\|/~~
`)

	// Epic narrative
	fmt.Println("✨ The roots of *Szelvarix* stir gently as you approach...")
	fmt.Println("🌿 A soft voice echoes from within the sacred bark:")
	fmt.Println()
	fmt.Println("    \"State your name, brave wanderer of the realm:\"")
	fmt.Print("    > ")

	fmt.Scanln(&name)

	nameToUpperCase := strings.ToUpper(name)
	greet := greetings.Hello(nameToUpperCase)
	fmt.Println()
	fmt.Println("🌳 The leaves shimmer as your essence is read...")
	fmt.Println(greet)
}
