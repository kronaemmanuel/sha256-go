# SHA-256 in Go

This is not a serious/efficient implementation of SHA-256 in Go. You can use the [official package](https://pkg.go.dev/crypto/sha256) for that. This is just for learning purposes, which you can probably tell because I'm using strings to hold data instead of storing it in binary/hex form.

This project is part of small fun experiments I do to challenge and improve my coding skills. I write about these experiments on my [blog](https://www.kronaemmanuel.com/blog). You can visit the blog post about this project [here](https://www.kronaemmanuel.com/2021/09/12/Implementing-SHA-256-in-Go-lang.html)

# How to run
- Clone the repo
- Make sure you have Go [installed](https://golang.org/doc/install)
- Go the the project directory
- Run `go run .`

You'll have to change the string inside the code itself, to be able to get a different hash at the end. the code output shows a lot of different things, the hash will be given at the end. For example, when I was testing, and put in the password to my bank account, it gave me this hash at the end:
```
Output of step 8: 
ec894281e55b2adf4b55f27cee33812bd4ac98ae65b21f4bf3e272b5567f5222
```

# Resources:
Main guide I used was:
- [Qvault: How SHA-256 works guide](https://qvault.io/cryptography/how-sha-2-works-step-by-step-sha-256/)

Other very useful resources were:
  - [Go Docs: Getting Started](https://golang.org/doc/tutorial/getting-started)
  - [Go Docs: Create a Module](https://golang.org/doc/tutorial/create-module)
  - [Willf: String right/left pad library](https://github.com/willf/pad/blob/master/pad.go)
  - [GeeksForGeeks: How to rotate a string](https://www.geeksforgeeks.org/left-rotation-right-rotation-string-2/)
  - [GeeksForGeeks: How to XOR two binary strings](https://www.geeksforgeeks.org/xor-of-two-binary-strings/)
  - [GeeksForGeeks: How to reverse a string in Go](https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/)
  - [Stack Overflow: Convert a binary string to hexadecimal string](https://stackoverflow.com/questions/25592084/converting-binary-string-to-a-hexadecimal-string-java)