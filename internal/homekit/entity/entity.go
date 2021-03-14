package entity

type BaseEntity struct {
	Name            string
	SerialNumber    string
	FirmwareVersion string
	Type            string
	Options         map[string]interface{}
}

type BaseEntities = []BaseEntity
