package main

import (
	"encoding/base64"
	"fmt"
	"redis/store"
	"redis/store/kv"
	"redis/store/ttl"
	"time"
)

func main() {
	keyValueStore := kv.NewStore(20)
	encrypted, err := SetKeyWithEncryption(keyValueStore, "a", "1")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(encrypted)

	ttlStore := ttl.NewTTLStore(2 * time.Second)
	encryptedTtlValue, err := SetKeyWithEncryption(ttlStore, "a", "1")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(encryptedTtlValue)
}

func SetKeyWithEncryption(store store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)
}
