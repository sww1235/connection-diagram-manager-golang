# File Formats

All file formats are currently based on YAML, (and inspired by
[WireViz](https://github.com/formatc1702/WireViz)).


Currently most of these pieces of information need to be in
one file. Eventually, this can be split into multiple files.
In preparation for this, (and also the possibility of
supporting alternate file syntax formats), YAML references are
not used.

```yaml
connector_types: #dictionary of all available connector types 

    <str>: # connectorType designator (must be unique)
	    
		# most tags are optional and will be filled with
		# blank values if unpopulated

		manufacturer: <str> # manufacturer name
		model: <str> # connector model description
		description: <str> # free text field for larger descriptions of connectors
		mounting_type: <str> # cable, pcb through hole, pcbsurface mount, panel
		panel_cutout: <str> # D, A, etc. optional
		pn: <str> # [internal] part number
		mpn: <str> # manufacturer part number
		supplier: <str> # supplier
		spn: <str> # supplier part number

		gender: <str> # (male, female, rpmale, rpfemale, hermaphroditic, unknown)

		# pinout information
		# at least one of the following must be specified
		pincount: <int>    # if omitted, is set to length of specified list(s)
        pins: <List>       # if omitted, is autofilled with [1, 2, ..., pincount]
        pinlabels: <List>  # if omitted, is autofilled with blanks
        # pin color marks (optional)
		pincolors: <List> # list of colors to be assigned
                          # goes in order of pin count/pin list
						  # if fewer colors are specified than pins, end of list will have no colors specified
		pin_signal_type: <list>

		# all images are specified as SVG images so they scale.
		visrep: <svg> # svg text (ideally minimized), can be multiline
		pin_visrep: <svg> # svg text, shows pin layout of connector

equipment_types: # dictionary of all available equipment types
    <str>: # equipmentType designator (must be unique)
	    
		# most tags are optional and will be filled with
		# blank values if unpopulated

		manufacturer: <str>
		model: <str>
		mounting_type: <str> # (19" rack, 23" rack, 1/2 19" rack, DIN rail, wall mount)
		type: <str> # (audio, video, mix, lighting, networking, patch panel)
		faces: # dictionary of faces that can have connectors associated with them, and an associated visual representation. should not include connectors themselves.
		    <str>: <svg>
			<str>: <svg>
		visrep: <svg> # overall visual representation of equipment

		connectors: # dictionary of connectors on equipment

		    <str>: # connectorType identifier
                
				id: <str> # identifier of connector on equipment
		        direction: <str> # (input, output, power input, power output, bidirectional, passive)
				face: <str> # face connector is located on
                x: <integer> # location of connector from bottom left of visrep of face to right
				y: <integer> # location of connector from bottom left of visrep of face up


		    



```
