package proxy

import (
	"fmt"
	"testing"
)

var TestProxy = &SSHProxy{}

func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setupTestCase")
	TestProxy = &SSHProxy{
		User: "test",
		Host: "127.0.0.1",
		Port: 22,
		Key:  "/tmp/konnect/key",
		Name: "host1",
	}
	return func(t *testing.T) {
		t.Log("teardownTestCase")
		TestProxy = &SSHProxy{}
	}
}

// func TestMain(m *testing.M) {
// 	setUp()
// 	fmt.Println("hello....")
// 	code := m.Run()
// 	tearDown()
// 	os.Exit(code)
// }

func TestString(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	expectedStr1 := "<SSHProxy: test@127.0.0.1>"
	proxyStr1 := fmt.Sprintf("%v", TestProxy)
	// Test populated SSHProxy.
	if proxyStr1 != expectedStr1 {
		t.Errorf("SSHProxy String. Expected '%v' Got '%v'", expectedStr1, proxyStr1)
	}

	// Test empty SSHProxy.
	TestProxy = &SSHProxy{}
	expectedStr2 := "<SSHProxy: empty>"
	proxyStr2 := fmt.Sprintf("%v", TestProxy)
	if proxyStr2 != expectedStr2 {
		t.Errorf("SSHProxy String. Expected '%v' Got '%v'", expectedStr2, proxyStr2)
	}
}

func TestInfo(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	expectedStr := fmt.Sprintf("[host1]\n" +
		"  User: test\n" +
		"  Host: 127.0.0.1\n" +
		"  Port: 22\n" +
		"  Key: /tmp/konnect/key\n")
	infoStr := TestProxy.Info()

	// Test SSHProxy Info.
	if infoStr != expectedStr {
		t.Errorf("SSHProxy Info. Expected '%v' Got '%v'", expectedStr, infoStr)
	}
}

func TestPrintStatus(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	expectedStr1 := fmt.Sprintf("Connection FAIL\t-> [host1]")
	infoStr1 := fmt.Sprintf("%v", TestProxy.PrintStatus())

	// Test PrintStatus Info with invalid connection.
	if infoStr1 != expectedStr1 {
		t.Errorf("SSHProxy Info. Expected '%v' Got '%v'", expectedStr1, infoStr1)
	}

	TestProxy.Connection = true
	expectedStr2 := fmt.Sprintf("Connection OK\t-> [host1]")
	infoStr2 := fmt.Sprintf("%v", TestProxy.PrintStatus())

	// Test PrintStatus Info with valid connection.
	if infoStr2 != expectedStr2 {
		t.Errorf("SSHProxy Info. Expected '%v' Got '%v'", expectedStr2, infoStr2)
	}
}

func TestArgs(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	expectedArgs := []string{
		"ssh",
		"-i",
		"/tmp/konnect/key",
		"-p",
		"22",
		"test@127.0.0.1",
	}
	proxyArgs := TestProxy.Args()

	// Test SSHProxy Args length.
	if len(proxyArgs) != len(expectedArgs) {
		// t.Errorf("SSHProxy Args. Expected %v Got %v", expectedArgs, proxyArgs)
		t.Errorf("SSHProxy Args. Expected %v Got %v", len(expectedArgs), len(proxyArgs))
	}

	// Test SSHProxy Args values.
	for i := range expectedArgs {
		if proxyArgs[i] != expectedArgs[i] {
			t.Errorf("SSHProxy Args. Expected %v Got %v", expectedArgs[i], proxyArgs[i])
		}
	}
}
