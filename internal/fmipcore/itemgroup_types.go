package fmipcore

type FmItemGroup struct {
	Name                   string           `json:"name"`
	Identifier             string           `json:"identifier"`
	ItemPairingStateMap    FmItemPairingMap `json:"itemPairingStateMap"`
	Items                  []FmItem         `json:"items"`
	GroupedItemIdentifiers [][]string       `json:"groupedItemIdentifiers"`
	ItemIdentifiers        []string         `json:"itemIdentifiers"`
}

type FmItemPairingMap = map[string][]string

func (g FmItemGroup) PairingState(identifier string) string {
	for _, v := range g.ItemPairingStateMap {
		state := v[0]
		id := v[1]
		if id == identifier {
			return state
		}
	}

	return ""
}
