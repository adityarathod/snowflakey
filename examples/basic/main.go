package main

import (
	"fmt"
	"time"

	"github.com/adityarathod/snowflakey/epoch"
	"github.com/adityarathod/snowflakey/original"
)

/**
 * This is a minimal example of how to use the snowflakey package to generate a snowflake.
 * Note that this example does not include any error handling or validation.
 */

// Some significant date (here, this is the launch of chatgpt)
var epochDate = time.Date(2022, time.November, 30, 19, 38, 0, 0, time.UTC)
var customEpoch = epoch.New(epochDate.UnixMilli())

func main() {
	definition := &original.Data{
		EpochTimeMillis: customEpoch.GenerateEpochTimeMillis(),
		MachineID:       1,
		SequenceNumber:  0,
	}
	fmt.Printf("Snowflake definition: %+v\n", definition)
	snowflake := original.Generate(definition)
	fmt.Println("Generated snowflake:", snowflake)
}
