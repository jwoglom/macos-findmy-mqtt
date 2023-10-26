package fmipcore

type FmItem struct {
	Name                 string            `json:"name"`
	Address              FmItemAddress     `json:"address"`
	ProductType          FmItemProductType `json:"productType"`
	Location             FmItemLocation    `json:"location"`
	CrowdSourcedLocation FmItemLocation    `json:"crowdSourcedLocation"`
	Identifier           string            `json:"identifier"`
	GroupIdentifier      string            `json:"groupIdentifier"`
	ProductIdentifier    string            `json:"productIdentifier"`
	SerialNumber         string            `json:"serialNumber"`
	SystemVersion        string            `json:"systemVersion"`
	PartInfo             FmItemPartInfo    `json:"partInfo"`
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

type FmItemPartInfo struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
