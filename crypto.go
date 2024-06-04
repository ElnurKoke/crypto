package crypto

import "fmt"

func AES128(word16byte, key16byte string) {
	key := key16byte
	plaintext := word16byte
	plaintextbite := padText(plaintext)
	fmt.Println("text bite: ", plaintextbite)

	generateRoundKeys([]byte(key))
	for h, i := range roundKeys {
		fmt.Println("key - ", h+1, " :", i)
	}

	ciphertext := encrypt(plaintextbite)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	decrypted := decrypt(ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
