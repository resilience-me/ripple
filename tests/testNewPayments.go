package tests

import (
    "encoding/binary"
    "fmt"
    "log"
    "ripple/commands"
    "ripple/config"
)

func TestNewPayments() {
    senderUsername := "testsender"
    receiverUsername := "testreceiver"
    senderServerAddress := "127.0.0.1"
    receiverServerAddress := "127.0.0.1"
    paymentAmount := uint32(500)
    senderCounter := uint32(1)
    receiverCounter := uint32(1)

    // Setup accounts with initial balances
    if err := SetupAccount(senderUsername, receiverUsername, senderServerAddress, "sendersecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up sender account: %v", err)
    }
    if err := SetupAccount(receiverUsername, senderUsername, receiverServerAddress, "receiversecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up receiver account: %v", err)
    }

    // Hardcoded server address using the config port
    serverAddress := fmt.Sprintf("127.0.0.1:%d", config.Port)
    log.Printf("Using server address: %s", serverAddress) // Debugging line

    // Initiate the payment process
    if err := initiatePayment(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, serverAddress, paymentAmount, senderCounter, receiverCounter); err != nil {
        log.Fatalf("Payment initiation failed: %v", err)
    }

    // **3. Verify Payment Registration Using GetPayment**

    data, err := prepareAndSignDatagram(senderUsername, receiverUsername, receiverServerAddress, commands.ClientPayments_GetPayment, senderCounter+1, nil)
    if err != nil {
        log.Fatalf("Failed to prepare GetPayment datagram for sender: %v", err)
    }

    response, err := sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive GetPayment for sender: %v", err)
    }

    paymentDetails, err := parsePaymentResponse(response)
    if err != nil {
        log.Fatalf("Failed to parse GetPayment response for sender: %v", err)
    }

    log.Printf("GetPayment response for sender: %s", paymentDetails)

    data, err = prepareAndSignDatagram(receiverUsername, senderUsername, senderServerAddress, commands.ClientPayments_GetPayment, receiverCounter+1, nil)
    if err != nil {
        log.Fatalf("Failed to prepare GetPayment datagram for receiver: %v", err)
    }

    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive GetPayment for receiver: %v", err)
    }

    paymentDetails, err = parsePaymentResponse(response)
    if err != nil {
        log.Fatalf("Failed to parse GetPayment response for receiver: %v", err)
    }

    log.Printf("GetPayment response for receiver: %s", paymentDetails)

    log.Println("Test passed: both sender and receiver initiated the payment successfully.")
}
