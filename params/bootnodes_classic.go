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

// ClassicBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the Ethereum Classic network.
var ClassicBootnodes = []string{

	"enode://b9e893ea9cb4537f4fed154233005ae61b441cd0ecd980136138c304fefac194c25a16b73dac05fc66a4198d0c15dd0f33af99b411882c68a019dfa6bb703b9d@18.130.93.66:30303",
}

var dnsPrefixETC = "enrtree://AJE62Q4DUX4QMMXEHCSSCSC65TDHZYSMONSD64P3WULVLSF6MRQ3K@"

var ClassicDNSNetwork1 = dnsPrefixETC + "all.classic.blockd.info"
