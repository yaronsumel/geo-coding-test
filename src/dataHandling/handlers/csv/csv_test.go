package csv

import (
	"testing"
	"github.com/yaronsumel/geo-coding-test/src/place"
)

// TestNewCsvHandler
func TestNewCsvHandler(t *testing.T) {
	// should fail
	_,err := NewCsvHandler("shouldfail.txt")
	if err==nil{
		t.FailNow()
	}
	h,err2 := NewCsvHandler("testFile.csv")
	if err2!=nil {
		t.FailNow()
	}
	if h.Close()!=nil{
		t.FailNow()
	}

}

// TestNext
func TestNext(t *testing.T)  {
	h,err := NewCsvHandler("testFile.csv")
	if err!=nil {
		t.FailNow()
	}
	defer h.Close()
	p := place.Place{}
	if h.Next(&p)==false{
		t.FailNow()
	}
	// first line
	// 382582,37.1768672,-3.608897
	if p.Id != 382582 || p.Lat != 37.1768672 || p.Lon !=-3.608897{
		t.FailNow()
	}
	// keep nexting till eof
	for h.Next(&p){
	}
	// eof
	if h.Next(&p)==true{
		t.FailNow()
	}
}

// TestParseLine
func TestParseLine(t *testing.T) {
	h := Handler{}
	p := place.Place{}
	if h.parseLine([]string{"123","12","12"},&p)!=nil{
		t.FailNow()
	}
	if h.parseLine([]string{"aaaaaaaaa","12","12aaaaaaaaaa"},&p)==nil{
		t.FailNow()
	}
	if h.parseLine([]string{"1","12asdfafs","12aaaaaaaaaa"},&p)==nil{
		t.FailNow()
	}
	if h.parseLine([]string{"1","12","12aaaaaaaaaa"},&p)==nil{
		t.FailNow()
	}
}
