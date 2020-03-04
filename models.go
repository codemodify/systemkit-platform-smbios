package smbios

const (
	StructureTypeBIOS              = 0
	StructureTypeSystemInformation = 1
	StructureTypeBaseboard         = 2
	StructureTypeCPU               = 4
	StructureTypeMemoryStick       = 17
	StructureTypeBattery           = 22
)

// A Header is a Structure's header.
type Header struct {
	Type   uint8
	Length uint8
	Handle uint16
}

// A Structure is an SMBIOS structure.
type Structure struct {
	Header    Header
	Formatted []byte
	Strings   []string
}

type SMBIOS struct {
	EntryPoint EntryPoint

	Bios        string
	BiosVendor  string
	BiosVersion string
	BiosDate    string

	SystemInformation             string
	SystemInformationManufacturer string
	SystemInformationProductName  string
	SystemInformationVersion      string
	SystemInformationSerialNumber string
	SystemInformationSKU          string
	SystemInformationFamily       string

	Baseboard                  string
	BaseboardManufacturer      string
	BaseboardProduct           string
	BaseboardVersion           string
	BaseboardSerialNumber      string
	BaseboardAssetTag          string
	BaseboardLocationInChassis string

	CPU                      string
	CPUSocketDesignation     string
	CPUProcessorManufacturer string
	CPUProcessorVersion      string
	CPUSerialNumber          string // Processor serial numbers were basically only in Pentium III processors. Intel removed it from later models due to the privacy concerns that were raised
	CPUAssetTag              string
	CPUPartNumber            string

	Memory []MemoryStick

	Battery []PortableBattery
}

type MemoryStick struct {
	Memory        string
	DeviceLocator string
	BankLocator   string
	Manufacturer  string
	SerialNumber  string
	AssetTag      string
	PartNumber    string
}

type PortableBattery struct {
	Battery             string
	Location            string
	Manufacturer        string
	ManufactureDate     string
	SerialNumber        string
	DeviceName          string
	SBDSVersionNumber   string
	SBDSDeviceChemistry string
}
