package json

import (
	"encoding/json"
	"fmt"
)

const (
	nullstr = "null"
)

type SampleStruct struct {
	Str  string                 `json:"str"`
	Pstr *string                `json:"pstr"`
	Num  int                    `json:"num"`
	Pnum *int                   `json:"pnum"`
	Inf  interface{}            `json:"inf"`
	Imap map[string]interface{} `json:"imap"`
}

// https://go.dev/play/p/8aEgE7NK2Iu
func UnmarshalNullStr() {
	var astr string
	var pstr *string = new(string)
	var aint int
	var pint *int = new(int)
	var ainf interface{}
	infmap := make(map[string]interface{})
	rmmap := make(map[string]json.RawMessage)
	sample := &SampleStruct{}

	var err error

	err = json.Unmarshal([]byte(nullstr), &astr)
	fmt.Printf("string: '%v', addr: %v, err: %v\n", astr, &astr, err)

	err = json.Unmarshal([]byte(nullstr), pstr)
	fmt.Printf("pstring: %v, addr: %v, err: %v\n", pstr, &pstr, err)

	err = json.Unmarshal([]byte(nullstr), &aint)
	fmt.Printf("int: %v, : %v, err: %v\n", aint, &aint, err)

	err = json.Unmarshal([]byte(nullstr), pint)
	fmt.Printf("pint: %v, err: %v\n", pint, err)

	err = json.Unmarshal([]byte(nullstr), &ainf)
	fmt.Printf("interface %v, err: %v\n", ainf, err)

	err = json.Unmarshal([]byte(nullstr), &infmap)
	fmt.Printf("interface map %v, nil: %t, err: %v\n", infmap, infmap == nil, err)

	err = json.Unmarshal([]byte(nullstr), &rmmap)
	fmt.Printf("json map %v, nil: %t, err: %v\n", rmmap, rmmap == nil, err)

	err = json.Unmarshal([]byte(nullstr), sample)
	fmt.Printf("sample map %+v, sample.Imap == nil: %t, err: %v\n", sample, sample.Imap == nil, err)

}
