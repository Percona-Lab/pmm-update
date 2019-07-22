// pmm-update
// Copyright (C) 2019 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package main // import "github.com/percona/pmm-update/pmm2-update"

import (
	"context"
	"flag"
	"log"

	"github.com/percona/pmm/version"

	"github.com/percona/pmm-update/pmm2-update/yum"
)

func main() {
	log.SetFlags(0)
	log.Print(version.FullInfo())
	log.SetPrefix("pmm-update: ")
	flag.Parse()

	versions, err := yum.CheckVersions(context.Background())
	if err != nil {
		log.Fatalf("%+v", err)
	}
	log.Printf("%+v", versions)
}