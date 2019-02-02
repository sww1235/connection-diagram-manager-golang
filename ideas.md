# Ideas
used for representing static configurations of interconnected equipment.

Can create multiple configurations of the same equipment with different
connections.

## Misc Ideas

somehow provide checks to prevent weird interconnections. Probably have to do it
in connection object.

may want concept of route that connections can be bundled into

## Connection
concept of "connection" is an interconnection between two connectors.

connection can have:
-   length
-   cable type
-   type (temporary, permanent, virtual)

connector types are assigned by which connector the connection is connected to.

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

patchbay type has internal connections between ports.

## Connector

concept of "connector" which is a physical connection on a piece of equipment.
examples are XLR, DVI, BNC, IEC, etc.

may need to implement subconector / port for creating split cables

connector can have
-   gender (male, female, rpmale, rpfemale, hermaphroditic, unknown)
-   direction (input, outlet, power input, power outlet, bidirectional)
-   type (see below)

## Connector type

this is a unique representation of number of pins, type of connector and type of
pins. May also hold subconnectors as well.
