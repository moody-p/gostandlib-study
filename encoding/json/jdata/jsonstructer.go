package jdata

type Person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby,omitempty"`
}

type TagExample struct {
	Name         string `json:"name"`
	Ignored      int    `json:"-"`
	EmptyIgnored int    `json:"emptyIgnored,omitempty"`
	Field        int    `json:",omitempty"`
}
