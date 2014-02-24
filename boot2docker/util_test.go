package main

import (
	"testing"
)

var getenvTests = map[string]map[string]string{
    "1": map[string]string{
        "in":"", 
        "fb":"", 
        "out":"",
        "res":"",
    },
    "test": map[string]string{
        "in":"", 
        "fb":"", 
        "out":"test1",
        "res":"test1",
    },
    "test1": map[string]string{
        "in":"ab", 
        "fb":"fallback", 
        "out":"",
        "res":"fallback",
    },
    "test2": map[string]string{
        "in":"ab", 
        "fb":"", 
        "out":"test1",
        "res":"test1",
    },
}

/*var getenvTests = []getenvx {
	in,fb ,out string
}{
	{"","" ,""},
	{"test","", "test"},
	{"fallback","fallback", ""},
}
*/
func Test_getenvSys(t *testing.T) {

	for k, test := range getenvTests {
		var fakeGetenv = func(key string) string {
			return getenvTests[k]["out"]
		}

		result := getenvSys(test["in"],test["fb"],fakeGetenv)
		if result != test["res"] {t.Errorf("Test in: %v Expected %v to equal %v",k, result, test["res"])}
	}
}