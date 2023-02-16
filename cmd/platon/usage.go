// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Contains the platon command usage template and generator.

package main

import (
	"io"
	"sort"

	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/internal/debug"
)

// AppHelpTemplate is the test template for the default, global app help topic.
var AppHelpTemplate = `NAME:
   {{.App.Name}} - {{.App.Usage}}

   Copyright 2019 The PlatON-Go Authors

USAGE:
   {{.App.HelpName}} [options]{{if .App.Commands}} command [command options]{{end}} {{if .App.ArgsUsage}}{{.App.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if .App.Version}}
VERSION:
   {{.App.Version}}
   {{end}}{{if len .App.Authors}}
AUTHOR(S):
   {{range .App.Authors}}{{ . }}{{end}}
   {{end}}{{if .App.Commands}}
COMMANDS:
   {{range .App.Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
   {{end}}{{end}}{{if .FlagGroups}}
{{range .FlagGroups}}{{.Name}} OPTIONS:
  {{range .Flags}}{{.}}
  {{end}}
{{end}}{{end}}{{if .App.Copyright }}
COPYRIGHT:
   {{.App.Copyright}}
   {{end}}
`

// flagGroup is a collection of flags belonging to a single topic.
type flagGroup struct {
	Name  string
	Flags []cli.Flag
}

// AppHelpFlagGroups is the application flags, grouped by functionality.
var AppHelpFlagGroups = []flagGroup{
	{
		Name: "PLATON",
		Flags: []cli.Flag{
			configFileFlag,
			utils.DataDirFlag,
			utils.AncientFlag,
			utils.KeyStoreDirFlag,
			utils.NoUSBFlag,
			utils.NetworkIdFlag,
			utils.MainFlag,
			utils.TestnetFlag,
			utils.SyncModeFlag,
			//	utils.EthStatsURLFlag,
			utils.IdentityFlag,
			utils.TxLookupLimitFlag,
			utils.LightKDFFlag,
		},
	},
	{
		Name: "DEVELOPER CHAIN",
		Flags: []cli.Flag{
			utils.DeveloperPeriodFlag,
		},
	},
	{
		Name: "TRANSACTION POOL",
		Flags: []cli.Flag{
			utils.TxPoolLocalsFlag,
			utils.TxPoolNoLocalsFlag,
			utils.TxPoolJournalFlag,
			utils.TxPoolRejournalFlag,
			utils.TxPoolPriceBumpFlag,
			utils.TxPoolAccountSlotsFlag,
			utils.TxPoolGlobalSlotsFlag,
			utils.TxPoolAccountQueueFlag,
			utils.TxPoolGlobalQueueFlag,
			utils.TxPoolGlobalTxCountFlag,
			utils.TxPoolLifetimeFlag,
			utils.TxPoolCacheSizeFlag,
		},
	},
	{
		Name: "PERFORMANCE TUNING",
		Flags: []cli.Flag{
			utils.CacheFlag,
			utils.CacheDatabaseFlag,
			utils.CacheTrieJournalFlag,
			utils.CacheTrieRejournalFlag,
			utils.CacheGCFlag,
			utils.CacheTrieDBFlag,
			utils.CachePreimagesFlag,
		},
	},
	{
		Name: "ACCOUNT",
		Flags: []cli.Flag{
			utils.UnlockedAccountFlag,
			utils.PasswordFileFlag,
			utils.InsecureUnlockAllowedFlag,
		},
	},
	{
		Name: "API AND CONSOLE",
		Flags: []cli.Flag{
			utils.IPCDisabledFlag,
			utils.IPCPathFlag,
			utils.HTTPEnabledFlag,
			utils.HTTPListenAddrFlag,
			utils.HTTPPortFlag,
			utils.HTTPApiFlag,
			utils.HTTPEnabledEthCompatibleFlag,
			utils.HTTPCORSDomainFlag,
			utils.HTTPVirtualHostsFlag,
			utils.WSEnabledFlag,
			utils.WSListenAddrFlag,
			utils.WSPortFlag,
			utils.WSApiFlag,
			utils.WSAllowedOriginsFlag,
			utils.GraphQLEnabledFlag,
			utils.GraphQLCORSDomainFlag,
			utils.GraphQLVirtualHostsFlag,
			utils.RPCGlobalGasCapFlag,
			utils.RPCGlobalTxFeeCapFlag,
			utils.JSpathFlag,
			utils.ExecFlag,
			utils.PreloadJSFlag,
		},
	},
	{
		Name: "NETWORKING",
		Flags: []cli.Flag{
			utils.BootnodesFlag,
			utils.BootnodesV4Flag,
			//	utils.BootnodesV5Flag,
			utils.ListenPortFlag,
			utils.MaxPeersFlag,
			utils.MaxConsensusPeersFlag,
			utils.MaxPendingPeersFlag,
			utils.NATFlag,
			utils.NoDiscoverFlag,
			//	utils.DiscoveryV5Flag,
			utils.NetrestrictFlag,
			utils.NodeKeyFileFlag,
			utils.NodeKeyHexFlag,
		},
	},
	{
		Name: "MINER",
		Flags: []cli.Flag{
			utils.MinerGasPriceFlag,
		},
	},
	{
		Name: "GAS PRICE ORACLE",
		Flags: []cli.Flag{
			utils.GpoBlocksFlag,
			utils.GpoPercentileFlag,
			utils.GpoMaxGasPriceFlag,
		},
	},
	{
		Name: "LOGGING AND DEBUGGING",
		Flags: append([]cli.Flag{
			utils.NoCompactionFlag,
		}, debug.Flags...),
	},
	{
		Name:  "METRICS AND STATS",
		Flags: metricsFlags,
	},
	{
		Name: "CBFT",
		Flags: []cli.Flag{
			utils.CbftPeerMsgQueueSize,
			utils.CbftWalDisabledFlag,
			utils.CbftMaxPingLatency,
			utils.CbftBlsPriKeyFileFlag,
			utils.CbftBlacklistDeadlineFlag,
		},
	},
	{
		Name: "DB",
		Flags: []cli.Flag{
			utils.DBNoGCFlag,
			utils.DBGCIntervalFlag,
			utils.DBGCTimeoutFlag,
			utils.DBGCMptFlag,
			utils.DBGCBlockFlag,
			utils.DBValidatorsHistoryFlag,
		},
	},
	{
		Name: "VM",
		Flags: []cli.Flag{
			utils.VMWasmType,
			utils.VmTimeoutDuration,
		},
	},
	{
		Name: "ALIASED (deprecated)",
		Flags: append([]cli.Flag{
			utils.LegacyRPCEnabledFlag,
			utils.LegacyRPCListenAddrFlag,
			utils.LegacyRPCPortFlag,
			utils.LegacyRPCCORSDomainFlag,
			utils.LegacyRPCVirtualHostsFlag,
			utils.LegacyRPCApiFlag,
			utils.LegacyWSListenAddrFlag,
			utils.LegacyWSPortFlag,
			utils.LegacyWSAllowedOriginsFlag,
			utils.LegacyWSApiFlag,
			utils.LegacyGpoBlocksFlag,
			utils.LegacyGpoPercentileFlag,
		}, debug.DeprecatedFlags...),
	},
	{
		Name: "MISC",
		Flags: []cli.Flag{
			cli.HelpFlag,
		},
	},
}

