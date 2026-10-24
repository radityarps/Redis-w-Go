package main

import (
	"encoding/base64"
	"fmt"
	"redis/store"
	"redis/store/kv"
)

func main() {
	keyValueStore := kv.NewStore(20)
	PopulateDefaults(keyValueStore)
	newStore := keyValueStore.Clone()
	newStore.Set("env", "production")
	fmt.Println(keyValueStore)
	fmt.Println(newStore)
}

func PopulateDefaults(store store.Storer) error {
	defaults := map[string]string{
		"env":   "development",
		"ver":   "1.0.0",
		"debug": "true",
	}
	for key, value := range defaults {
		if err := store.Set(key, value); err != nil {
			return fmt.Errorf("PopulateDefaults: %w", err)
		}
	}
	return nil
}

func SetKeyWithEncryption(store store.Storer, key, value string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	store.Len()
	return store.Get(key)
}
