package proxy

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	expectedStr1 := "<SSHProxy: test@127.0.0.1>"
	proxyStr1 := fmt.Sprintf("%v", defaultProxy())

	if proxyStr1 != expectedStr1 {
		t.Errorf("Expected '%v' Got '%v'", expectedStr1, proxyStr1)
	}

	expectedStr2 := "<SSHProxy: empty>"
	proxyStr2 := fmt.Sprintf("%v", emptyProxy())
	if proxyStr2 != expectedStr2 {
		t.Errorf("Expected '%v' Got '%v'", expectedStr2, proxyStr2)
	}
}

func TestInfo(t *testing.T) {
	proxy := defaultProxy()
	expectedStr := fmt.Sprintf("%v@%v\n", proxy.User, proxy.Host)
	infoStr := proxy.Info()

	if infoStr != expectedStr {
		t.Errorf("Expected '%v' Got '%v'", expectedStr, infoStr)
	}
}

func TestPrintStatus(t *testing.T) {
	proxy := defaultProxy()

	expectedStr1 := fmt.Sprintf("Connection FAIL\t-> [host1]")
	infoStr1 := fmt.Sprintf("%v", proxy.PrintStatus())

	if infoStr1 != expectedStr1 {
		t.Errorf("Expected '%v' Got '%v'", expectedStr1, infoStr1)
	}

	proxy.Connection = true
	expectedStr2 := fmt.Sprintf("Connection OK\t-> [host1]")
	infoStr2 := fmt.Sprintf("%v", proxy.PrintStatus())

	if infoStr2 != expectedStr2 {
		t.Errorf("Expected '%v' Got '%v'", expectedStr2, infoStr2)
	}
}

func TestArgs(t *testing.T) {
	expectedArgs := []string{
		"ssh",
		"-i",
		"/tmp/konnect/key",
		"-p",
		"22",
		"test@127.0.0.1",
	}
	proxyArgs := defaultProxy().Args()

	if len(proxyArgs) != len(expectedArgs) {
		t.Errorf("Expected %v Got %v", len(expectedArgs), len(proxyArgs))
	}

	// Test SSHProxy Args values.
	for i := range expectedArgs {
		if proxyArgs[i] != expectedArgs[i] {
			t.Errorf("Expected %v Got %v", expectedArgs[i], proxyArgs[i])
		}
	}
}

func defaultProxy() *SSHProxy {
	return newProxy("host1", "127.0.0.1", "test", 22, "/tmp/konnect/key")
}

func emptyProxy() *SSHProxy {
	return &SSHProxy{}
}

func newProxy(name, host, user string, port int, key string) *SSHProxy {
	return &SSHProxy{
		Host: host,
		User: user,
		Port: port,
		Key:  key,
		Name: name,
	}
}
