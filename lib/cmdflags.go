package cbdozer

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type RequestFlags struct {
	Args        []string
	RequestType string
	Rate        uint64
	RunTime     int64
	Username    string
	Password    string
	Method      string
	URL         string
	RequestBody string
	FTSFlags    *FTSRequestFlags
}

type FTSRequestFlags struct {
	FTSQueryType  string
	FTSQueryStr   string
	FTSResultSize uint64
}

// parse top-level args and set test flag parsing mode
func NewRequestFlags() RequestFlags {

	f := RequestFlags{
		Args: os.Args[1:],
	}

	// detect mode
	if len(f.Args) > 0 {
		if strings.Index(f.Args[0], "-") != 0 {
			// is not a flag, thus mode
			f.RequestType = f.Args[0]
		} else {
			f.RequestType = "default"
		}
	}

	return f
}

// set flag vals for subcommands
func (f *RequestFlags) Parse() {

	flagSet := flag.NewFlagSet("default", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Println("Usage: cbdozer <request> [request flags]")
		fmt.Printf("\n flags for %s:\n", f.RequestType)
		flagSet.PrintDefaults()
	}

	flagSet.Uint64Var(&f.Rate, "rate", 1,
		"request rate")
	flagSet.Int64Var(&f.RunTime, "duration", 10,
		"request duration (in seconds)")
	flagSet.StringVar(&f.URL, "url", "http://undefined",
		"url to send request")
	flagSet.StringVar(&f.Method, "method", "GET",
		"request method (GET | POST ...)")

	switch f.RequestType {
	case "fts":
		// add fts flags
		ftsFlags := FTSRequestFlags{}
		flagSet.StringVar(&ftsFlags.FTSQueryType, "type", "query", "type of query")
		flagSet.StringVar(&ftsFlags.FTSQueryStr, "query", "", "query string")
		flagSet.Uint64Var(&ftsFlags.FTSResultSize, "size", 1000, "response size")
		f.FTSFlags = &ftsFlags
		flagSet.Parse(f.Args[1:])
	default:
		flagSet.StringVar(&f.RequestBody, "body", "",
			"data to send as body of request")
		flagSet.Parse(f.Args)
	}
}

func (f *RequestFlags) GetBody() []byte {
	var body []byte

	switch f.RequestType {
	case "fts":
		ftsQuery := NewFTSQuery(f.FTSFlags)
		body = ftsQuery.Body()
	default:
		body = []byte(f.RequestBody)
	}
	return body
}
