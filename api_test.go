package api

import (
	"reflect"
	"testing"
)

func TestGetServer(t *testing.T) {

	initializationPort := 9090
	reinitializationPort := 8080 //should be not working
	serverVar := GetServer(initializationPort)
	secondServerVar := GetServer(reinitializationPort)

	if secondServerVar == nil || serverVar == nil {
		t.Fatalf("Server is uninitialized. server is %v", server)
	}

	if !reflect.DeepEqual(serverVar, secondServerVar) || secondServerVar.Addr != serverVar.Addr {
		t.Errorf("server is instanitiating multiple times. First:%v & second:%v", serverVar.Addr, secondServerVar.Addr)
	}

}
