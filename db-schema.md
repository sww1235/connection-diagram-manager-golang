# Database Schema

## gender

| Field       | Datatype     | Description          |
| ----------- | ------------ | -------------------- |
| id          | int          | primary key          |
| name        | varchar(50)  |                      |
| description | varchar(500) | optional description |

### Option List
- Male
- Female
- RP-Male
- RP-Female
- Hermaphroditic
- Unknown

## direction

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List

- Input
- Output
- Power Input
- Power Output
- Bidirectional
- N/A

## signalType

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List (initial)

- AC power
- DC power
- Mic Level Audio
- Consumer Line Level Audio
- Professional Line Level Audio
- Composite Video
- S video
- SDI Video
- Red (Component Video)
- Green (Component Video)
- Blue (Component Video)
- Clock+
- Clock-
- Ground
- Data+
- Data-
- 0-10V Analog
- 4-20mA Analog
- Unknown

## manufacturer

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List (initial)

- Shure
- Sennheiser
- Neutrik
- Phoenix Contact
- Switchcraft

## mountingType

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List (initial)

- 19 inch rack
- 23 inch rack
- DIN rail
- Panel mount - Flat
- Unknown

## equipmentType

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List (initial)

- Audio
- Video
- Mix
- Lighting
- Networking
- Patch Panel
- Unknown

## color

| Field        | Datatype     | Description          | Nullable |
| ------------ | ------------ | -------------------- | -------- |
| id           | int          | primary key          | No       |
| name         | varchar(50)  |                      | No       |
| abbreviation | varchar(3)   | Unique constraint    | No       |
| description  | varchar(500) | optional description | Yes      |

### Option List (initial)

- Red (RED)
- Orange (ORN)
- Yellow (YEL)
- Green (GRN)
- Blue (BLU)
- Purple (PUR)
- Brown (BRN)
- Black (BLK)
- Gray (GRY)
- Slate (SLT)
- Clear (CLR)
- Cyan (CYN)

## locationType

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List (initial)

- 19 inch rack
- 23 inch rack
- Backer board
- Enclosure

## connectorType

Represents connector types

| Field        | Datatype     | Foreign Key Link | Description                                | Nullable |
| ------------ | ------------ | ---------------- | ------------------------------------------ | -------- |
| id           | int          |                  | primary key                                | No       |
| name         | varchar(50)  |                  | friendly name                              | No       |
| description  | varchar(500) |                  | optional description                       | Yes      |
| gender       | int          | gender.id        |                                            | No       |
| direction    | int          | direction.id     |                                            | No       |
| numberPins   | int          |                  | number of pins                             | Yes      |
| manufacturer | int          | manufacturer.id  |                                            | Yes      |
| model        | varchar(50)  |                  | model number                               | Yes      |
| color        | int          | color.id         |                                            | Yes      |
| pluggable    | bit          |                  | If the connector is pluggable or hardwired | No       |

## connectorTypeMating

which connectors mate with which other connectors

<https://dba.stackexchange.com/a/48663/185504>
## equipment

Represents a piece of equipment.

If a piece of equipment has multiple faces, then they will be represented
separately in visual representations

| Field        | Datatype     | Foreign Key Link | Description                                 | Nullable |
| ------------ | ------------ | ---------------- | ------------------------------------------- | -------- |
| id           | int          |                  | primary key                                 | No       |
| name         | varchar(50)  |                  | friendly name                               | No       |
| description  | varchar(500) |                  | optional description                        | Yes      |
| width        | decimal(8,3) |                  | stored in mm                                | Yes      |
| height       | decimal(8,3) |                  | stored in mm, used for rack unit height     | Yes      |
| depth        | decimal(8,3) |                  | stored in mm                                | Yes      |
| diameter     | decimal(8,3) |                  | stored in mm                                | Yes      |
| manufacturer | int          | manufacturer.id  |                                             | Yes      |
| model        | varchar(50)  |                  | model number                                | Yes      |
| type         | int          | equipmentType.id |                                             | No       |
| numFaces     | int          |                  | number of faces with connectors             | No       |
| rackPercent  | bit          |                  | how much of a rack space is equipment width | Yes      |

## equipmentInst

Represents an instance of a piece of equipment. Also ties to location

MountingType linked here since a piece of equipment can have multiple potential mounting options.

| Field         | Datatype     | Foreign Key Link | Description                                               | Nullable |
| ------------- | ------------ | ---------------- | --------------------------------------------------------- | -------- |
| id            | int          |                  | primary key                                               | No       |
| name          | varchar(50)  |                  | friendly name                                             | No       |
| identifier    | varchar(50)  |                  | structured name                                           | Yes      |
| description   | varchar(500) |                  | optional description                                      | Yes      |
| mountingType  | int          | mountingType.id  |                                                           | No       |
| location      | int          | location.id      |                                                           | Yes      |
| rackUnit      | int          |                  | which rack unit in the location                           | Yes      |
| dinRailNumber | int          |                  | which DIN rail in panel the equipment is on               | Yes      |
| equipmentType | int          | equipment.id     |                                                           | No       |
| xPos          | decimal(8,3) |                  | stored in mm, distance from left side of location edge    | No       |
| yPos          | decimal(8,3) |                  | stored in mm, distance from bottom side of location edge  | No       |
| zPos          | decimal(8,3) |                  | stored in mm, distance from face of location (default 0)  | Yes      |

