package comm

// Default Send with standard importance (5 retries)
func Send(destinationAddr string, data []byte) error {
	return sendWithResolvedAddress(destinationAddr, data, LowImportance)
}

// Send with priority importance (12 retries)
func SendPriority(destinationAddr string, data []byte) error {
	return sendWithResolvedAddress(destinationAddr, data, HighImportance)
}
