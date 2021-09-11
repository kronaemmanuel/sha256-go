/*
  Implementing SHA-256 algorithm in GO Language

  This algorithm was implemented by following this guide:
  - https://qvault.io/cryptography/how-sha-2-works-step-by-step-sha-256/

  Along the way, the following resources were also very helpful:
  - https://golang.org/doc/tutorial/getting-started
  - https://golang.org/doc/tutorial/create-module
  - https://github.com/willf/pad/blob/master/pad.go
  - https://www.geeksforgeeks.org/left-rotation-right-rotation-string-2/
  - https://www.geeksforgeeks.org/xor-of-two-binary-strings/
  - https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
  - https://stackoverflow.com/questions/25592084/converting-binary-string-to-a-hexadecimal-string-java

  This was my first Go program, so I didn't know how to implement this in a low
  level way, so its done using strings which is not very memory efficient of
  course. But this is more for learning purposes rather than efficiency of code.

  Limitations:
  - Doesn't accept big strings, just small strings, since there is just one block
*/
package main

import (
  "fmt"
  "errors"
  "strings"
  "log"
  "strconv"
)

func main() {
  // Setting up error logging
  log.SetPrefix("Hashing: ")
  log.SetFlags(0)

  // Initial String
  initString := "hello world"

  // Step 1: Preprocessing the data:
  chunk := preprocess(initString)
  fmt.Printf("Output of step 1: \n")
  binaryPrettyPrint(chunk)

  // Step 2: Initialize Hash Values
  h0 := "6a09e667"
  h1 := "bb67ae85"
  h2 := "3c6ef372"
  h3 := "a54ff53a"
  h4 := "510e527f"
  h5 := "9b05688c"
  h6 := "1f83d9ab"
  h7 := "5be0cd19"

  fmt.Printf("\n\nOutput of step 2: \n")
  fmt.Printf("%x %x %x %x %x %x %x, %x", h0, h1, h2, h3, h4, h5, h6, h7)

  // Step 3: Initialize Round Constants
  k := []string{
    "428a2f98", "71374491", "b5c0fbcf", "e9b5dba5", "3956c25b", "59f111f1", "923f82a4", "ab1c5ed5",
    "d807aa98", "12835b01", "243185be", "550c7dc3", "72be5d74", "80deb1fe", "9bdc06a7", "c19bf174",
    "e49b69c1", "efbe4786", "0fc19dc6", "240ca1cc", "2de92c6f", "4a7484aa", "5cb0a9dc", "76f988da",
    "983e5152", "a831c66d", "b00327c8", "bf597fc7", "c6e00bf3", "d5a79147", "06ca6351", "14292967",
    "27b70a85", "2e1b2138", "4d2c6dfc", "53380d13", "650a7354", "766a0abb", "81c2c92e", "92722c85",
    "a2bfe8a1", "a81a664b", "c24b8b70", "c76c51a3", "d192e819", "d6990624", "f40e3585", "106aa070",
    "19a4c116", "1e376c08", "2748774c", "34b0bcb5", "391c0cb3", "4ed8aa4a", "5b9cca4f", "682e6ff3",
    "748f82ee", "78a5636f", "84c87814", "8cc70208", "90befffa", "a4506ceb", "bef9a3f7", "c67178f2",
  }

  fmt.Printf("\n\nOutput of step 3: \n")
  kBinary := make([]string, len(k))
  for index, constant := range k {
    kBinary[index] = hexToBinaryString(constant)
    fmt.Printf("Constant_%d: %s\n", index, kBinary[0])
  }

  // Step 4: Chunk Loop
  // In our case there is only one chunk, so we don't have a loop here

  // Step 5.1: Copy the chunk from step 1 into a new array where each entry is a
  // 32-bit word
  numOfWordsInBlock := 16 // 512 / 32 = 16
  chunkWords := make([]string, numOfWordsInBlock)
  for i := 0; i < numOfWordsInBlock; i++ {
    startIndex := i * 32
    chunkWords[i] = chunk[startIndex:startIndex+32]
  }

  fmt.Printf("\n\nOutput of step 5.1: \n")
  blockPrettyPrint(chunkWords)

  // Step 5.2: Add 48 more words, initialized to zero, so total length is 64
  numOfZeroWords := 48
  zeroWords := make([]string, numOfZeroWords)
  for index := range zeroWords {
    zeroWords[index] = strings.Repeat("0", 32)
  }
  block := append(chunkWords, zeroWords...)

  fmt.Printf("\n\nOutput of step 5.2: \n")
  blockPrettyPrint(block)

  // Step 5.3: Modify the words from block[16...63] as following:
  // For i from w[16…63]:
  //    s0 = (w[i-15] rightrotate 7) xor (w[i-15] rightrotate 18) xor (w[i-15] rightshift 3)
  //    s1 = (w[i- 2] rightrotate 17) xor (w[i- 2] rightrotate 19) xor (w[i- 2] rightshift 10)
  //    w[i] = w[i-16] + s0 + w[i-7] + s1

  for i := 16; i < len(block); i++ {
    s0 := xor(xor(rightRotate(block[i-15], 7), rightRotate(block[i-15], 18)), rightShift(block[i-15], 3))
    s1 := xor(xor(rightRotate(block[i-2], 17), rightRotate(block[i-2], 19)), rightShift(block[i-2], 10))
    block[i] = binaryAdd(binaryAdd(binaryAdd(block[i-16], s0), block[i-7]), s1)
  }

  fmt.Printf("\n\nOutput of step 5.3: \n")
  blockPrettyPrint(block)

  // Step 6: Compression

  // Step 6.1: Initialize variables a...h with same values as binary strings of
  // h0...h7
  a := "01101010000010011110011001100111"
  b := "10111011011001111010111010000101"
  c := "00111100011011101111001101110010"
  d := "10100101010011111111010100111010"
  e := "01010001000011100101001001111111"
  f := "10011011000001010110100010001100"
  g := "00011111100000111101100110101011"
  h := "01011011111000001100110100011001"

  // Step 6.2 Run the compression loop
  /* for i from 0 to 63
        S1 = (e rightrotate 6) xor (e rightrotate 11) xor (e rightrotate 25)
        ch = (e and f) xor ((not e) and g)
        temp1 = h + S1 + ch + k[i] + w[i]
        S0 = (a rightrotate 2) xor (a rightrotate 13) xor (a rightrotate 22)
        maj = (a and b) xor (a and c) xor (b and c)
        temp2 := S0 + maj
        h = g
        g = f
        f = e
        e = d + temp1
        d = c
        c = b
        b = a
        a = temp1 + temp2
  */

  for i := 0; i < 64; i++ {
    s1 := xor(xor(rightRotate(e, 6), rightRotate(e, 11)), rightRotate(e, 25))
    ch := xor(and(e, f), and(not(e), g))
    temp1 := binaryAdd(binaryAdd(binaryAdd(binaryAdd(h, s1), ch), kBinary[i]), block[i])
    s0 := xor(xor(rightRotate(a, 2), rightRotate(a, 13)), rightRotate(a, 22))
    maj := xor(xor(and(a, b), and(a, c)), and(b, c))
    temp2:= binaryAdd(s0, maj)
    h = g
    g = f
    f = e
    e = binaryAdd(d, temp1)
    d = c
    c = b
    b = a
    a = binaryAdd(temp1, temp2)

  }

  fmt.Printf("\n\nOutput of step 6.2: \n")

  fmt.Printf("h0: %s\n", hexToBinaryString(h0))
  fmt.Printf("h1: %s\n", hexToBinaryString(h1))
  fmt.Printf("h2: %s\n", hexToBinaryString(h2))
  fmt.Printf("h3: %s\n", hexToBinaryString(h3))
  fmt.Printf("h4: %s\n", hexToBinaryString(h4))
  fmt.Printf("h5: %s\n", hexToBinaryString(h5))
  fmt.Printf("h6: %s\n", hexToBinaryString(h6))
  fmt.Printf("h7: %s\n\n", hexToBinaryString(h7))
  fmt.Printf("a: %s\n", a)
  fmt.Printf("b: %s\n", b)
  fmt.Printf("c: %s\n", c)
  fmt.Printf("d: %s\n", d)
  fmt.Printf("e: %s\n", e)
  fmt.Printf("f: %s\n", f)
  fmt.Printf("g: %s\n", g)
  fmt.Printf("h: %s\n", h)

  // Step 7: Modify Final Values
  h0 = binaryAdd(hexToBinaryString(h0), a)
  h1 = binaryAdd(hexToBinaryString(h1), b)
  h2 = binaryAdd(hexToBinaryString(h2), c)
  h3 = binaryAdd(hexToBinaryString(h3), d)
  h4 = binaryAdd(hexToBinaryString(h4), e)
  h5 = binaryAdd(hexToBinaryString(h5), f)
  h6 = binaryAdd(hexToBinaryString(h6), g)
  h7 = binaryAdd(hexToBinaryString(h7), h)

  fmt.Printf("\n\nOutput of step 7: \n")
  fmt.Printf("h0: %s\n", h0)
  fmt.Printf("h1: %s\n", h1)
  fmt.Printf("h2: %s\n", h2)
  fmt.Printf("h3: %s\n", h3)
  fmt.Printf("h4: %s\n", h4)
  fmt.Printf("h5: %s\n", h5)
  fmt.Printf("h6: %s\n", h6)
  fmt.Printf("h7: %s\n\n", h7)

  // Step 8: Concatenate Final Hash
  hash := h0 + h1 + h2 + h3 + h4 + h5 + h6 + h7
  hashString := binaryToHexString(hash)

  fmt.Printf("\n\nOutput of step 8: \n")
  fmt.Printf("%s\n", hashString)
}

