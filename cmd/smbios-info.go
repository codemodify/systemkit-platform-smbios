package main

import (
	"encoding/json"
	"fmt"

	smbios "github.com/codemodify/systemkit-platform-smbios"
)

func main() {
	smbiosInfo, err := smbios.Info()
	if err != nil {
		fmt.Println(err.Error())
	}

	smbiosInfoAsJSON, err := json.MarshalIndent(&smbiosInfo, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(smbiosInfoAsJSON))
}
