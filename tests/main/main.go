package main

import (
    "flag"
    "fmt"
    "log"
    "ripple/tests"
)

func main() {
    // Define a flag to select the test
    testName := flag.String("test", "TestTrustlineUpdate", "Name of the test to run")
    flag.Parse()

    // Initialize the test account (common setup)
    if err := tests.SetupAccount("testuser", "peeruser", "127.0.0.1", "mysecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up account: %v", err)
    }

    // Map of test functions
    tests := map[string]func(){
        "TestTrustlineUpdate": tests.TestTrustlineUpdate,
        // Add more test functions here
    }

    // Run the selected test
    if testFunc, exists := tests[*testName]; exists {
        fmt.Printf("Running test: %s\n", *testName)
        testFunc()
    } else {
        log.Fatalf("Unknown test: %s", *testName)
    }
}
