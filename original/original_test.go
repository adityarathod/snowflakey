package original

import "testing"

const wikipediaExampleSnowflake = 1541815603606036480

var wikipediaExampleInfo = &Data{
	EpochTimeMillis: 367597485448,
	MachineID:       378,
	SequenceNumber:  0,
}

func TestGenerateWithWikipediaExample(t *testing.T) {
	info := wikipediaExampleInfo
	snowflake := Generate(info)
	if snowflake != wikipediaExampleSnowflake {
		t.Errorf("Expected %d, got %d", wikipediaExampleSnowflake, snowflake)
	}
}

func TestParseWithWikipediaExample(t *testing.T) {
	info := Parse(wikipediaExampleSnowflake)
	if *info != *wikipediaExampleInfo {
		t.Errorf("Expected %v, got %v", wikipediaExampleInfo, info)
	}
}