/*
  Converts an input string and outputs a string representing the binary data
  of the input string
*/
func stringToBin(s string) (binString string) {
    for _, c := range s {
        binString = fmt.Sprintf("%s%.8b",binString, c)
    }
    return
}

/*
  Takes a string and prints it with a space after every 8 characters
*/
func binaryPrettyPrint(s string) {
   for i := 0; i < len(s); i = i+8 {
     if (i % 64 == 0 && i != 0) {
       fmt.Printf("\n")
     }
     if (i+8 < len(s)) {
       fmt.Printf("%s ", s[i:i+8])
     } else {
       fmt.Printf("%s", s[i:])
     }
   }
}

// TODO: Refactor Code to only include one pad function in which right/left
// is passed as a parameter
/*
  Takes an input string and pads it on the right with the given character until
  the total length of the padded string is equal to the passed totalLength
*/
func padRight(input string, totalLength int, character rune) (paddedString string, err error) {
  if (totalLength < len(input)) {
    return "", errors.New("Total length of padded string cannot be smaller than the length of input string")
  }

  return input + strings.Repeat(string(character), totalLength-len(input)), nil
}

/*
  Takes an input string and pads it on the right with the given character until
  the total length of the padded string is equal to the passed totalLength
*/
func padLeft(input string, totalLength int, character rune) (paddedString string, err error) {
  if (totalLength < len(input)) {
    return "", errors.New("Total length of padded string cannot be smaller than the length of input string")
  }

  return strings.Repeat(string(character), totalLength-len(input)) + input, nil
}

