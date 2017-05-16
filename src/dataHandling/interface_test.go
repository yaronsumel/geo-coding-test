package dataHandling

import (
	"testing"
	"github.com/yaronsumel/geo-coding-test/src/place"
	"reflect"
)

// empty demo handler
type X struct {}
func (x *X)Next(p *place.Place)bool{return false}
func (x *X)Close()error{return nil}

// TestRegHandler
func TestRegHandler(t *testing.T) {
	demoHandler := &X{}
	RegHandler("test",demoHandler)
	if !reflect.DeepEqual(handlers["test"],demoHandler){
		t.FailNow()
	}
}

// TestGetHandler
func TestGetHandler(t *testing.T) {
	demoHandler := &X{}
	RegHandler("test",demoHandler)
	x,_:=GetHandler("test")
	if !reflect.DeepEqual(handlers["test"],x){
		t.FailNow()
	}
	//
	_,err:=GetHandler("testNotValid")
	if err==nil{
		t.FailNow()
	}
}
