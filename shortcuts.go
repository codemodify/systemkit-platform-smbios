package smbios

import "strings"

// Info -
func Info() (SMBIOS, error) {
	rc, entryPoint, err := Stream()
	if err != nil {
		return SMBIOS{}, err
	}
	defer rc.Close()

	decoder := NewDecoder(rc)
	smBiosStructures, err := decoder.Decode()
	if err != nil {
		return SMBIOS{}, err
	}

	smBios := SMBIOS{
		EntryPoint: entryPoint,
		Memory:     []MemoryStick{},
		Battery:    []PortableBattery{},
	}

	for _, smBiosStructure := range smBiosStructures {
		setBios(&smBios, smBiosStructure)
		setSystemInformation(&smBios, smBiosStructure)
		setBaseboard(&smBios, smBiosStructure)
		setCPU(&smBios, smBiosStructure)
		setMemory(&smBios, smBiosStructure)
		setBattery(&smBios, smBiosStructure)
	}

	return smBios, nil
}

func setBios(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeBIOS {
		smBios.Bios = strings.Join(smBiosStructure.Strings, "/")

		if len(smBiosStructure.Strings) > 0 {
			smBios.BiosVendor = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			smBios.BiosVersion = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			smBios.BiosDate = smBiosStructure.Strings[2]
		}
	}
}

func setSystemInformation(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeSystemInformation {
		smBios.SystemInformation = strings.Join(smBiosStructure.Strings, "/")

		if len(smBiosStructure.Strings) > 0 {
			smBios.SystemInformationManufacturer = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			smBios.SystemInformationProductName = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			smBios.SystemInformationVersion = smBiosStructure.Strings[2]
		}
		if len(smBiosStructure.Strings) > 3 {
			smBios.SystemInformationSerialNumber = smBiosStructure.Strings[3]
		}
		if len(smBiosStructure.Strings) > 4 {
			smBios.SystemInformationSKU = smBiosStructure.Strings[4]
		}
		if len(smBiosStructure.Strings) > 5 {
			smBios.SystemInformationFamily = smBiosStructure.Strings[5]
		}
	}
}

func setBaseboard(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeBaseboard {
		smBios.Baseboard = strings.Join(smBiosStructure.Strings, "/")

		if len(smBiosStructure.Strings) > 0 {
			smBios.BaseboardManufacturer = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			smBios.BaseboardProduct = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			smBios.BaseboardVersion = smBiosStructure.Strings[2]
		}
		if len(smBiosStructure.Strings) > 3 {
			smBios.BaseboardSerialNumber = smBiosStructure.Strings[3]
		}
		if len(smBiosStructure.Strings) > 4 {
			smBios.BaseboardAssetTag = smBiosStructure.Strings[4]
		}
		if len(smBiosStructure.Strings) > 5 {
			smBios.BaseboardLocationInChassis = smBiosStructure.Strings[5]
		}
	}
}

func setCPU(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeCPU {
		smBios.CPU = strings.Join(smBiosStructure.Strings, "/")

		if len(smBiosStructure.Strings) > 0 {
			smBios.CPUSocketDesignation = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			smBios.CPUProcessorManufacturer = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			smBios.CPUProcessorVersion = smBiosStructure.Strings[2]
		}
		if len(smBiosStructure.Strings) > 3 {
			smBios.CPUSerialNumber = smBiosStructure.Strings[3]
		}
		if len(smBiosStructure.Strings) > 4 {
			smBios.CPUAssetTag = smBiosStructure.Strings[4]
		}
		if len(smBiosStructure.Strings) > 5 {
			smBios.CPUPartNumber = smBiosStructure.Strings[5]
		}
	}
}

func setMemory(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeMemoryStick {
		ms := MemoryStick{
			Memory: strings.Join(smBiosStructure.Strings, "/"),
		}

		if len(smBiosStructure.Strings) > 0 {
			ms.DeviceLocator = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			ms.BankLocator = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			ms.Manufacturer = smBiosStructure.Strings[2]
		}
		if len(smBiosStructure.Strings) > 3 {
			ms.SerialNumber = smBiosStructure.Strings[3]
		}
		if len(smBiosStructure.Strings) > 4 {
			ms.AssetTag = smBiosStructure.Strings[4]
		}
		if len(smBiosStructure.Strings) > 5 {
			ms.PartNumber = smBiosStructure.Strings[5]
		}

		smBios.Memory = append(smBios.Memory, ms)
	}
}

func setBattery(smBios *SMBIOS, smBiosStructure *Structure) {
	if smBiosStructure.Header.Type == StructureTypeBattery {
		pb := PortableBattery{
			Battery: strings.Join(smBiosStructure.Strings, "/"),
		}

		if len(smBiosStructure.Strings) > 0 {
			pb.Location = smBiosStructure.Strings[0]
		}
		if len(smBiosStructure.Strings) > 1 {
			pb.Manufacturer = smBiosStructure.Strings[1]
		}
		if len(smBiosStructure.Strings) > 2 {
			pb.ManufactureDate = smBiosStructure.Strings[2]
		}
		if len(smBiosStructure.Strings) > 3 {
			pb.SerialNumber = smBiosStructure.Strings[3]
		}
		if len(smBiosStructure.Strings) > 4 {
			pb.DeviceName = smBiosStructure.Strings[4]
		}
		if len(smBiosStructure.Strings) > 5 {
			pb.SBDSVersionNumber = smBiosStructure.Strings[5]
		}
		if len(smBiosStructure.Strings) > 6 {
			pb.SBDSDeviceChemistry = smBiosStructure.Strings[6]
		}

		smBios.Battery = append(smBios.Battery, pb)
	}
}
