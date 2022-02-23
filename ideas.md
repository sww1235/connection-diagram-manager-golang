# Ideas
used for representing static configurations of interconnected equipment.

Can create multiple configurations of the same equipment with different
connections.

## Misc Ideas

somehow provide checks to prevent weird interconnections. Probably have to do it
in connection object.

may want concept of route that connections can be bundled into

Want to have wire, cable, route

## Connection
concept of "connection" is an interconnection between two or more connectors.

connection can have:
-   length
-   cable type
-   type (temporary, permanent, virtual)

connector types are assigned by which connector the connection is connected to.

connection contains logic to validate IO direction and gender

Optionally contains pin mapping data for breakout cables

## Equipment

concept of "equipment/gear" which is a single unit with no representable
internal connections. (may add the ability for internal connections between
external connections, for patch bay or switch type things. Not reallly meant for
documenting changing routing.) examples would be a monitor, a audio mixing desk,
a network switch etc.

Equipment can have:
-   connector
-   mounting
-   manufacturer
-   model
-   type (audio, video, mix, lighting, networking, patch panel)
- 	location

patchbay type has internal connections between ports.

Equipment does not represent a structure like a rack or panel

## ConnectorType

concept of "connector" which is a physical connection on a piece of equipment.
examples are XLR, DVI, BNC, IEC, etc.

connector can have
-   gender (male, female, rpmale, rpfemale, hermaphroditic, unknown)
-   direction (input, outlet, power input, power outlet, bidirectional)
-	Number of pins
- 	Manufacturer
-	Model
-	SignalType: to allow some differentiation and connection logic.

Connectors are linked to equipment via a table that establishes individual
instances of connectors, and their relationships

## Location

represents where equipment resides. Does not represent an address or building.
