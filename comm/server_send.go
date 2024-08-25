package comm

// Default Send with standard importance
func Send(destinationAddr string, data []byte) error {
	return sendWithResolvedAddress(destinationAddr, data, lowImportance)
}

// Send with priority importance
func SendPriority(destinationAddr string, data []byte) error {
	return sendWithResolvedAddress(destinationAddr, data, highImportance)
}