/*
  Inputs a string and outputs preprocessed chunk of data which has 512 bits
*/
func preprocess(input string) string {
  // Convert string to binary string
  binString := stringToBin(input)

  // Append a single 1 to the bindary string
  binStringAppended := binString + string('1')

  // Pad with 0's until data is a multiple of 512, less 64 bits
  // TODO: Handle case where binString is larger than 512 bits
  const paddedStringLength = 448 // 512 - 64 = 448
  paddedString, err := padRight(binStringAppended, paddedStringLength, '0')
  if err != nil {
    log.Fatal(err)
  }

  // Append 64 bits to the end, where the 64 bits are big endian representing
  // the length of our binString
  binStringLength := stringToBin(string(len(binString)))
  paddedBinStringLength, err := padLeft(binStringLength, 64, '0')

  return paddedString + paddedBinStringLength
}

/*
  Pretty prints a block which has 32 bit words in it
*/
func blockPrettyPrint(block []string) {
  for index, word := range block {
    if (index % 2 == 0 && index != 0) {
      fmt.Printf("\n")
    }
    fmt.Printf("%s ", word)
  }
}

/*
  Shifts the first character to the end x times where x is provided as rotateBy
*/
func leftRotate(input string, rotateBy int) (rotatedString string) {
  rotatedString = input
  rotatedString = rotatedString[rotateBy:len(input)] + rotatedString[0:rotateBy]
  return
}

