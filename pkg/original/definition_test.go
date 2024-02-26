package original

import "testing"

func TestValidateEpochTimeMillisValidValue(t *testing.T) {
	def := &Definition{EpochTimeMillis: 0}
	if def.ValidateEpochTimeMillis() != nil {
		t.Error("Expected nil error for valid epoch time")
	}
}

func TestValidateEpochTimeMillisNegativeValue(t *testing.T) {
	def := &Definition{EpochTimeMillis: -1}
	if def.ValidateEpochTimeMillis() == nil {
		t.Error("Expected non-nil error for negative epoch time")
	}
}

func TestValidateEpochTimeMillisTooLargeValue(t *testing.T) {
	def := &Definition{EpochTimeMillis: 1 << 41}
	if def.ValidateEpochTimeMillis() == nil {
		t.Error("Expected non-nil error for too large epoch time")
	}
}

func TestValidateMachineIDValidValue(t *testing.T) {
	def := &Definition{MachineID: 0}
	if def.ValidateMachineID() != nil {
		t.Error("Expected nil error for valid machine ID")
	}
}

func TestValidateMachineIDNegativeValue(t *testing.T) {
	def := &Definition{MachineID: -1}
	if def.ValidateMachineID() == nil {
		t.Error("Expected non-nil error for negative machine ID")
	}
}

func TestValidateMachineIDTooLargeValue(t *testing.T) {
	def := &Definition{MachineID: 1 << 10}
	if def.ValidateMachineID() == nil {
		t.Error("Expected non-nil error for too large machine ID")
	}
}

func TestValidateSequenceNumberValidValue(t *testing.T) {
	def := &Definition{SequenceNumber: 0}
	if def.ValidateSequenceNumber() != nil {
		t.Error("Expected nil error for valid sequence number")
	}
}

func TestValidateSequenceNumberNegativeValue(t *testing.T) {
	def := &Definition{SequenceNumber: -1}
	if def.ValidateSequenceNumber() == nil {
		t.Error("Expected non-nil error for negative sequence number")
	}
}

func TestValidateSequenceNumberTooLargeValue(t *testing.T) {
	def := &Definition{SequenceNumber: 1 << 12}
	if def.ValidateSequenceNumber() == nil {
		t.Error("Expected non-nil error for too large sequence number")
	}
}

func TestValidateAllValidValues(t *testing.T) {
	def := &Definition{
		EpochTimeMillis: 0,
		MachineID:       0,
		SequenceNumber:  0,
	}
	if def.ValidateAll() != nil {
		t.Error("Expected nil error for valid values")
	}
}

func TestValidateAllInvalidEpochTimeMillis(t *testing.T) {
	def := &Definition{
		EpochTimeMillis: -1,
		MachineID:       0,
		SequenceNumber:  0,
	}
	if def.ValidateAll() == nil {
		t.Error("Expected non-nil error for invalid epoch time")
	}
}

func TestValidateAllInvalidMachineID(t *testing.T) {
	def := &Definition{
		EpochTimeMillis: 0,
		MachineID:       -1,
		SequenceNumber:  0,
	}
	if def.ValidateAll() == nil {
		t.Error("Expected non-nil error for invalid machine ID")
	}
}

func TestValidateAllInvalidSequenceNumber(t *testing.T) {
	def := &Definition{
		EpochTimeMillis: 0,
		MachineID:       0,
		SequenceNumber:  -1,
	}
	if def.ValidateAll() == nil {
		t.Error("Expected non-nil error for invalid sequence number")
	}
}
