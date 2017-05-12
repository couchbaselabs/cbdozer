package main

import (
	"fmt"
	C "github.com/tahmmee/cbdozer/lib"
	"github.com/tsenart/vegeta/lib"
	"time"
)

type Auth struct {
	Username string
	Password string
}

func MakeTargeter(flags C.RequestFlags) vegeta.Targeter {
	// unpack request flags to create a new targeter
	body := flags.GetBody()

	tgt := vegeta.Target{
		Method: flags.Method,
		URL:    flags.URL,
		Body:   body,
	}
	return vegeta.NewStaticTargeter(tgt)
}

func main() {

	// parse cli flags
	flags := C.NewRequestFlags()
	flags.Parse()

	// set up targeter
	tr := MakeTargeter(flags)
	rate := flags.Rate
	runDuration := time.Duration(flags.RunTime) * time.Second

	// run attacker
	attacker := vegeta.NewAttacker()
	ch := attacker.Attack(tr, rate, runDuration)
	for result := range ch {
		fmt.Println(result)
	}
}
