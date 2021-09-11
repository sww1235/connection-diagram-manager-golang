# Database Schema

## connector\_gender

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

### Option List
- Male
- Female
- RP-Male
- RP-Female
- Hermaphroditic
- Unknown

### physical\_media

| Field       | Datatype     | Description          | Nullable |
| ----------- | ------------ | -------------------- | -------- |
| id          | int          | primary key          | No       |
| name        | varchar(50)  |                      | No       |
| description | varchar(500) | optional description | Yes      |

# direction

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

## signal\_type

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

## mounting\_type

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

## equipment\_type

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

## location\_type

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

## connector\_model

Represents connector models

| Field        | Datatype     | Foreign Key Link | Description                                | Nullable |
| ------------ | ------------ | ---------------- | ------------------------------------------ | -------- |
| id           | int          |                  | primary key                                | No       |
| name         | varchar(50)  |                  | friendly name                              | No       |
| description  | varchar(500) |                  | optional description                       | Yes      |
| gender       | int          | gender.id        |                                            | No       |
| direction    | int          | direction.id     |                                            | No       |
| number\_of\_pins   | int          |                  | number of pins                             | Yes      |
| manufacturer | int          | manufacturer.id  |                                            | Yes      |
| model        | varchar(50)  |                  | model number                               | Yes      |
| color        | int          | color.id         |                                            | Yes      |
| pluggable    | bit          |                  | If the connector is pluggable or hardwired | No       |

## connector\_model\_mating

which connectors mate with which other connectors

<https://dba.stackexchange.com/a/48663/185504>
## equipmentModel

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
| type         | int          | equipment\_type.id |                                             | No       |
| number\_faces     | int          |                  | number of faces with connectors             | No       |
| rack\_percent  | int          |                  | how much of a rack space is equipment width | Yes      |

## equipment\_inst

Represents an instance of a piece of equipment. Also ties to location

MountingType linked here since a piece of equipment can have multiple potential mounting options.

| Field         | Datatype     | Foreign Key Link | Description                                               | Nullable |
| ------------- | ------------ | ---------------- | --------------------------------------------------------- | -------- |
| id            | int          |                  | primary key                                               | No       |
| name          | varchar(50)  |                  | friendly name                                             | No       |
| identifier    | varchar(50)  |                  | structured name                                           | Yes      |
| description   | varchar(500) |                  | optional description                                      | Yes      |
| mounting\_type  | int          | mounting\_type.id  |                                                           | No       |
| location      | int          | location.id      |                                                           | Yes      |
| rackUnit      | int          |                  | which rack unit in the location                           | Yes      |
| din\_rail\_number | int          |                  | which DIN rail in panel the equipment is on               | Yes      |
| equipment\_type | int          | equipment.id     |                                                           | No       |
| x\_pos          | decimal(8,3) |                  | stored in mm, distance from left side of location edge    | No       |
| y\_pos          | decimal(8,3) |                  | stored in mm, distance from bottom side of location edge  | No       |
| z\_pos          | decimal(8,3) |                  | stored in mm, distance from face of location (default 0)  | Yes      |

## equipment\_connector

Represents individual instance of connector on equipment.

| Field         | Datatype     | Foreign Key Link | Description                                               | Nullable |
| ------------- | ------------ | ---------------- | --------------------------------------------------------- | -------- |
| id            | int          |                  | primary key                                               | No       |
| name          | varchar(50)  |                  | friendly name                                             | No       |
| identifier    | varchar(50)  |                  | structured name                                           | Yes      |
| description   | varchar(500) |                  | optional description                                      | Yes      |
| x\_pos          | decimal(8,3) |                  | stored in mm, distance from left side of equipment edge   | No       |
| y\_pos          | decimal(8,3) |                  | stored in mm, distance from bottom side of equipment edge | No       |
| z\_pos          | decimal(8,3) |                  | stored in mm, distance from face of equipment (default 0) | No      |
| face          | int          |                  | which face of the equipment the connector is attached to  | No       |
| connector\_type | int          | connector\_type.id |                                                           | No       |
| equipment     | int          | equipment\_tnst.id |                                                           | No       |
| signal\_type    | int          | signal\_type.id    |                                                           | No       |


