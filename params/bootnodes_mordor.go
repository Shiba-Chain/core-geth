// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package params

// MordorBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Ethereum Classic Mordor network.
// https://github.com/etclabscore/mordor/blob/master/static-nodes.json
var MordorBootnodes = []string{
	"enode://534d18fd46c5cd5ba48a68250c47cea27a1376869755ed631c94b91386328039eb607cf10dd8d0aa173f5ec21e3fb45c5d7a7aa904f97bc2557e9cb4ccc703f1@51.158.190.99:30303", // @q9f lyrae
}

var MordorDNSNetwork1 = dnsPrefixETC + "all.mordor.blockd.info"
