package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/adityarathod/snowflakey/epoch"
	"github.com/adityarathod/snowflakey/original"
)

/**
 * This is a simple microservice that generates snowflakes and returns them as HTTP responses.
 * There's no real error handling, atomicity around SequenceNumber, or logging.
 * It's not suitable for production use :)
 */

var defaultEpoch = time.Date(2022, time.November, 30, 19, 38, 0, 0, time.UTC)

type SnowflakeGeneratorGlobals struct {
	Epoch      *epoch.Epoch
	MachineID  int16
	Port       int
	SequenceID int16
}

var globals *SnowflakeGeneratorGlobals

// Parses command line flags and initializes the snowflake generator context.
func initCtx() {
	epochMillis := flag.Int64("epoch", defaultEpoch.UnixMilli(), "The epoch time in milliseconds")
	machineID := flag.Int("machine", 1, "The machine ID")
	port := flag.Int("port", 4040, "The port to listen on for HTTP requests")
	flag.Parse()

	globals = &SnowflakeGeneratorGlobals{
		Epoch:     epoch.New(*epochMillis),
		MachineID: int16(*machineID),
		Port:      *port,
	}
	slog.Info("Initialized snowflake generator globals", "ctx", globals)
}

// Main HTTP request handler.
func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	definition := &original.Data{
		EpochTimeMillis: globals.Epoch.GenerateEpochTimeMillis(),
		MachineID:       globals.MachineID,
		SequenceNumber:  globals.SequenceID,
	}
	err := definition.Validate()
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid snowflake definition: %v", err), http.StatusInternalServerError)
		return
	}
	globals.SequenceID++
	snowflake := original.Generate(definition)
	slog.Info("Generated snowflake", slog.Int64("snowflake", snowflake), "definition", fmt.Sprintf("%+v", definition))
	fmt.Fprintf(w, "%d", snowflake)
}

func main() {
	initCtx()
	http.HandleFunc("/", handleRequest)
	slog.Info("Listening for HTTP requests", "port", globals.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", globals.Port), nil)
}
