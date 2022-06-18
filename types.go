package main

//ConnectorType represents a particular type of connector
type ConnectorType struct {
	Manufacturer  string   `yaml:"manufacturer"`
	Model         string   `yaml:"model"`
	Description   string   `yaml:"description"`
	MountingType  []string `yaml:"mounting_type"`
	PanelCutout   string   `yaml:"panel_cutout"`
	Pn            string   `yaml:"pn"`
	Mpn           string   `yaml:"mpn"`
	Supplier      string   `yaml:"supplier"`
	Spn           string   `yaml:"spn"`
	Gender        string   `yaml:"gender"`
	Height        float64  `yaml:"height"`
	Width         float64  `yaml:"width"`
	Depth         float64  `yaml:"depth"`
	Diameter      float64  `yaml:"diameter"`
	Pincount      int      `yaml:"pincount"`
	Pins          []string `yaml:"pins"`
	Pinlabels     []string `yaml:"pinlabels"`
	Pincolors     []string `yaml:"pincolors"`
	PinSignalType []string `yaml:"pin_signal_type"`
	Visrep        string   `yaml:"visrep"`
	PinVisrep     string   `yaml:"pin_visrep"`
}

//ConnectorTypes stores multiple ConnectorType definitions
type ConnectorTypes map[string]ConnectorType

//EquipmentConnector stores a particular instance
//of a connectorType and its location on a particular type of equipment.
type EquipmentConnector struct {
	Connector ConnectorType
	Direction string `yaml:"direction:"`
	Face      string `yaml:"face"`
	X         int    `yaml:"x"`
	Y         int    `yaml:"y"`
}

//EquipmentType represents a particular type of equipment
type EquipmentType struct {
	Manufacturer string   `yaml:"manufacturer"`
	Model        string   `yaml:"model"`
	Pn           string   `yaml:"pn"`
	Mpn          string   `yaml:"mpn"`
	Supplier     string   `yaml:"supplier"`
	Spn          string   `yaml:"spn"`
	MountingType []string `yaml:"mounting_type"`
	Type         string   `yaml:"type"`
	Faces        struct {
		Face1 string `yaml:"face1"`
		Face2 string `yaml:"face2"`
	} `yaml:"faces"`
	Visrep     string `yaml:"visrep"`
	Connectors map[string]EquipmentConnector
}

//EquipmentTypes stores multiple EquipmentType definitions
type EquipmentTypes map[string]EquipmentType

//PathwayType represents a particular type of pathway
type PathwayType struct {
	Type          string  `yaml:"type"`
	Pn            string  `yaml:"pn"`
	Mpn           string  `yaml:"mpn"`
	Supplier      string  `yaml:"supplier"`
	Spn           string  `yaml:"spn"`
	Size          string  `yaml:"size"`
	TradeSize     string  `yaml:"trade_size"`
	CrossSectArea float64 `yaml:"cross_sect_area"`
	Material      string  `yaml:"material"`
}

//PathwayTypes stores multiple PathwayType defintions
type PathwayTypes map[string]PathwayType

//WireType represents a particular type of wire
type WireType struct {
	Material           string  `yaml:"material"`
	Manufacturer       string  `yaml:"manufacturer"`
	Pn                 string  `yaml:"pn"`
	Mpn                string  `yaml:"mpn"`
	Supplier           string  `yaml:"supplier"`
	Spn                string  `yaml:"spn"`
	InsulationMaterial string  `yaml:"insulation_material"`
	WireTypeCode       string  `yaml:"wire_type_code"`
	CrossSectArea      float64 `yaml:"cross_sect_area"`
	Stranded           bool    `yaml:"stranded"`
	NumStrands         int     `yaml:"num_strands"`
	StrandCrossArea    float64 `yaml:"strand_cross_area"`
	InsulVoltRating    float64 `yaml:"insul_volt_rating"`
	InsulTempRating    float64 `yaml:"insul_temp_rating"`
}

//WireTypes stores multiple WireType defintions
type WireTypes map[string]WireType

//CableCore represents an individual wire inside a cableType
type CableCore struct {
	Type  string `yaml:"type"`
	Color string `yaml:"color"`
}

//CableType represents a particular type of cable
type CableType struct {
	CableCore     map[string]CableCore
	Manufacturer  string  `yaml:"manufacturer"`
	Pn            string  `yaml:"pn"`
	Mpn           string  `yaml:"mpn"`
	Supplier      string  `yaml:"supplier"`
	Spn           string  `yaml:"spn"`
	CableTypeCode string  `yaml:"cable_type_code"`
	CrossSectArea float64 `yaml:"cross_sect_area"`
	CrossSection  string  `yaml:"cross_section"`
	Height        float64 `yaml:"height"`
	Width         float64 `yaml:"width"`
	Diameter      float64 `yaml:"diameter"`
	Layer         []struct {
		LayerNbr   int     `yaml:"layer"`
		Type       string  `yaml:"type"`
		Material   string  `yaml:"material"`
		VoltRating float64 `yaml:"volt_rating"`
		TempRating float64 `yaml:"temp_rating"`
		Color      string  `yaml:"color"`
	} `yaml:"layer"`
}

//CableTypes stores multiple CableType defintions
type CableTypes map[string]CableType

type TermCableType struct {
	Manufacturer string  `yaml:"manufacturer"`
	Pn           string  `yaml:"pn"`
	Mpn          string  `yaml:"mpn"`
	Supplier     string  `yaml:"supplier"`
	Spn          string  `yaml:"spn"`
	Cable        string  `yaml:"cable"`
	NomLength    float64 `yaml:"nom_length"`
	Length       float64 `yaml:"length"`
	End1         []struct {
		Type        string `yaml:"type"`
		AutoTerm    string `yaml:"autoTerm"`
		Termination []struct {
			Core int `yaml:"core"`
			Pin  int `yaml:"pin"`
		} `yaml:"termination"`
	} `yaml:"end1"`
	End2 []struct {
		Type        string `yaml:"type"`
		AutoTerm    string `yaml:"autoTerm"`
		Termination []struct {
			Core int `yaml:"core"`
			Pin  int `yaml:"pin"`
		} `yaml:"termination"`
	} `yaml:"end2"`
}
