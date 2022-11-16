package main

//TODO: remove debug lines that use reflect

import (
	"reflect"
)

func ParseTermCableTypes(dictValue map[string]interface{}) error {
	// each instance is unique key of TermCableType dictionary Entry
	for instance, content := range dictValue {
		// checking if instance is already present in map
		if _, ok := TermCableTypeDict[instance]; ok {
			// TODO: prompt user to merge two instances
			infoLogger.Printf("TermCableInstance %s already present in dictionary, skipping.\n", instance)
			continue //skip importing duplicate value
		}
		var tempTermCableType TermCableType
		debugLogger.Printf("instance: %s, value: %s,  type of value: %s\n", instance, content, reflect.TypeOf(content))
		// Type assert checking that content of instance is map[string]interface{}
		contentValue, ok := content.(map[string]interface{})
		if !ok {
			debugLogger.Printf("not map. key: %s, type of value: %s\n", instance, reflect.TypeOf(content))
		}
		for parameter, value := range contentValue {

			switch parameter {
			case "manufacturer":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Manufacturer of %s is not a string.\n", instance)
				}

				tempTermCableType.Manufacturer = value
			case "pn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Pn of %s is not a string.\n", instance)
				}

				tempTermCableType.Pn = value
			case "mpn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Mpn of %s is not a string.\n", instance)
				}

				tempTermCableType.Mpn = value
			case "supplier":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Supplier of %s is not a string.\n", instance)
				}

				tempTermCableType.Supplier = value
			case "spn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("SPN of %s is not a string.\n", instance)
				}

				tempTermCableType.Spn = value
			case "cable":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("CableType ID of %s is not a string.\n", instance)
				}
				CableTypeValue, present := cableTypeDict[value]
				if !present {
					// if not present, insert an empty CableType and warn the user
					var tempCableType CableType
					cableTypeDict[value] = tempCableType
					CableTypeValue = cableTypeDict[value]
					debugLogger.Printf("CableType %s not present in imported files thus far.\n", value)

				}
				tempTermCableType.Cable = CableTypeValue

			case "wire":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("WireType ID of %s is not a string.\n", instance)
				}
				WireTypeValue, present := wireTypeDict[value]
				if !present {
					// if not present, insert an empty WireType and warn the user.
					// The empty type may be merged with another imported instance later
					var tempWireType WireType
					wireTypeDict[value] = tempWireType
					WireTypeValue = wireTypeDict[value]
					debugLogger.Printf("WireType %s not present in imported files thus far.\n", value)

				}
				tempTermCableType.Wire = WireTypeValue

			case "nom_length":
				floatValue, floatOk := value.(float64)
				intValue, intOk := value.(int)
				//TODO: prompt to fix
				if !floatOk && !intOk {
					infoLogger.Printf("Nominal Length of %s is not a number.\n", instance)
					debugLogger.Printf("Type of value: %s.\n", reflect.TypeOf(value))
				}
				if floatOk {
					tempTermCableType.NomLength = floatValue
				} else if intOk {

					tempTermCableType.NomLength = float64(intValue)
				}
			case "length":
				floatValue, floatOk := value.(float64)
				intValue, intOk := value.(int)
				//TODO: prompt to fix
				if !floatOk && !intOk {
					infoLogger.Printf("Length of %s is not a number.\n", instance)
					debugLogger.Printf("Type of value: %s\n.", reflect.TypeOf(value))
				}
				if floatOk {
					tempTermCableType.Length = floatValue
				} else if intOk {

					tempTermCableType.Length = float64(intValue)
				}
			case "end1":
				_, ok := value.(map[string]interface{})
				if ok {
					infoLogger.Printf("End1 of %s is a map. It needs to be a dictionary.\n", instance)
					debugLogger.Printf("value: %s, typeof value: %s\n", value, reflect.TypeOf(value))
					//TODO: return error here
				}

				//create TermCableConnector Struct
				var tempTermCableConnectors []TermCableConnector
				tempTermCableType.End1 = tempTermCableConnectors

				// loop through all connectors in end1 slice
				// WTF, why can I iterate over a []interface{} when I type assert, but not directly?
				for _, connector := range value.([]interface{}) {
					connector, ok := connector.(map[string]interface{})
					if !ok {
						infoLogger.Printf("Connector %s in end1 is not a map[string]interface{}\n", connector)
					}
					var tempTermCableConnector TermCableConnector
					for endParameter, contents := range connector {
						switch endParameter {

						case "type":
							value, ok := contents.(string)
							if !ok {
								infoLogger.Printf("end1 connectorType not a string\n")
							}
							ConnectorTypeValue, present := connectorTypeDict[value]
							if !present {
								// if not present, insert an empty ConnectorType and
								// warn the user.
								// The empty type may be merged with another imported
								// instance later
								var tempConnectorType ConnectorType
								connectorTypeDict[value] = tempConnectorType
								ConnectorTypeValue = connectorTypeDict[value]
								debugLogger.Printf("ConnectorType %s not present in imported files thus far.\n", value)
							}
							// need to append here
							tempTermCableConnector.Type = ConnectorTypeValue

						case "termination":

						}
					}

					tempTermCableType.End1 = append(tempTermCableType.End1, tempTermCableConnector)

				}

			case "end2":
				_, ok := value.(map[string]interface{})
				if ok {
					infoLogger.Printf("End2 of %s is a map. It needs to be a dictionary.\n", instance)
					debugLogger.Printf("value: %s, typeof value: %s\n", value, reflect.TypeOf(value))
					//TODO: return error here
				}

				//create TermCableConnector Struct
				var tempTermCableConnectors []TermCableConnector
				tempTermCableType.End2 = tempTermCableConnectors

				// loop through all connectors in end2 slice
				// WTF, why can I iterate over a []interface{} when I type assert, but not directly?
				for _, connector := range value.([]interface{}) {
					connector, ok := connector.(map[string]interface{})
					if !ok {
						infoLogger.Printf("Connector %s in end1 is not a map[string]interface{}\n", connector)
					}
					var tempTermCableConnector TermCableConnector
					for endParameter, contents := range connector {
						switch endParameter {

						case "type":
							value, ok := contents.(string)
							if !ok {
								infoLogger.Printf("end2 connectorType not a string\n")
							}
							ConnectorTypeValue, present := connectorTypeDict[value]
							if !present {
								// if not present, insert an empty ConnectorType and
								// warn the user.
								// The empty type may be merged with another imported
								// instance later
								var tempConnectorType ConnectorType
								connectorTypeDict[value] = tempConnectorType
								ConnectorTypeValue = connectorTypeDict[value]
								debugLogger.Printf("ConnectorType %s not present in imported files thus far.\n", value)
							}
							// need to append here
							tempTermCableConnector.Type = ConnectorTypeValue

						case "termination":

						}
					}

					tempTermCableType.End2 = append(tempTermCableType.End2, tempTermCableConnector)

				}
			default:
				// TODO: maybe throw an error here? Need to somehow print which file it is from.
				infoLogger.Printf("parameter in TermCableType %s is not recognized, please fix, attempting to continue gracefully", parameter)

			}
		}

		TermCableTypeDict[instance] = tempTermCableType
	}
	return nil
}

