// Copyright 2016 The go-ethereum Authors
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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var customGenesisTests = []struct {
	genesis string
	query   string
	result  string
}{
	// Plain genesis file without anything extra
	{
		genesis: `{
    "alloc":{
        "1000000000000000000000000000000000000001":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000002":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000003":{
            "balance":"200000000000000000000000000"
        },
        "1000000000000000000000000000000000000004":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000005":{
            "balance":"0"
        },
        "60ceca9c1290ee56b98d4e160ef0453f7c40d219":{
            "balance":"8050000000000000000000000000"
        },
        "55bfd49472fd41211545b01713a9c3a97af78b05":{
            "balance":"2000000000000000000000000000"
        }
    },
    "EconomicModel":{
        "Common":{
            "MaxEpochMinutes":4,
            "MaxConsensusVals":4,
            "AdditionalCycleTime":16
        },
        "Staking":{
            "StakeThreshold":               1000000000000000000000000,
            "OperatingThreshold":             10000000000000000000,
            "MaxValidators":            24,
            "HesitateRatio":                1,
            "UnStakeFreezeDuration":           2
        },
        "Slashing":{
           "PackAmountAbnormal":   6,
           "SlashFractionDuplicateSign": 100,
           "DuplicateSignReportReward": 50,
           "SlashBlocksReward":20, 
           "MaxEvidenceAge":1
        },
        "Gov": {
            "VersionProposalVoteDurationSeconds": 160,
            "VersionProposalActive_ConsensusRounds": 5,
            "VersionProposalSupportRate": 0.667,
            "TextProposalVoteDurationSeconds": 160,
            "TextProposalVoteRate": 0.5,
            "TextProposalSupportRate": 0.667,          
            "CancelProposalVoteRate": 0.50,
            "CancelProposalSupportRate": 0.667
        },
        "Reward":{
            "NewBlockRate": 50,
            "PlatONFoundationYear": 10 
        },
        "InnerAcc":{
            "PlatONFundAccount": "0x493301712671ada506ba6ca7891f436d29185821",
            "PlatONFundBalance": 0,
            "CDFAccount": "0xc1f330b214668beac2e6418dd651b09c759a4bf5",
            "CDFBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0x0000000000000000000000000000000000000000",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "cbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "epoch":1,
            "amount":10,
            "validatorMode":"ppos",
            "period":10000
        }
    }
}`,
		query:  "platon.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
	//Genesis file with only cbft config
	{
		genesis: `{
    "alloc":{
        "1000000000000000000000000000000000000001":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000002":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000003":{
            "balance":"200000000000000000000000000"
        },
        "1000000000000000000000000000000000000004":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000005":{
            "balance":"0"
        },
        "60ceca9c1290ee56b98d4e160ef0453f7c40d219":{
            "balance":"8050000000000000000000000000"
        },
        "55bfd49472fd41211545b01713a9c3a97af78b05":{
            "balance":"2000000000000000000000000000"
        }
    },
    "EconomicModel":{
        "Common":{
            "MaxEpochMinutes":4,
            "MaxConsensusVals":4,
            "AdditionalCycleTime":16
        },
        "Staking":{
            "StakeThreshold":               1000000000000000000000000,
            "OperatingThreshold":             10000000000000000000,
            "MaxValidators":            24,
            "HesitateRatio":                1,
            "UnStakeFreezeDuration":           2
        },
        "Slashing":{
           "PackAmountAbnormal":   6,
           "SlashFractionDuplicateSign": 100,
           "DuplicateSignReportReward": 50,
           "SlashBlocksReward":20, 
           "MaxEvidenceAge":1
        },
        "Gov": {
            "VersionProposalVoteDurationSeconds": 160,
            "VersionProposalActive_ConsensusRounds": 5,
            "VersionProposalSupportRate": 0.667,
            "TextProposalVoteDurationSeconds": 160,
            "TextProposalVoteRate": 0.5,
            "TextProposalSupportRate": 0.667,          
            "CancelProposalVoteRate": 0.50,
            "CancelProposalSupportRate": 0.667
        },
        "Reward":{
            "NewBlockRate": 50,
            "PlatONFoundationYear": 10 
        },
        "InnerAcc":{
            "PlatONFundAccount": "0x493301712671ada506ba6ca7891f436d29185821",
            "PlatONFundBalance": 0,
            "CDFAccount": "0xc1f330b214668beac2e6418dd651b09c759a4bf5",
            "CDFBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0x0000000000000000000000000000000000000000",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "cbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "epoch":1,
            "amount":10,
            "validatorMode":"ppos",
            "period":10000
        }
    }
}`,
		query:  "platon.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
	//Genesis file with specific chain configurations
	{
		genesis: `{
    "alloc":{
        "1000000000000000000000000000000000000001":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000002":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000003":{
            "balance":"200000000000000000000000000"
        },
        "1000000000000000000000000000000000000004":{
            "balance":"0"
        },
        "1000000000000000000000000000000000000005":{
            "balance":"0"
        },
        "60ceca9c1290ee56b98d4e160ef0453f7c40d219":{
            "balance":"8050000000000000000000000000"
        },
        "55bfd49472fd41211545b01713a9c3a97af78b05":{
            "balance":"2000000000000000000000000000"
        }
    },
    "EconomicModel":{
        "Common":{
            "MaxEpochMinutes":4,
            "MaxConsensusVals":4,
            "AdditionalCycleTime":16
        },
        "Staking":{
            "StakeThreshold":               1000000000000000000000000,
            "OperatingThreshold":             10000000000000000000,
            "MaxValidators":            24,
            "HesitateRatio":                1,
            "UnStakeFreezeDuration":           2
        },
        "Slashing":{
           "PackAmountAbnormal":   6,
           "SlashFractionDuplicateSign": 100,
           "DuplicateSignReportReward": 50,
           "SlashBlocksReward":20, 
           "MaxEvidenceAge":1
        },
        "Gov": {
            "VersionProposalVoteDurationSeconds": 160,
            "VersionProposalActive_ConsensusRounds": 5,
            "VersionProposalSupportRate": 0.667,
            "TextProposalVoteDurationSeconds": 160,
            "TextProposalVoteRate": 0.5,
            "TextProposalSupportRate": 0.667,          
            "CancelProposalVoteRate": 0.50,
            "CancelProposalSupportRate": 0.667
        },
        "Reward":{
            "NewBlockRate": 50,
            "PlatONFoundationYear": 10 
        },
        "InnerAcc":{
            "PlatONFundAccount": "0x493301712671ada506ba6ca7891f436d29185821",
            "PlatONFundBalance": 0,
            "CDFAccount": "0xc1f330b214668beac2e6418dd651b09c759a4bf5",
            "CDFBalance": 331811981000000000000000000
        }
    },
    "coinbase":"0x0000000000000000000000000000000000000000",
    "extraData":"",
    "gasLimit":"0x2fefd8",
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
    "timestamp":"0x00",
    "config":{
        "chainId":101,
        "eip155Block":0,
        "interpreter":"wasm",
        "cbft":{
            "initialNodes":[
                {
                    "node":"enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789",
                    "blsPubKey":"d341a0c485c9ec00cecf7ea16323c547900f6a1bacb9daacb00c2b8bacee631f75d5d31b75814b7f1ae3a4e18b71c617bc2f230daa0c893746ed87b08b2df93ca4ddde2816b3ac410b9980bcc048521562a3b2d00e900fd777d3cf88ce678719"
                }
            ],
            "epoch":1,
            "amount":10,
            "validatorMode":"ppos",
            "period":10000
        }
    }
}`,
		query:  "platon.getBlock(0).nonce",
		result: "0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	},
}

// Tests that initializing Geth with a custom genesis block and chain definitions
// work properly.
func TestCustomGenesis(t *testing.T) {
	for i, tt := range customGenesisTests {
		// Create a temporary data directory to use and inspect later
		datadir := tmpdir(t)
		defer os.RemoveAll(datadir)

		// Initialize the data directory with the custom genesis block
		json := filepath.Join(datadir, "genesis.json")
		if err := ioutil.WriteFile(json, []byte(tt.genesis), 0600); err != nil {
			t.Fatalf("test %d: failed to write genesis file: %v", i, err)
		}
		runGeth(t, "--datadir", datadir, "init", json).WaitExit()

		// Query the custom genesis block
		geth := runGeth(t,
			"--datadir", datadir, "--maxpeers", "0", "--port", "0",
			"--nodiscover", "--nat", "none", "--ipcdisable",
			"--exec", tt.query, "console")
		t.Log("testi", i)
		geth.ExpectRegexp(tt.result)
		geth.ExpectExit()
	}
}
