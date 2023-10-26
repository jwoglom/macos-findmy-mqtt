package fmipcore

type FmItem struct {
	Name                 string               `json:"name"`
	Address              FmItemAddress        `json:"address"`
	ProductType          FmItemProductType    `json:"productType"`
	Location             FmItemLocation       `json:"location"`
	CrowdSourcedLocation FmItemLocation       `json:"crowdSourcedLocation"`
	SafeLocations        []FmItemSafeLocation `json:"safeLocations"`
	Identifier           string               `json:"identifier"`
	GroupIdentifier      string               `json:"groupIdentifier"`
	ProductIdentifier    string               `json:"productIdentifier"`
	SerialNumber         string               `json:"serialNumber"`
	SystemVersion        string               `json:"systemVersion"`
	Role                 FmItemRole           `json:"role"`
	PartInfo             FmItemPartInfo       `json:"partInfo"`
}

type FmItemAddress struct {
	Label                 string   `json:"Label"`
	MapItemFullAddress    string   `json:"mapItemFullAddress"`
	FormattedAddressLines []string `json:"formattedAddressLines"`
}

type FmItemProductType struct {
	ProductInformation FmItemProductTypeInfo `json:"productInformation"`
	Type               string                `json:"type"`
}

func (t FmItem) IsAirTag() bool {
	return t.ProductType.Type == "b389"
}

type FmItemProductTypeInfo struct {
	ModelName string `json:"modelName"`
}

type FmItemLocation struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Timestamp    int64   `json:"timeStamp"`
	IsInaccurate bool    `json:"isInaccurate"`
	IsOld        bool    `json:"isOld"`
}

type FmItemSafeLocation struct {
	Name     string         `json:"name"`
	Location FmItemLocation `json:"location"`
	Address  FmItemAddress  `json:"address"`
}

type FmItemPartInfo struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type FmItemRole struct {
	Emoji string `json:"string"`
}
