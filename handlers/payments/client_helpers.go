package payments

import (
    "fmt"
    "ripple/types"
    "ripple/pathfinding"
)

// serializePaymentDetails constructs a byte array from the payment details
func serializePaymentDetails(payment *pathfinding.Payment, amount uint32) []byte {
    buffer := concatNameAndServer(payment.Counterpart.Username, payment.Counterpart.ServerAddress)
    buffer = append(buffer, payment.InOrOut)
    amountAndNonce := append(types.Uint32ToBytes(amount), types.Uint32ToBytes(payment.Nonce)...)
    buffer = append(buffer, amountAndNonce...)
    return buffer
}

// getPaymentAndPath retrieves the Payment and Path objects associated with the given username.
func getPaymentAndPath(username string) (*pathfinding.Payment, *pathfinding.Path, error) {
    account := pathfinding.GetPathManager().Find(username)
    if account == nil {
        return nil, nil, fmt.Errorf("no account found for username: %s", username)
    }
    if account.Payment == nil {
        return nil, nil, fmt.Errorf("no payment registered for username: %s", username)
    }

    // Extract payment details
    paymentDetails := account.Payment

    // Find the Path using the identifier in the Payment
    path := account.Find(paymentDetails.Identifier)
    if path == nil {
        return paymentDetails, nil, fmt.Errorf("no path found for payment with identifier: %x", paymentDetails.Identifier)
    }

    return paymentDetails, path, nil
}

// GetPaymentDetails retrieves payment identifier, amount, and direction.
func GetPaymentDetails(username string) ([32]byte, uint32, byte, error) {
    payment, path, err := getPaymentAndPath(username)
    if err != nil {
        // More verbose error handling here
        return "", 0, 0, fmt.Errorf("failed to retrieve payment details for user %s: %w", username, err)
    }
    return payment.Identifier, path.Amount, payment.InOrOut, nil
}

// Wrapper function to fetch and serialize payment details
func FetchAndSerializePaymentDetails(username string) []byte {
    payment, path, err := getPaymentAndPath(username)
    if err != nil {
        return []byte{}
    }
    return serializePaymentDetails(payment, path.Amount)
}
