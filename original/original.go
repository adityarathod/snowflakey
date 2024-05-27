package original

// Generates a snowflake from the given definition.
func Generate(data *Data) int64 {
	snowflake := int64(0)
	// get last 12 bits of sequence number and add to end of snowflake
	snowflake |= int64(data.SequenceNumber & 0xfff)
	// get last 10 bits of machine id and add to snowflake starting at bit 12 from the right
	snowflake |= int64(data.MachineID&0x3ff) << 12
	// get last 41 bits of epoch time and add to snowflake starting at bit 22 from the right
	snowflake |= int64(data.EpochTimeMillis&0x1ffffffffff) << 22

	return snowflake
}

// Parses a snowflake into its constituent definition.
func Parse(snowflake int64) *Data {
	info := &Data{}
	// get last 12 bits of snowflake and set them as the sequence number
	info.SequenceNumber = int16(snowflake & 0xfff)
	// get the next 10 bits of snowflake and set them as the machine id
	info.MachineID = int16((snowflake >> 12) & 0x3ff)
	// get the last 41 bits of snowflake and set them as the time since epoch
	info.EpochTimeMillis = snowflake >> 22
	return info
}
