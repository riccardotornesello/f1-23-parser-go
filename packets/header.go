package packets

type PacketId uint8

const (
	PacketMotionId              PacketId = 0  // Contains all motion data for player’s car – only sent while player is in control
	PacketSessionId             PacketId = 1  // Data about the session – track, time left
	PacketLapDataId             PacketId = 2  // Data about all the lap times of cars in the session
	PacketEventDataId           PacketId = 3  // Various notable events that happen during a session
	PacketParticipantsDataId    PacketId = 4  // List of participants in the session, mostly relevant for multiplayer
	PacketCarSetupId            PacketId = 5  // Packet detailing car setups for cars in the race
	PacketCarTelemetryId        PacketId = 6  // Telemetry data for all cars
	PacketCarStatusId           PacketId = 7  // Status data for all cars
	PacketFinalClassificationId PacketId = 8  // Final classification confirmation at the end of a race
	PacketLobbyInfoId           PacketId = 9  // Information about players in a multiplayer lobby
	PacketCarDamageId           PacketId = 10 // Damage status for all cars
	PacketSessionHistoryId      PacketId = 11 // Lap and tyre data for session
	PacketTyreSetsId            PacketId = 12 // Extended tyre set data
	PacketMotionExId            PacketId = 13 // Extended motion data for player car
)

type PacketHeader struct {
	PacketFormat            uint16   // 2023
	GameYear                uint8    // Game year - last two digits e.g. 23
	GameMajorVersion        uint8    // Game major version - "X.00"
	GameMinorVersion        uint8    // Game minor version - "1.XX"
	PacketVersion           uint8    // Version of this packet type, all start from 1
	PacketId                PacketId // Identifier for the packet type, see below
	SessionUID              uint64   // Unique identifier for the session
	SessionTime             float32  // Session timestamp
	FrameIdentifier         uint32   // Identifier for the frame the data was retrieved on
	OverallFrameIdentifier  uint32   // Overall identifier for the frame the data was retrieved on, doesn't go back after flashbacks
	PlayerCarIndex          uint8    // Index of player's car in the array
	SecondaryPlayerCarIndex uint8    // Index of secondary player's car in the array (splitscreen), 255 if no second player
}