## location

represents locations where equipment can reside.

Each physical rack cabinet or panel would be a separate location

| Field            | Datatype      | Foreign Key Link | Description                                                 | Nullable |
| ---------------- | ------------- | ---------------- | ----------------------------------------------------------- | -------- |
| id               | int           |                  | primary key                                                 | No       |
| name             | varchar(50)   |                  | friendly name                                               | No       |
| identifier       | varchar(50)   |                  | structured name                                             | Yes      |
| description      | varchar(500)  |                  | optional description                                        | Yes      |
| location\_type     | int           | location\_type.id  |                                                             | Yes      |
| number\_din\_rails      | int           |                  |                                                             | Yes      |
| din\_rail\_posns     | varchar(1000) |                  | contains encoded xyz, length coordinates of each DIN rail   | Yes      |
| width            | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| height           | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| depth            | decimal(8,3)  |                  | stored in mm, external dimension                            | Yes      |
| avail\_width       | decimal(8,3)  |                  | stored in mm, internal dimension                            | Yes      |
| avail\_height      | decimal(8,3)  |                  | stored in mm, internal dimension, used for rack unit height | Yes      |
| avail\_depth       | decimal(8,3)  |                  | stored in mm, internal dimension                            | Yes      |
| physical\_location | varchar(100)  |                  | street address, coordinates, description                    | Yes      |

## connection

represents a wire or cable with connectors.

An example of this is a wire or cable. Siamese cables are considered one connection.

| Field         | Datatype      | Foreign Key Link   | Description                                                  | Nullable |
| ------------- | ------------- | ------------------ | ------------------------------------------------------------ | -------- |
| id            | int           |                    | primary key                                                  | No       |
| name          | varchar(50)   |                    | friendly name                                                | No       |
| identifier    | varchar(50)   |                    | structured name                                              | Yes      |
| description   | varchar(500)  |                    | optional description                                         | Yes      |
| route         | int           | connection\_route.id |                                                              | Yes      |
| length        | decimal(8,3)  |                    | stored in mm, nominal length of premade cable                | Yes      |
| physical\_media | int           | physical\_media.id   |                                                              | Yes      |
| number\_of\_ends  | int           |                    | number of cable or wire ends that the connection consists of | Yes      |

## connection\_route

represents a bundle of connections that are all routed along the same path

| Field         | Datatype      | Foreign Key Link | Description                                                 | Nullable |
| ------------- | ------------- | ---------------- | ----------------------------------------------------------- | -------- |
| id            | int           |                  | primary key                                                 | No       |
| name          | varchar(50)   |                  | friendly name                                               | No       |
| identifier    | varchar(50)   |                  | structured name                                             | Yes      |
| description   | varchar(500)  |                  | optional description                                        | Yes      |

## connection\_connector

represents the ends of a connection and the associated connectors

| Field         | Datatype      | Foreign Key Link | Description                                                   | Nullable |
| ------------- | ------------- | ---------------- | ------------------------------------------------------------- | -------- |
| id            | int           |                  | primary key                                                   | No       |
| name          | varchar(50)   |                  | friendly name                                                 | No       |
| identifier    | varchar(50)   |                  | structured name                                               | Yes      |
| description   | varchar(500)  |                  | optional description                                          | Yes      |
| connection    | int           | connection.id    |                                                               | No       |
| connection\_end | int           |                  | which end of the connection this connector is associated with | No       |
| connector\_model | int           | connector\_model.id |                                                               | No       |


## equipment\_connection\_links

represents relationship between equipmentConnectors and connectionConnectors

| Field               | Datatype      | Foreign Key Link       | Description                    | Nullable |
| ------------------- | ------------- | ---------------------- | ------------------------------ | -------- |
| id                  | int           |                        | primary key                    | No       |
| name                | varchar(50)   |                        | friendly name                  | No       |
| identifier          | varchar(50)   |                        | structured name                | Yes      |
| description         | varchar(500)  |                        | optional description           | Yes      |
| equipment\_connector  | int           | equipment\_connector.id  |                                | No       |
| connection\_connector | int           | connection\_connector.id |                                | No       |


