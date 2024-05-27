package original

import "errors"

// Represents the data encoded in a snowflake.
type Data struct {
	// Time since the snowflake's epoch, in milliseconds
	EpochTimeMillis int64
	// The ID of the machine that generated the snowflake
	MachineID int16
	// The sequence number of the snowflake on that machine
	SequenceNumber int16
}

// Validates the epoch time. It must be non-negative and fit within 41 bits.
func (def *Data) ValidateEpochTimeMillis() error {
	if def.EpochTimeMillis < 0 {
		return errors.New("epoch time must be non-negative")
	}
	if def.EpochTimeMillis >= 1<<41 {
		return errors.New("epoch time must fit within 41 bits")
	}
	return nil
}

// Validates the machine ID. It must be non-negative and fit within 10 bits.
func (def *Data) ValidateMachineID() error {
	if def.MachineID < 0 {
		return errors.New("machine ID must be non-negative")
	}
	if def.MachineID >= 1<<10 {
		return errors.New("machine ID must fit within 10 bits")
	}
	return nil
}

// Validates the sequence number. It must be non-negative and fit within 12 bits.
func (def *Data) ValidateSequenceNumber() error {
	if def.SequenceNumber < 0 {
		return errors.New("sequence number must be non-negative")
	}
	if def.SequenceNumber >= 1<<12 {
		return errors.New("sequence number must fit within 12 bits")
	}
	return nil
}

func (def *Data) Validate() error {
	return errors.Join(
		def.ValidateEpochTimeMillis(),
		def.ValidateMachineID(),
		def.ValidateSequenceNumber(),
	)
}