// byCategory sorts an array of flagGroup by Name in the order
// defined in AppHelpFlagGroups.
type byCategory []flagGroup

func (a byCategory) Len() int      { return len(a) }
func (a byCategory) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byCategory) Less(i, j int) bool {
	iCat, jCat := a[i].Name, a[j].Name
	iIdx, jIdx := len(AppHelpFlagGroups), len(AppHelpFlagGroups) // ensure non categorized flags come last

	for i, group := range AppHelpFlagGroups {
		if iCat == group.Name {
			iIdx = i
		}
		if jCat == group.Name {
			jIdx = i
		}
	}

	return iIdx < jIdx
}

func flagCategory(flag cli.Flag) string {
	for _, category := range AppHelpFlagGroups {
		for _, flg := range category.Flags {
			if flg.GetName() == flag.GetName() {
				return category.Name
			}
		}
	}
	return "MISC"
}

func init() {
	// Override the default app help template
	cli.AppHelpTemplate = AppHelpTemplate

	// Define a one shot struct to pass to the usage template
	type helpData struct {
		App        interface{}
		FlagGroups []flagGroup
	}

	// Override the default app help printer, but only for the global app help
	originalHelpPrinter := cli.HelpPrinter
	cli.HelpPrinter = func(w io.Writer, tmpl string, data interface{}) {
		if tmpl == AppHelpTemplate {
			// Iterate over all the flags and add any uncategorized ones
			categorized := make(map[string]struct{})
			for _, group := range AppHelpFlagGroups {
				for _, flag := range group.Flags {
					categorized[flag.String()] = struct{}{}
				}
			}
			deprecated := make(map[string]struct{})
			for _, flag := range utils.DeprecatedFlags {
				deprecated[flag.String()] = struct{}{}
			}
			// Only add uncategorized flags if they are not deprecated
			var uncategorized []cli.Flag
			for _, flag := range data.(*cli.App).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					if _, ok := deprecated[flag.String()]; !ok {
						uncategorized = append(uncategorized, flag)
					}
				}
			}
			if len(uncategorized) > 0 {
				// Append all ungategorized options to the misc group
				miscs := len(AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags)
				AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags = append(AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags, uncategorized...)

				// Make sure they are removed afterwards
				defer func() {
					AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags = AppHelpFlagGroups[len(AppHelpFlagGroups)-1].Flags[:miscs]
				}()
			}
			// Render out custom usage screen
			originalHelpPrinter(w, tmpl, helpData{data, AppHelpFlagGroups})
		} else if tmpl == utils.CommandHelpTemplate {
			// Iterate over all command specific flags and categorize them
			categorized := make(map[string][]cli.Flag)
			for _, flag := range data.(cli.Command).Flags {
				if _, ok := categorized[flag.String()]; !ok {
					categorized[flagCategory(flag)] = append(categorized[flagCategory(flag)], flag)
				}
			}

			// sort to get a stable ordering
			sorted := make([]flagGroup, 0, len(categorized))
			for cat, flgs := range categorized {
				sorted = append(sorted, flagGroup{cat, flgs})
			}
			sort.Sort(byCategory(sorted))

			// add sorted array to data and render with default printer
			originalHelpPrinter(w, tmpl, map[string]interface{}{
				"cmd":              data,
				"categorizedFlags": sorted,
			})
		} else {
			originalHelpPrinter(w, tmpl, data)
		}
	}
}
