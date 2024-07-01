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

func AES128enc(message, key string) string {
	generateRoundKeys([]byte(key))
	for h, i := range roundKeys {
		fmt.Println("key - ", h+1, " :", i)
	}
	if len(message) <= 16 {
		plaintextbite := padText(message)
		fmt.Println("text bite: ", plaintextbite)
		ciphertext := encrypt(plaintextbite)
		fmt.Printf("Encrypted: %x\n", ciphertext)
		return fmt.Sprintf("%x", ciphertext)
	} else {
		res := ""
		for len(message) != 0 {
			if len(message) >= 16 {
				block := message[:16]
				ciphertext := encrypt([]byte(block))
				res += fmt.Sprintf("%x", ciphertext)
				message = message[16:]
			} else {
				lastblock := padText(message)
				ciphertext := encrypt(lastblock)
				res += fmt.Sprintf("%x", ciphertext)
				message = ""
			}
		}
		return res
	}
}
