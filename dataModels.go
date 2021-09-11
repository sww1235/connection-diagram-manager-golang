package main

import (
	"time"

	"gorm.io/gorm"
)

// Definition of all data models used in project

// structs and fields must be exported for GORM to work
// gorm will normalize table and column names to
// snake_case all lower case, and pluralize table names

// data models use CamelCase here due to golang
// best practices.

// BasicFields is used as an embedded struct in other
// data model structs. It is not used by itself.
type BasicFields struct {
	ID          uint
	Name        string  `gorm:"size:50;not null"`
	Description *string `gorm:"size:500"`
}

// Identifier is used as an embedded struct in other
// data model structs. It is not used by itself
// It is not incorporated into BasicFields
// because some tables don't need an identifer
// field
type Identifier struct {
	Identifier *string `gorm:"size:50"`
}

// AuditFields is used as an embedded struct in other
// data model structs. It is not used by itself.
type AuditFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type ConnectorGender struct {
	BasicFields
}

type PhysicalMedia struct {
	BasicFields
}

type Direction struct {
	BasicFields
}

type SignalType struct {
	BasicFields
}

type Manufacturer struct {
	BasicFields
}

type MountingType struct {
	BasicFields
}

type EquipmentType struct {
	BasicFields
}

type Color struct {
	BasicFields
	Abbreviation string `gorm:"size:3;not null"`
}

type LocationType struct {
	BasicFields
}

// Location represents an individual location for equipment
type Location struct {
	BasicFields
	Identifier
	LocationTypeID   *uint
	LocationType     *LocationType
	NumberDinRails   *uint
	DinRailPosns     *uint //TODO: build relation table for DIN rails. Don't use JSON
	Width            *float64
	Height           *float64
	Depth            *float64
	AvailWidth       *float64
	AvailHeight      *float64
	AvailDepth       *float64
	PhysicalLocation *string `gorm:"size:100"`
}

// ConnectorModel represents a model or type of connector
// Examples include XLR3M, SMA-F, etc.
type ConnectorModel struct {
	BasicFields
	// lookup table foreign key
	// maps to belongs_to relationship.
	ConGenderID     uint
	ConnectorGender ConnectorGender `gorm:"foreignKey:ConGenderID"`
	DirectionID     uint            `gorm:"not null"`
	Direction       Direction       `gorm:"not null"`
	Model           *string         `gorm:"size:50"`
	Pluggable       bool            `gorm:"not null"` //TODO: fix this to account for screw terminal connectors
	NumberOfPins    *uint
	ManufacturerID  *uint
	Manufacturer    *Manufacturer
	ColorID         *uint
	Color           *Color
	Mates           []*ConnectorModel `gorm:"many2many:connector_model_mating"`
}

//type ConnectorModelMating struct {
//	ConnectorA uint `gorm:"primaryKey"`
//	ConnectorB uint `gorm:"primaryKey"`
//}

// EquipmentModel represents a model of equipment
type EquipmentModel struct {
	BasicFields
	Width           *float64
	Height          *float64
	Depth           *float64
	Diameter        *float64
	ManufacturerID  *uint
	Manufacturer    *Manufacturer
	Model           *string
	EquipmentTypeID uint          `gorm:"not null"`
	EquipmentType   EquipmentType `gorm:"not null"`
	NumberFaces     uint          `gorm:"not null"`
	RackPercent     *uint
}

// EquipmentInst represents a particular instance of a model of equipment
type EquipmentInst struct {
	BasicFields
	Identifier
	MountingTypeID   uint         `gorm:"not null"`
	MountingType     MountingType `gorm:"not null"`
	LocationID       *uint
	Location         *Location
	RackUnit         *uint
	DinRailNumber    *uint
	EquipmentModelID uint           `gorm:"not null"`
	EquipmentModel   EquipmentModel `gorm:"not null"`
	XPos             float64        `gorm:"not null"`
	YPos             float64        `gorm:"not null"`
	ZPos             float64        `gorm:"not null"`
}

// EquipmentConnector represents a particular connector on a particular piece of equipment.
type EquipmentConnector struct {
	BasicFields
	Identifier
	XPos             float64        `gorm:"not null"`
	YPos             float64        `gorm:"not null"`
	ZPos             float64        `gorm:"not null"`
	Face             uint           `gorm:"not null"`
	ConnectorModelID uint           `gorm:"not null"`
	ConnectorModel   ConnectorModel `gorm:"not null"`
	EquipmentInstID  uint           `gorm:"not null"`
	EquipmentInst    EquipmentInst  `gorm:"not null"`
	SignalTypeID     uint           `gorm:"not null"`
	SignalType       SignalType     `gorm:"not null"`
}

// Connection represents a wire or cable with connectors
//Siamese cables are considered one connection
type Connection struct {
	BasicFields
	Identifier
	RouteID         *uint
	ConnectionRoute *ConnectionRoute `gorm:"foreignKey:RouteID"`
	Length          *uint
	PhysicalMediaID *uint
	PhysicalMedia   *PhysicalMedia
	NumberOfEnds    *uint
}

type ConnectionRoute struct {
	BasicFields
	Identifer *string `gorm:"size:50"`
}

// ConnectionConnector represents a particular connector associated with a connection
type ConnectionConnector struct {
	BasicFields
	Identifier
	ConnectionID     uint           `gorm:"not null"`
	Connection       Connection     `gorm:"not null"`
	ConnectionEnd    uint           `gorm:"not null"`
	ConnectorModelID uint           `gorm:"not null"`
	ConnectorModel   ConnectorModel `gorm:"not null"`
}

type EquipmentConnectorConnectionConnector struct {
	BasicFields
	Identifier
	ConnectoionConnectorID uint                `gorm:"not null"`
	ConnectionConnector    ConnectionConnector `gorm:"not null"`
	EquipmentConnectorID   uint                `gorm:"not null"`
	EquipmentConnector     EquipmentConnector  `gorm:"not null"`
}
