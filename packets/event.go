package packets

type EventStringCode [4]uint8

var (
	SessionStartedCode     EventStringCode = EventStringCode{'S', 'S', 'T', 'A'}
	SessionEndedCode       EventStringCode = EventStringCode{'S', 'E', 'N', 'D'}
	FastestLapCode         EventStringCode = EventStringCode{'F', 'T', 'L', 'P'}
	RetirementCode         EventStringCode = EventStringCode{'R', 'T', 'M', 'T'}
	DrsEnabledCode         EventStringCode = EventStringCode{'D', 'R', 'S', 'E'}
	DrsDisabledCode        EventStringCode = EventStringCode{'D', 'R', 'S', 'D'}
	TeamMateInPitsCode     EventStringCode = EventStringCode{'T', 'M', 'P', 'T'}
	ChequeredFlagCode      EventStringCode = EventStringCode{'C', 'H', 'Q', 'F'}
	RaceWinnerCode         EventStringCode = EventStringCode{'R', 'C', 'W', 'N'}
	PenaltyIssuedCode      EventStringCode = EventStringCode{'P', 'E', 'N', 'A'}
	SpeedTrapTriggeredCode EventStringCode = EventStringCode{'S', 'P', 'T', 'P'}
	StartLightsCode        EventStringCode = EventStringCode{'S', 'T', 'L', 'G'}
	LightsOutCode          EventStringCode = EventStringCode{'L', 'G', 'O', 'T'}
	DriveThroughServedCode EventStringCode = EventStringCode{'D', 'T', 'S', 'V'}
	StopGoServedCode       EventStringCode = EventStringCode{'S', 'G', 'S', 'V'}
	FlashbackCode          EventStringCode = EventStringCode{'F', 'L', 'B', 'K'}
	ButtonStatusCode       EventStringCode = EventStringCode{'B', 'U', 'T', 'N'}
	RedFlagCode            EventStringCode = EventStringCode{'R', 'D', 'F', 'L'}
	OvertakeCode           EventStringCode = EventStringCode{'O', 'V', 'T', 'K'}
)

type FastestLap struct {
	VehicleIdx uint8   // Vehicle index of the car achieving the fastest lap
	LapTime    float32 // Lap time is in seconds
}

type Retirement struct {
	VehicleIdx uint8 // Vehicle index of the car retiring
}

type TeamMateInPits struct {
	VehicleIdx uint8 // Vehicle index of the team mate
}

type RaceWinner struct {
	VehicleIdx uint8 // Vehicle index of the race winner
}

type Penalty struct {
	PenaltyType      uint8 // Penalty type – see Appendices
	InfringementType uint8 // Infringement type – see Appendices
	VehicleIdx       uint8 // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8 // Vehicle index of the other car involved
	Time             uint8 // Time gained, or time spent doing the action in seconds
	LapNum           uint8 // Lap the penalty occurred on
	PlacesGained     uint8 // Number of places gained by this
}

type SpeedTrap struct {
	VehicleIdx                 uint8   // Vehicle index of the vehicle triggering the speed trap
	Speed                      float32 // Top speed achieved in kilometres per hour
	IsOverallFastestInSession  uint8   // Overall fastest speed in the session = 1, otherwise 0
	IsDriverFastestInSession   uint8   // Fastest speed for the driver in the session = 1, otherwise 0
	FastestVehicleIdxInSession uint8   // Vehicle index of the vehicle that is the fastest
	FastestSpeedInSession      float32 // Speed of the vehicle that is the fastest
}

type StartLIghts struct {
	NumLights uint8 // Number of lights showing
}

type DriveThroughPenaltyServed struct {
	VehicleIdx uint8 // Vehicle index of the vehicle serving the drive-through penalty
}

type StopGoPenaltyServed struct {
	VehicleIdx uint8 // Vehicle index of the vehicle serving the stop-go penalty
}

type Flashback struct {
	FlashbackFrameIdentifier uint32  // Frame identifier flashed back to
	FlashbackSessionTime     float32 // Session time flashed back to
}

type Buttons struct {
	ButtonStatus uint32 // Bit flags specifying which buttons are being pressed
}

type Overtake struct {
	OvertakingVehicleIdx     uint8 // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8 // Vehicle index of the vehicle being overtaken
}

type PacketEventData struct {
	Header          PacketHeader    // Header
	EventStringCode EventStringCode // Event string code, see below
	EventDetails    [12]byte        // Event details - should be interpreted differently for each type
}
