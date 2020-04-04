# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Access SMBIOS/DMI information exposed by hardware
[![](https://img.shields.io/github/v/release/codemodify/systemkit-platform-smbios?style=flat-square)](https://github.com/codemodify/systemkit-platform-smbios/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-platform-smbios?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-platform-smbios?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-platform-smbios/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-platform-smbios?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-platform-smbios?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-platform-smbios)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-platform-smbios)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-platform-smbios?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-platform-smbios?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-platform-smbios?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-platform-smbios?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-platform-smbios?style=flat-square)

#### Access SMBIOS/DMI information exposed by hardware
#### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-config
```



# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Specs
- https://www.dmtf.org/sites/default/files/standards/documents/DSP0134_3.1.1.pdf

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Refrences
- https://docs.rs/smbios/0.1.7/smbios/
- https://blog.gopheracademy.com/advent-2017/accessing-smbios-information-with-go
- https://github.com/codemodify/systemkit-platform-smbios
- https://github.com/jaypipes/ghw
- https://github.com/shirou/gopsutil
- https://github.com/ochapman/godmi
- https://github.com/intel-go/cpuid
- https://github.com/ebfe?tab=repositories


# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Dox
- "W1KS96C10B3" id produced by
	- `sudo cat /sys/class/dmi/id/board_serial`
	- `sudo dmidecode -s baseboard-serial-number`
- "89b0774c-2975-11b2-a85c-ef1c041b9a68" id produced by
	- `sudo cat /sys/class/dmi/id/product_uuid`
		- There is no such thing Linux PPC (Apple based)
	- `sudo dmidecode -s system-uuid`
	- `sudo dmidecode | grep -i uuid`
- "61fabe084dc648e7b70eda16cd771e79" id produced by
	- `cat /var/lib/dbus/machine-id`

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Debug
- `sudo dlv debug --headless --listen=:2345 --log --api-version=2 --`

```text
sudo ./ghwc baseboard 	-> baseboard vendor=LENOVO serial=W1KS96C10B3 version=Not Defined
sudo ./ghwc bios		-> bios vendor=LENOVO version=N2EET43W (1.25 ) date=10/28/2019
sudo ./ghwc block		-> /dev/nvme0n1 SSD (954GB) NVMe [@pci-0000:71:00.0-nvme-1 (node #0)] vendor=unknown model=SAMSUNG MZVLB1T0HALR-000L7 serial=S3TPNX0M522320 WWN=eui.0025388591be5238
sudo ./ghwc chassis		-> chassis type=Notebook vendor=LENOVO serial=R90VCAEW version=None
sudo ./ghwc cpu
sudo ./ghwc gpu
sudo ./ghwc memory
sudo ./ghwc net
sudo ./ghwc product		-> product family=ThinkPad X1 Extreme name=20MF000DUS vendor=LENOVO serial=R90VCAEW uuid=89b0774c-2975-11b2-a85c-ef1c041b9a68 sku=LENOVO_MT_20MF_BU_Think_FM_ThinkPad X1 Extreme version=ThinkPad X1 Extreme
sudo ./ghwc topology
```