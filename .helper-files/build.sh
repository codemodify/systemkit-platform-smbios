#!/bin/bash

SUPPORTED_TARGETS=(
  "darwin,amd64"
  "linux,386"
  "linux,amd64"
  "linux,arm"
  "linux,arm64"
  "windows,amd64"
)

OUTPUTFOLDER=temp-build

clear
echo ""
echo "Build PARAMS"
echo "    -> OUTPUTFOLDER             : ${OUTPUTFOLDER}"
echo ""

for target in "${SUPPORTED_TARGETS[@]}"; do

  # get the OS/Arch array
  IFS=',' read -ra targ <<<"$target"
  os="${targ[0]}"
  arch="${targ[1]}"
  ext=""

  if [ $os = "windows" ]; then
    ext=".exe"
  fi

  printf "Building %-22s" ${os}.${arch}
  echo ""

  fullOutputFilePath="$OUTPUTFOLDER/systemkit-platform-smbios.${os}.${arch}${ext}"
  GOOS=${os} GOARCH=${arch} go build -ldflags "${LDFLAGS}" -o ${fullOutputFilePath} ../cmd

done

echo ""

SUPPORTED_TARGETS_NOT=(
  "darwin,386"
  "linux,ppc64"
  "linux,ppc64le"
  "linux,mips"
  "linux,mipsle"
  "linux,mips64,true"
  "linux,mips64le,true"
  "linux,s390x,true"
  "windows,386"
  "aix,ppc64"
  "android,386"
  "android,amd64"
  "android,arm"
  "android,arm64"
  "darwin,arm"
  "darwin,arm64"
  "dragonfly,amd64"
  "freebsd,386,true"
  "freebsd,amd64,true"
  "freebsd,arm,true"
  "illumos,amd64"
  "js,wasm"
  "netbsd,386,true"
  "netbsd,amd64,true"
  "netbsd,arm,true"
  "openbsd,386,true"
  "openbsd,amd64,true"
  "openbsd,arm,true"
  "openbsd,arm64,true"
  "plan9,386"
  "plan9,amd64"
  "plan9,arm"
  "solaris,amd64"
)
