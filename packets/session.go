package packets

type MarshalZone struct {
	ZoneStart float32 // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow
}

type WeatherForecastSample struct {
	SessionType            uint8 // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1
	TimeOffset             uint8 // Time in minutes the forecast is for
	Weather                uint8 // Weather - 0 = clear, 1 = light cloud, 2 = overcast
	TrackTemperature       int8  // Track temp. in degrees Celsius
	TrackTemperatureChange int8  // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8  // Air temp. in degrees Celsius
	AirTemperatureChange   int8  // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8 // Rain percentage (0-100)
}

type PacketSessionData struct {
	Header                          PacketHeader              // Header
	Weather                         uint8                     // Weather - 0 = clear, 1 = light cloud, 2 = overcast
	TrackTemperature                int8                      // Track temp. in degrees celsius
	AirTemperature                  int8                      // Air temp. in degrees celsius
	TotalLaps                       uint8                     // Total number of laps in this race
	TrackLength                     uint16                    // Track length in metres
	SessionType                     uint8                     // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P
	TrackId                         int8                      // -1 for unknown, see appendix
	Formula                         uint8                     // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2,
	SessionTimeLeft                 uint16                    // Time left in session in seconds
	SessionDuration                 uint16                    // Session duration in seconds
	PitSpeedLimit                   uint8                     // Pit speed limit in kilometres per hour
	GamePaused                      uint8                     // Whether the game is paused – network game only
	IsSpectating                    uint8                     // Whether the player is spectating
	SpectatorCarIndex               uint8                     // Index of the car being spectated
	SliProNativeSupport             uint8                     // SLI Pro support, 0 = inactive, 1 = active
	NumMarshalZones                 uint8                     // Number of marshal zones to follow
	MarshalZones                    [21]MarshalZone           // List of marshal zones – max 21
	SafetyCarStatus                 uint8                     // 0 = no safety car, 1 = full
	NetworkGame                     uint8                     // 0 = offline, 1 = online
	NumWeatherForecastSamples       uint8                     // Number of weather samples to follow
	WeatherForecastSamples          [56]WeatherForecastSample // Array of weather forecast samples
	ForecastAccuracy                uint8                     // 0 = Perfect, 1 = Approximate
	AiDifficulty                    uint8                     // AI Difficulty rating – 0-110
	SeasonLinkIdentifier            uint32                    // Identifier for season - persists across saves
	WeekendLinkIdentifier           uint32                    // Identifier for weekend - persists across saves
	SessionLinkIdentifier           uint32                    // Identifier for session - persists across saves
	PitStopWindowIdealLap           uint8                     // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap          uint8                     // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition           uint8                     // Predicted position to rejoin at (player)
	SteeringAssist                  uint8                     // 0 = off, 1 = on
	BrakingAssist                   uint8                     // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist                   uint8                     // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist                       uint8                     // 0 = off, 1 = on
	PitReleaseAssist                uint8                     // 0 = off, 1 = on
	ERSAssist                       uint8                     // 0 = off, 1 = on
	DRSAssist                       uint8                     // 0 = off, 1 = on
	DynamicRacingLine               uint8                     // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType           uint8                     // 0 = 2D, 1 = 3D
	GameMode                        uint8                     // Game mode id - see appendix
	RuleSet                         uint8                     // Ruleset - see appendix
	TimeOfDay                       uint32                    // Local time of day - minutes since midnight
	SessionLength                   uint8                     // 0 = None, 2 = Very Short, 3 = Short, 4 = Medium
	SpeedUnitsLeadPlayer            uint8                     // 0 = MPH, 1 = KPH
	TemperatureUnitsLeadPlayer      uint8                     // 0 = Celsius, 1 = Fahrenheit
	SpeedUnitsSecondaryPlayer       uint8                     // 0 = MPH, 1 = KPH
	TemperatureUnitsSecondaryPlayer uint8                     // 0 = Celsius, 1 = Fahrenheit
	NumSafetyCarPeriods             uint8                     // Number of safety cars called during session
	NumVirtualSafetyCarPeriods      uint8                     // Number of virtual safety cars called
	NumRedFlagPeriods               uint8                     // Number of red flags called during session
}