func ParseLocationTypes(dictValue map[string]interface{}) error {
	// each instance is unique key of LocationType dictionary entry
	for instance, content := range dictValue {
		if _, ok := LocationTypeDict[instance]; ok {
			//TODO: prompt user to merge two instances
			infoLogger.Printf("LocationTypeInstance %s already present in dictionary, skipping.\n", instance)
			continue //skip importing duplicate value
		}
		var tempLocationType LocationType
		debugLogger.Printf("instance: %s, value: %s, type of value: %s\n", instance, content, reflect.TypeOf(content))
		// Type assert checking that content of instance is map[string]interface{}
		contentValue, ok := content.(map[string]interface{})
		if !ok {
			debugLogger.Printf("not map. key: %s, type of value: %s\n", instance, reflect.TypeOf(content))
		}
		for parameter, value := range contentValue {
			switch parameter {
			case "manufacturer":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Manufacturer of %s is not a string.\n", instance)
				}

				tempLocationType.Manufacturer = value
			case "pn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Pn of %s is not a string.\n", instance)
				}

				tempLocationType.Pn = value
			case "mpn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Mpn of %s is not a string.\n", instance)
				}

				tempLocationType.Mpn = value
			case "supplier":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Supplier of %s is not a string.\n", instance)
				}

				tempLocationType.Supplier = value
			case "spn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("SPN of %s is not a string.\n", instance)
				}

				tempLocationType.Spn = value
			}
		}

	}
	return nil
}