/*
  Shifts the last character to the start x times where x is provided as rotateBy
*/
func rightRotate(input string, rotateBy int) string {
  return leftRotate(input, len(input) - rotateBy)
}

/*
  Add 0s to the start of the string by the amount specified by shiftBy
*/
func rightShift(input string, shiftBy int) (shiftedString string) {
  shiftedString = rightRotate(input, shiftBy)
  shiftedString = strings.Repeat("0", shiftBy) + shiftedString[shiftBy:]
  return
}

/*
  Function to find the XOR of two binary strings of the same length
*/
func xor(input1 string, input2 string) (result string) {
  for i := 0; i < len(input1); i++ {
    if input1[i] == input2[i] {
      result += "0"
    } else {
      result += "1"
    }
  }
  return
}


func binaryAdd(input1 string, input2 string) (result string) {
  carry := 0
  sum := 0
  for i := len(input1)-1; i >= 0; i-- {
    firstDigit, _ := strconv.Atoi(input1[i:i+1])
    secondDigit, _ := strconv.Atoi(input2[i:i+1])
    sum = firstDigit + secondDigit + carry

    if sum == 0 {
      result += "0"
      carry = 0
    } else if sum == 1 {
      result += "1"
      carry = 0
    } else if sum == 2 {
      result += "0"
      carry = 1
    } else if sum == 3 {
      result += "1"
      carry = 1
    }
  }

  result = reverseString(result)
  return
}

/*
  function to reverse a string
*/
func reverseString(input string) string {
  runes := []rune(input)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

/*
  Function to find the AND of two binary strings of the same length
*/
func and(input1 string, input2 string) (result string) {
  for i := 0; i < len(input1); i++ {
    if input1[i] == '1' && input2[i] == '1' {
      result += "1"
    } else {
      result += "0"
    }
  }
  return
}

/*
  Function to find the NOT of a binary string
*/
func not(input string) (result string) {
  for i := 0; i < len(input); i++ {
    if input[i] == '1' {
      result += "0"
    } else {
      result += "1"
    }
  }
  return
}

/*
  Function for converting Hex string into binary string
*/
func hexToBinaryString(hexString string) string {
  hexInt, _ := strconv.ParseUint(hexString, 16, 32)

	bits := []uint64{}
	for i := 0; i < 32; i++ {
		bits = append([]uint64{hexInt & 0x1}, bits...)
		hexInt = hexInt >> 1
	}

  bitString := ""
  for _, bit := range bits {
    bitString += strconv.Itoa(int(bit))
  }

	return bitString
}

/*
  Function for converting binary string into hex string
*/
func binaryToHexString(binString string) string {
  hexString := ""
  digitNumber := 1;
  sum := 0;
  binary := binString;
  for i := 0; i < len(binString); i++ {
      num, _ := strconv.Atoi(binary[i:i+1])
      if(digitNumber == 1){
        sum += num*8;
      } else if digitNumber == 2 {
        sum += num*4;
      } else if digitNumber == 3 {
        sum += num*2;
      } else if digitNumber == 4 || i < len(binary)+1 {
        sum += num*1;
        digitNumber = 0;
        if(sum < 10) {
          hexString += strconv.Itoa(sum)
        } else if(sum == 10) {
          hexString += "a"
        } else if(sum == 11) {
          hexString += "b"
        } else if(sum == 12) {
          hexString += "c"
        } else if(sum == 13) {
          hexString += "d"
        } else if(sum == 14) {
          hexString += "e"
        } else if(sum == 15) {
          hexString += "f"
        }
        sum=0;
      }
    digitNumber++;
  }
  return hexString
}
