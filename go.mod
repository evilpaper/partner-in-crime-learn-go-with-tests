module github.com/evilpaper/partner-in-crime-learn-go-with-tests

go 1.24.5

// wft is go.mod
// It's like package.json for Go. Plus more.
// When you for, for example, go test, Go does roughly the following:
// 1. Search upward for the nearest go.mod
// 2. Treat that folder as the module root
// 3. Reads the go.mod file
// 4. Finds the module name
// 5. Finds the module version
// 6. Finds the module dependencies
// 7. Finds the module tests
// 8. Runs the tests
// 9. Prints the results
// 10. Prints the coverage
// 11. Prints the errors
// 12. Prints the warnings
// 13. Prints the summary