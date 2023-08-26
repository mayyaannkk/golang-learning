## Path

1. [Hello World](helloworld/hello.go) - [Testing hello world](helloworld/hello_test.go)
2. [Intgers main](integers/adder.go) - [Testing Integers](integers/adder_test.go)

---

## Notes

Writing a test is just like writing a function, with a few rules:

1. It needs to be in a file with a name like `xxx_test.go`
2. The test function must start with the word `Test`
3. The test function takes one argument only `t *testing.T`
4. In order to use the `*testing.T` type, you need to import "testing", like we did with fmt in the other file
