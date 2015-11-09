// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/fixtures"
	"github.com/maxbrunsfeld/counterfeiter/fixtures/another_package"
)

type FakeEmbedsInterfaces struct {
	ServeHTTPStub        func(http.ResponseWriter, *http.Request)
	serveHTTPMutex       sync.RWMutex
	serveHTTPArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	AnotherMethodStub        func([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType)
	anotherMethodMutex       sync.RWMutex
	anotherMethodArgsForCall []struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}
	EmbeddedMethodStub        func() string
	embeddedMethodMutex       sync.RWMutex
	embeddedMethodArgsForCall []struct{}
	embeddedMethodReturns struct {
		result1 string
	}
	DoThingsStub        func()
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct{}
}

func (fake *FakeEmbedsInterfaces) ServeHTTP(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.serveHTTPMutex.Lock()
	fake.serveHTTPArgsForCall = append(fake.serveHTTPArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.serveHTTPMutex.Unlock()
	if fake.ServeHTTPStub != nil {
		fake.ServeHTTPStub(arg1, arg2)
	}
}

func (fake *FakeEmbedsInterfaces) ServeHTTPCallCount() int {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return len(fake.serveHTTPArgsForCall)
}

func (fake *FakeEmbedsInterfaces) ServeHTTPArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.serveHTTPMutex.RLock()
	defer fake.serveHTTPMutex.RUnlock()
	return fake.serveHTTPArgsForCall[i].arg1, fake.serveHTTPArgsForCall[i].arg2
}

func (fake *FakeEmbedsInterfaces) AnotherMethod(arg1 []another_package.SomeType, arg2 map[another_package.SomeType]another_package.SomeType, arg3 *another_package.SomeType, arg4 another_package.SomeType, arg5 chan another_package.SomeType) {
	fake.anotherMethodMutex.Lock()
	fake.anotherMethodArgsForCall = append(fake.anotherMethodArgsForCall, struct {
		arg1 []another_package.SomeType
		arg2 map[another_package.SomeType]another_package.SomeType
		arg3 *another_package.SomeType
		arg4 another_package.SomeType
		arg5 chan another_package.SomeType
	}{arg1, arg2, arg3, arg4, arg5})
	fake.anotherMethodMutex.Unlock()
	if fake.AnotherMethodStub != nil {
		fake.AnotherMethodStub(arg1, arg2, arg3, arg4, arg5)
	}
}

func (fake *FakeEmbedsInterfaces) AnotherMethodCallCount() int {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	return len(fake.anotherMethodArgsForCall)
}

func (fake *FakeEmbedsInterfaces) AnotherMethodArgsForCall(i int) ([]another_package.SomeType, map[another_package.SomeType]another_package.SomeType, *another_package.SomeType, another_package.SomeType, chan another_package.SomeType) {
	fake.anotherMethodMutex.RLock()
	defer fake.anotherMethodMutex.RUnlock()
	return fake.anotherMethodArgsForCall[i].arg1, fake.anotherMethodArgsForCall[i].arg2, fake.anotherMethodArgsForCall[i].arg3, fake.anotherMethodArgsForCall[i].arg4, fake.anotherMethodArgsForCall[i].arg5
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethod() string {
	fake.embeddedMethodMutex.Lock()
	fake.embeddedMethodArgsForCall = append(fake.embeddedMethodArgsForCall, struct{}{})
	fake.embeddedMethodMutex.Unlock()
	if fake.EmbeddedMethodStub != nil {
		return fake.EmbeddedMethodStub()
	} else {
		return fake.embeddedMethodReturns.result1
	}
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodCallCount() int {
	fake.embeddedMethodMutex.RLock()
	defer fake.embeddedMethodMutex.RUnlock()
	return len(fake.embeddedMethodArgsForCall)
}

func (fake *FakeEmbedsInterfaces) EmbeddedMethodReturns(result1 string) {
	fake.EmbeddedMethodStub = nil
	fake.embeddedMethodReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeEmbedsInterfaces) DoThings() {
	fake.doThingsMutex.Lock()
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct{}{})
	fake.doThingsMutex.Unlock()
	if fake.DoThingsStub != nil {
		fake.DoThingsStub()
	}
}

func (fake *FakeEmbedsInterfaces) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

var _ fixtures.EmbedsInterfaces = new(FakeEmbedsInterfaces)