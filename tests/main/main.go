package main

import (
    "fmt"
    "log"
    "ripple/tests"
)

func main() {
    // Initialize the test account (common setup)
    if err := tests.SetupAccount("testuser", "peeruser", "127.0.0.1", "mysecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up account: %v", err)
    }

    // Map of test functions
    testFuncs := map[int]struct {
        name string
        fn   func()
    }{
        0: {"TestTrustlineUpdate", tests.TestTrustlineUpdate},
        1: {"TestSenderAndReceiverInitiatePayment", tests.TestSenderAndReceiverInitiatePayment},
        // Add more test functions here
    }

    // Display the list of tests
    fmt.Println("Available tests:")
    for i, test := range testFuncs {
        fmt.Printf("%d: %s\n", i, test.name)
    }

    // Prompt the user to select a test
    var choice int
    fmt.Print("Enter the number of the test to run: ")
    _, err := fmt.Scan(&choice)
    if err != nil {
        log.Fatalf("Failed to read input: %v", err)
    }

    // Run the selected test
    if test, exists := testFuncs[choice]; exists {
        fmt.Printf("Running test: %s\n", test.name)
        test.fn()
    } else {
        log.Fatalf("Invalid selection: %d", choice)
    }
}
