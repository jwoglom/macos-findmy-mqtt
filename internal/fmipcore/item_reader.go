package fmipcore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func itemsDataPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Library/Caches/com.apple.findmy.fmipcore/Items.data")
}

func ReadItems() ([]FmItem, error) {
	f, err := os.Open(itemsDataPath())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return readItems(bytes)
}

func readItems(bytes []byte) ([]FmItem, error) {
	var items []FmItem
	json.Unmarshal(bytes, &items)

	return items, nil
}
