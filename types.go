package main

import "fmt"

//TODO: better document types

type SVG string

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
	Visrep        SVG      `yaml:"visrep"`
	PinVisrep     SVG      `yaml:"pin_visrep"`
}

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
	Manufacturer string         `yaml:"manufacturer"`
	Model        string         `yaml:"model"`
	Pn           string         `yaml:"pn"`
	Mpn          string         `yaml:"mpn"`
	Supplier     string         `yaml:"supplier"`
	Spn          string         `yaml:"spn"`
	MountingType []string       `yaml:"mounting_type"`
	Type         string         `yaml:"type"`
	Faces        map[string]SVG `yaml:"faces"`
	Visrep       SVG            `yaml:"visrep"`
	Connectors   map[string]EquipmentConnector
}

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

//WireType represents a particular type of wire
type WireType struct {
	Material           string  `yaml:"material"`
	Manufacturer       string  `yaml:"manufacturer"`
	Pn                 string  `yaml:"pn"`
	Mpn                string  `yaml:"mpn"`
	Supplier           string  `yaml:"supplier"`
	Spn                string  `yaml:"spn"`
	Insulated          bool    `yaml:"insulated"`
	InsulationMaterial string  `yaml:"insulation_material"`
	WireTypeCode       string  `yaml:"wire_type_code"`
	CrossSectArea      float64 `yaml:"cross_sect_area"`
	Stranded           bool    `yaml:"stranded"`
	NumStrands         int     `yaml:"num_strands"`
	StrandCrossArea    float64 `yaml:"strand_cross_area"`
	InsulVoltRating    float64 `yaml:"insul_volt_rating"`
	InsulTempRating    float64 `yaml:"insul_temp_rating"`
}

//TODO: Implement
func (wt WireType) String() string {
	var tempString string

	return tempString

}

//CableCore represents an individual wire inside a cableType
type CableCore struct {
	Type  string `yaml:"type"`
	Color string `yaml:"color"`
}

type CableLayer struct {
	LayerNbr   int     `yaml:"layer"`
	Type       string  `yaml:"type"`
	Material   string  `yaml:"material"`
	VoltRating float64 `yaml:"volt_rating"`
	TempRating float64 `yaml:"temp_rating"`
	Color      string  `yaml:"color"`
}

//CableType represents a particular type of cable
type CableType struct {
	CableCore     map[string]CableCore
	Manufacturer  string       `yaml:"manufacturer"`
	Pn            string       `yaml:"pn"`
	Mpn           string       `yaml:"mpn"`
	Supplier      string       `yaml:"supplier"`
	Spn           string       `yaml:"spn"`
	CableTypeCode string       `yaml:"cable_type_code"`
	CrossSectArea float64      `yaml:"cross_sect_area"`
	CrossSection  string       `yaml:"cross_section"`
	Height        float64      `yaml:"height"`
	Width         float64      `yaml:"width"`
	Diameter      float64      `yaml:"diameter"`
	Layer         []CableLayer `yaml:"layer"`
}

//TODO: Implement
func (ct CableType) String() string {
	var tmpString string

	return tmpString

}

type TermCableConnectorTermination struct {
	Core int `yaml:"core"`
	Pin  int `yaml:"pin"`
}

type TermCableConnector struct {
	//connectorType
	Type        ConnectorType                   `yaml:"type"`
	Termination []TermCableConnectorTermination `yaml:"termination"`
}

//TODO: Implement
func (tcc TermCableConnector) String() string {
	var tempString string
	return tempString
}

type TermCableType struct {
	Manufacturer string               `yaml:"manufacturer"`
	Pn           string               `yaml:"pn"`
	Mpn          string               `yaml:"mpn"`
	Supplier     string               `yaml:"supplier"`
	Spn          string               `yaml:"spn"`
	Cable        CableType            `yaml:"cable"`
	Wire         WireType             `yaml:"wire"`
	NomLength    float64              `yaml:"nom_length"`
	Length       float64              `yaml:"length"`
	End1         []TermCableConnector `yaml:"end1"` // cables can have multiple connectors on each end
	End2         []TermCableConnector `yaml:"end2"`
}

//TODO: maybe implement formatter to allow for printing of just values, all fields and values or just present fields.
func (tct TermCableType) String() string {
	var tmpString string

	tmpString = "---TermCableType Struct---\n"
	//if tct == (TermCableType{}) {
	tmpString += fmt.Sprintf("Manufacturer: %s\n", tct.Manufacturer)
	tmpString += fmt.Sprintf("Part Number: %s\n", tct.Pn)
	tmpString += fmt.Sprintf("Manufacturer Part Number: %s\n", tct.Mpn)
	tmpString += fmt.Sprintf("Supplier: %s\n", tct.Supplier)
	tmpString += fmt.Sprintf("Supplier Part Number: %s\n", tct.Spn)
	tmpString += fmt.Sprintf("Cable Type: \t%s\n", tct.Cable)
	tmpString += fmt.Sprintf("Wire Type: \t%s\n", tct.Wire)
	tmpString += fmt.Sprintf("Nominal Length: %g\n", tct.NomLength)
	tmpString += fmt.Sprintf("Actual Length: %g\n", tct.Length)
	tmpString += fmt.Sprintf("End1: \t%s\n", tct.End1)
	tmpString += fmt.Sprintf("End2: \t%s\n", tct.End2)
	//}
	return tmpString
}

type LocationType struct {
	Manufacturer string `yaml:"manufacturer"`
	Pn           string `yaml:"pn"`
	Mpn          string `yaml:"mpn"`
	Supplier     string `yaml:"supplier"`
	Spn          string `yaml:"spn"`
}
