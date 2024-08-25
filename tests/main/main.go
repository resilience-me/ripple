package main

import (
    "fmt"
    "log"
    "ripple/tests"
)

func main() {
    // Common setup for all tests
    if err := tests.SetupAccount("testuser", "mysecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up account: %v", err)
    }

    // Slice of test functions with their names
    testFuncs := []struct {
        name string
        fn   func()
    }{
        {"TestTrustlineUpdate", tests.TestTrustlineUpdate},
        {"TestNewPayments", tests.TestNewPayments},
        {"TestStartPayment", tests.TestStartPayment}, // Add the new test here
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
    if choice >= 0 && choice < len(testFuncs) {
        test := testFuncs[choice]
        fmt.Printf("Running test: %s\n", test.name)
        test.fn()
    } else {
        log.Fatalf("Invalid selection: %d", choice)
    }
}
