package tspgocrypt

type XORCrypter struct {
}

func (*XORCrypter) GetName() string {
  return "XORCrypter"
}

func (*XORCrypter) Encrypt(data []byte, key []byte, counter uint) int {
  return xorcrypt(data, key, counter)
}

func (*XORCrypter) Decrypt(data []byte, key []byte, counter uint) int {
  return xorcrypt(data, key, counter)
}

// in case of simple XOR based "encryption" encrypt and decrypt functions are
// the identical
func xorcrypt(data []byte, key []byte, counter uint) int {
  key_len := uint(len(key))
  encrypt_len := 0
  start_counter := counter % key_len

  for i := 0; i < len(data); i++ {
    data[i] = data[i] ^ key[start_counter % key_len]
    start_counter++
    encrypt_len++
  }

  return encrypt_len
}
