package main

import (
    "os"
    "fmt"
    "flag"

    "github.com/cciva/pkflux/build"
)

const usage = `Usage:
    pkflux [--enable] [--disable] [--sync]
    pkflux [--version]
    pkflux [--help]

Options:
    -e, --enable               	Enables the firewall.
    -d, --disable               Disables the firewall.
    -u, --use                   Use/activate profile.
    -s, --sync                  Syncs user profile changes.
    -v, --version               Prints app version.
    -h, --help                  Prints help.

Example:
    $ pkflux --sync
`

var cmdHelp = flag.Bool("help", false, "Prints help.")
var cmdVersion = flag.Bool("version", false, "Prints version.")
var cmdEnable = flag.Bool("enable", false, "Enables the firewall.")
var cmdDisable = flag.Bool("disable", false, "Disables the firewall.")
var cmdUse = flag.Bool("use", false, "Use/activate profile.")
var cmdSync = flag.Bool("sync", false, "Syncs user profile changes.")

var version build.VersionInfo

func main() {
    flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s\n", usage)
	}

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}

    flag.BoolVar(cmdHelp, "n", false, "Prints help.")
    flag.BoolVar(cmdVersion, "v", false, "Prints version.")
    flag.BoolVar(cmdEnable, "e", false, "Enables the firewall.")
    flag.BoolVar(cmdDisable, "d", false, "Disables the firewall.")
    flag.BoolVar(cmdUse, "u", false, "Use/activate profile.")
    flag.BoolVar(cmdSync, "s", false, "Syncs user profile changes.")
    flag.Parse()

    switch {
    case *cmdVersion:
        build.GetVersion(&version)
		fmt.Fprintf(os.Stdout, "Version: %s\n", version.Version)
		fmt.Fprintf(os.Stdout, "Timestamp: %s\n", version.Timestamp)
        return
    case *cmdEnable:
        fmt.Fprintf(os.Stdout, "enabling...\n")
    case *cmdDisable:
        fmt.Fprintf(os.Stdout, "disabling...\n")
    case *cmdUse:
        fmt.Fprintf(os.Stdout, "switching...\n")
    case *cmdSync:
        fmt.Fprintf(os.Stdout, "syncing...\n")
    default:
        flag.Usage()
    }
}