func ParseConnectorTypes(dictValue map[string]interface{}) error {
	for instance, content := range dictValue {
		// checking if instance is already present in map
		if _, ok := TermCableTypeDict[instance]; ok {
			// TODO: prompt user to merge two instances
			infoLogger.Printf("TermCableInstance %s already present in dictionary, skipping.\n", instance)
			continue //skip importing duplicate value
		}
		var tempConnectorType ConnectorType
		debugLogger.Printf("instance: %s, value: %s,  type of value: %s\n", instance, content, reflect.TypeOf(content))
		// Type assert checking that content of instance is map[string]interface{}
		contentValue, ok := content.(map[string]interface{})
		if !ok {
			debugLogger.Printf("not map. key: %s, type of value: %s\n", instance, reflect.TypeOf(content))
		}
		for parameter, value := range contentValue {
			//Description   string   `yaml:"description"`
			//MountingType  []string `yaml:"mounting_type"`
			//PanelCutout   string   `yaml:"panel_cutout"`
			//Pn            string   `yaml:"pn"`
			//Mpn           string   `yaml:"mpn"`
			//Supplier      string   `yaml:"supplier"`
			//Spn           string   `yaml:"spn"`
			//Gender        string   `yaml:"gender"`
			//Height        float64  `yaml:"height"`
			//Width         float64  `yaml:"width"`
			//Depth         float64  `yaml:"depth"`
			//Diameter      float64  `yaml:"diameter"`
			//Pincount      int      `yaml:"pincount"`
			//Pins          []string `yaml:"pins"`
			//Pinlabels     []string `yaml:"pinlabels"`
			//Pincolors     []string `yaml:"pincolors"`
			//PinSignalType []string `yaml:"pin_signal_type"`
			//Visrep        SVG      `yaml:"visrep"`
			//PinVisrep     SVG      `yaml:"pin_visrep"`
			switch parameter {
			case "manufacturer":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Manufacturer of %s is not a string.\n", instance)
				}

				tempConnectorType.Manufacturer = value
			case "model":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Model of %s is not a string.\n", instance)
				}

				tempConnectorType.Model = value
			case "pn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Pn of %s is not a string.\n", instance)
				}

				tempConnectorType.Pn = value
			case "mpn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Mpn of %s is not a string.\n", instance)
				}

				tempConnectorType.Mpn = value
			case "supplier":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("Supplier of %s is not a string.\n", instance)
				}

				tempConnectorType.Supplier = value
			case "spn":
				value, ok := value.(string)
				//TODO: prompt to fix
				if !ok {
					infoLogger.Printf("SPN of %s is not a string.\n", instance)
				}

				tempConnectorType.Spn = value

			}

		}
	}

	return nil
}

func ParseConnectorTypeMates(dictValue map[string]interface{}) error {
	//for instance, content := range dictValue {
	//	debugLogger.Printf("instance: %s, type of value: %s", instance, reflect.TypeOf(content))
	//}
	return nil
}

func ParseEquipmentTypes(dictValue map[string]interface{}) error {
	//for instance, content := range dictValue {
	//	debugLogger.Printf("instance: %s, type of value: %s", instance, reflect.TypeOf(content))
	//}
	return nil
}

func ParsePathwayTypes(dictValue map[string]interface{}) error {
	//for instance, content := range dictValue {
	//	debugLogger.Printf("instance: %s, type of value: %s", instance, reflect.TypeOf(content))
	//}
	return nil
}

func ParseWireTypes(dictValue map[string]interface{}) error {
	//for instance, content := range dictValue {
	//	debugLogger.Printf("instance: %s, type of value: %s", instance, reflect.TypeOf(content))
	//}
	return nil
}

func ParseCableTypes(dictValue map[string]interface{}) error {
	//for instance, content := range dictValue {
	//	debugLogger.Printf("instance: %s, type of value: %s", instance, reflect.TypeOf(content))
	//}
	return nil
}
