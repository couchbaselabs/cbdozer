package cbdozer

import (
	"encoding/json"
	"fmt"
)

type FTSQuery struct {
	Fields []string
	Query  map[string]string
	Size   uint64
	body   string
}

func NewFTSQuery(flags *FTSRequestFlags) *FTSQuery {
	ftsQuery := FTSQuery{
		Fields: []string{"*"},
		Size:   flags.FTSResultSize,
	}
	ftsQuery.Query = make(map[string]string)
	ftsQuery.Query[flags.FTSQueryType] = flags.FTSQueryStr

	queryJson, err := json.Marshal(ftsQuery)
	if err != nil {
		fmt.Println("all bad", err)
	}

	ftsQuery.body = string(queryJson)
	return &ftsQuery
}

func (fq *FTSQuery) Body() []byte {
	return []byte(fq.body)
}