## equipmentConnector

Represents individual instance of connector on equipment.

| Field         | Datatype     | Foreign Key Link | Description                                               | Nullable |
| ------------- | ------------ | ---------------- | --------------------------------------------------------- | -------- |
| id            | int          |                  | primary key                                               | No       |
| name          | varchar(50)  |                  | friendly name                                             | No       |
| identifier    | varchar(50)  |                  | structured name                                           | Yes      |
| description   | varchar(500) |                  | optional description                                      | Yes      |
| xPos          | decimal(8,3) |                  | stored in mm, distance from left side of equipment edge   | No       |
| yPos          | decimal(8,3) |                  | stored in mm, distance from bottom side of equipment edge | No       |
| zPos          | decimal(8,3) |                  | stored in mm, distance from face of equipment (default 0) | Yes      |
| face          | int          |                  | which face of the equipment the connector is attached to  | No       |
| connectorType | int          | connectorType.id |                                                           | No       |
| equipment     | int          | equipmentInst.id |                                                           | No       |
| signalType    | int          | signalType.id    |                                                           | No       |


## location

represents locations where equipment can reside.

Each physical rack cabinet or panel would be a separate location

| Field            | Datatype      | Foreign Key Link | Description                                                 | Nullable |
| ---------------- | ------------- | ---------------- | ----------------------------------------------------------- | -------- |
| id               | int           |                  | primary key                                                 | No       |
| name             | varchar(50)   |                  | friendly name                                               | No       |
| identifier       | varchar(50)   |                  | structured name                                             | Yes      |
| description      | varchar(500)  |                  | optional description                                        | Yes      |
| locationType     | int           | locationType.id  |                                                             | Yes      |
| numDINRails      | int           |                  |                                                             | Yes      |
| dinRailPosns     | varchar(1000) |                  | contains encoded xyz, length coordinates of each DIN rail   | Yes      |
| width            | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| height           | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| depth            | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| availWidth       | decimal(8,3)  |                  | stored in mm, internal dimension                            | Yes      |
| availHeight      | decimal(8,3)  |                  | stored in mm, internal dimension, used for rack unit height | Yes      |
| availDepth       | decimal(8,3)  |                  | stored in mm, internal dimension                            | Yes      |
| physicalLocation | varchar(100)  |                  | street address, coordinates, description                    | Yes      |

## connection

represents a wire or cable with connectors.

An example of this is a wire or cable. Siamese cables are considered one connection.

| Field         | Datatype      | Foreign Key Link   | Description                                                  | Nullable |
| ------------- | ------------- | ------------------ | ------------------------------------------------------------ | -------- |
| id            | int           |                    | primary key                                                  | No       |
| name          | varchar(50)   |                    | friendly name                                                | No       |
| identifier    | varchar(50)   |                    | structured name                                              | Yes      |
| description   | varchar(500)  |                    | optional description                                         | Yes      |
| route         | int           | connectionRoute.id |                                                              | Yes      |
| length        | decimal(8,3)  |                    | stored in mm, nominal length of premade cable                | Yes      |
| physicalMedia | int           | physicalMedia.id   |                                                              | Yes      |
| numberOfEnds  | int           |                    | number of cable or wire ends that the connection consists of | Yes      |

## connectionRoute

represents a bundle of connections that are all routed along the same path

| Field         | Datatype      | Foreign Key Link | Description                                                 | Nullable |
| ------------- | ------------- | ---------------- | ----------------------------------------------------------- | -------- |
| id            | int           |                  | primary key                                                 | No       |
| name          | varchar(50)   |                  | friendly name                                               | No       |
| identifier    | varchar(50)   |                  | structured name                                             | Yes      |
| description   | varchar(500)  |                  | optional description                                        | Yes      |

## connectionConnector

represents the ends of a connection and the associated connectors

| Field         | Datatype      | Foreign Key Link | Description                                                   | Nullable |
| ------------- | ------------- | ---------------- | ------------------------------------------------------------- | -------- |
| id            | int           |                  | primary key                                                   | No       |
| name          | varchar(50)   |                  | friendly name                                                 | No       |
| identifier    | varchar(50)   |                  | structured name                                               | Yes      |
| description   | varchar(500)  |                  | optional description                                          | Yes      |
| connection    | int           | connection.id    |                                                               | No       |
| connectionEnd | int           |                  | which end of the connection this connector is associated with | No       |
| connectorType | int           | connectorType.id |                                                               | No       |


## equipmentConnectionLinks

represents relationship between equipmentConnectors and connectionConnectors

| Field               | Datatype      | Foreign Key Link       | Description                    | Nullable |
| ------------------- | ------------- | ---------------------- | ------------------------------ | -------- |
| id                  | int           |                        | primary key                    | No       |
| name                | varchar(50)   |                        | friendly name                  | No       |
| identifier          | varchar(50)   |                        | structured name                | Yes      |
| description         | varchar(500)  |                        | optional description           | Yes      |
| equipmentConnector  | int           | equipmentConnector.id  |                                | No       |
| connectionConnector | int           | connectionConnector.id |                                | No       |


