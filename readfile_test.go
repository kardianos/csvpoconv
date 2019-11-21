package main

import (
	"bytes"
	"encoding/csv"
	"strings"
	"testing"
)

func TestReadFileReader(t *testing.T) {
	rf := strings.NewReader(`"Week Start Date","Week","Day","PONumber","Delivery Location","Processor","Customer Item Num","Line Sequence","Quantity (IN)","Rate (CO)"
"11/18/2019","Week 12","11/19/2019","10419P","DSP","1015","FSHINS","1","42,000.00","$2.9225"
"11/18/2019","Week 12","11/19/2019","10481P","DSP","144","FSHINS","1","42,000.00","$2.9394"
"11/18/2019","Week 12","11/19/2019","10496P","DSP","151","FSHINS","1","40,000.00","$2.9294"
"11/18/2019","Week 12","11/19/2019","10614P","HEN","1015","FZNINS","1","19,250.00","$3.0500"
"11/18/2019","Week 12","11/20/2019","10599P","HEN","515","BEEFHEARTS","1","39,000.00","$0.8100"
"11/18/2019","Week 12","11/22/2019","10448P","DSP","509","FSHINS","1","42,000.00","$2.9394"
"11/18/2019","Week 12","11/22/2019","10461P","DSP","507","FSHINS","1","42,000.00","$2.9190"
"11/18/2019","Week 12","11/22/2019","10482P","DSP","144","FSHINS","1","42,000.00","$2.9394"
"11/18/2019","Week 12","11/22/2019","10516P","HEN","1015","BEEF65","1","38,500.00","$1.0500"
"11/18/2019","Week 12","11/22/2019","10616P","HEN","212","BEEF50","1","22,470.00","$1.0900"
"11/18/2019","Week 12","11/22/2019","10617P","HEN","507","FZNINS","1","42,000.00","$3.0200"
`)
	wf := &bytes.Buffer{}
	want := `
H,10419P,1015,11/19/2019,DSP,DSP
R,FSHINS,42000,2.9225
H,10481P,144,11/19/2019,DSP,DSP
R,FSHINS,42000,2.9394
H,10496P,151,11/19/2019,DSP,DSP
R,FSHINS,40000,2.9294
H,10614P,1015,11/19/2019,HEN,HEN
R,FZNINS,19250,3.0500
H,10599P,515,11/20/2019,HEN,HEN
R,BEEFHEARTS,39000,0.8100
H,10448P,509,11/22/2019,DSP,DSP
R,FSHINS,42000,2.9394
H,10461P,507,11/22/2019,DSP,DSP
R,FSHINS,42000,2.9190
H,10482P,144,11/22/2019,DSP,DSP
R,FSHINS,42000,2.9394
H,10516P,1015,11/22/2019,HEN,HEN
R,BEEF65,38500,1.0500
H,10616P,212,11/22/2019,HEN,HEN
R,BEEF50,22470,1.0900
H,10617P,507,11/22/2019,HEN,HEN
R,FZNINS,42000,3.0200
`

	w := csv.NewWriter(wf)
	err := readFileReader(w, rf)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.TrimSpace(wf.String())
	want = strings.TrimSpace(want)
	if want != got {
		t.Fatalf("incorrect result, got:\n%s\n", got)
	}
}
