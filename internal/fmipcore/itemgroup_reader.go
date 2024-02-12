package fmipcore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func itemGroupsDataPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "Library/Caches/com.apple.findmy.fmipcore/ItemGroups.data")
}

func ReadItemGroups() ([]FmItemGroup, error) {
	f, err := os.Open(itemGroupsDataPath())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return readItemGroups(bytes)
}

func readItemGroups(bytes []byte) ([]FmItemGroup, error) {
	var items []FmItemGroup
	json.Unmarshal(bytes, &items)

	return items, nil
}

func GetItemGroup(identifier string) (*FmItemGroup, error) {
	groups, err := ReadItemGroups()
	if err != nil {
		return nil, err
	}

	for _, g := range groups {
		if g.Identifier == identifier {
			return &g, nil
		}
	}

	return nil, nil
}
