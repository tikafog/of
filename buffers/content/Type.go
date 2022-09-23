// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package content

import "strconv"

type Type uint32

const (
	TypeNone              Type = 0
	TypeCore              Type = 1
	TypeUser              Type = 2
	TypeBootstrap         Type = 3
	TypeInformation       Type = 4
	TypeDecryptedInstruct Type = 5
	TypeUpdate            Type = 6
	TypeAnnounce          Type = 7
	TypeTopList           Type = 8
	TypeDiscovery         Type = 9
	TypePage              Type = 10
	TypeMax               Type = 11
)

var EnumNamesType = map[Type]string{
	TypeNone:              "None",
	TypeCore:              "Core",
	TypeUser:              "User",
	TypeBootstrap:         "Bootstrap",
	TypeInformation:       "Information",
	TypeDecryptedInstruct: "DecryptedInstruct",
	TypeUpdate:            "Update",
	TypeAnnounce:          "Announce",
	TypeTopList:           "TopList",
	TypeDiscovery:         "Discovery",
	TypePage:              "Page",
	TypeMax:               "Max",
}

var EnumValuesType = map[string]Type{
	"None":              TypeNone,
	"Core":              TypeCore,
	"User":              TypeUser,
	"Bootstrap":         TypeBootstrap,
	"Information":       TypeInformation,
	"DecryptedInstruct": TypeDecryptedInstruct,
	"Update":            TypeUpdate,
	"Announce":          TypeAnnounce,
	"TopList":           TypeTopList,
	"Discovery":         TypeDiscovery,
	"Page":              TypePage,
	"Max":               TypeMax,
}

func (v Type) String() string {
	if s, ok := EnumNamesType[v]; ok {
		return s
	}
	return "Type(" + strconv.FormatInt(int64(v), 10) + ")"
}
