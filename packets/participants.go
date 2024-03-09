package packets

type ParticipantData struct {
	AiControlled    uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId        uint8    // Driver ID - see appendix, 255 if network human
	NetworkId       uint8    // Network ID – unique identifier for network players
	TeamId          uint8    // Team ID - see appendix
	MyTeam          uint8    // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber      uint8    // Race number of the car
	Nationality     uint8    // Nationality of the driver
	Name            [48]byte // Name of the participant in UTF-8 format – null-terminated
	YourTelemetry   uint8    // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8    // The player's show online names setting, 0 = off, 1 = on
	Platform        uint8    // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
}

type PacketParticipantsData struct {
	Header        PacketHeader        // Header
	NumActiveCars uint8               // Number of active cars in the data – should match the number of cars on HUD
	Participants  [22]ParticipantData // Participant data for all active cars
}
