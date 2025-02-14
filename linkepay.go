package linkepay

// Export functions by starting with uppercase
func HelloWorld() string {
	return "Hello from mypackage!"
}

// Unexported function (private) starts with lowercase
func helper() string {
	return "I'm a helper function"
}
