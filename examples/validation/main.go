package main

import (
	"fmt"
	"time"

	"github.com/adityarathod/snowflakey/epoch"
	"github.com/adityarathod/snowflakey/original"
)

/**
 * This is a minimal example of how validation works for a snowflake definition.
 */

// Date significantly in the future, to generate negative snowflake epoch times
var epochDate = time.Date(2050, time.January, 1, 0, 0, 0, 0, time.UTC)
var epochHelper = epoch.New(epochDate.UnixMilli())

func main() {
	definition := &original.Data{
		// This will generate a negative number
		EpochTimeMillis: epochHelper.GenerateEpochTimeMillis(),
		// This is outside the 10-bit range
		MachineID: 5000,
		// This is outside the 12 bit range
		SequenceNumber: 6000,
	}
	fmt.Printf("Snowflake definition: %+v\n", definition)
	err := definition.Validate()
	fmt.Printf("Validation error(s):\n%v\n", err)
}
