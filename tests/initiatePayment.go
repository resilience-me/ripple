package tests

import (
    "encoding/binary"
    "fmt"
    "log"
    "ripple/commands"
)

func InitiatePayment(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, serverAddress string, paymentAmount, senderCounter, receiverCounter uint32) error {
    // Prepare the arguments (the amount to be paid)
    arguments := make([]byte, 4)
    binary.BigEndian.PutUint32(arguments[:4], paymentAmount)

    // **1. Sender Initiates Payment**
    data, err := prepareAndSignDatagram(senderUsername, receiverUsername, receiverServerAddress, commands.ClientPayments_NewPaymentOut, senderCounter, arguments)
    if err != nil {
        return fmt.Errorf("failed to prepare payment datagram for sender: %w", err)
    }

    response, err := sendAndReceive(serverAddress, data)
    if err != nil {
        return fmt.Errorf("failed to send and receive for sender: %w", err)
    }

    log.Printf("Sender response: %s", string(response))

    // **2. Receiver Initiates Payment**
    data, err = prepareAndSignDatagram(receiverUsername, senderUsername, senderServerAddress, commands.ClientPayments_NewPaymentIn, receiverCounter, arguments)
    if err != nil {
        return fmt.Errorf("failed to prepare payment datagram for receiver: %w", err)
    }

    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        return fmt.Errorf("failed to send and receive for receiver: %w", err)
    }

    log.Printf("Receiver response: %s", string(response))

    return nil
}
