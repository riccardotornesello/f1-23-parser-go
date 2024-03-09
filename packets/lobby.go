package packets

type LobbyInfoData struct {
	AiControlled uint8    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamID       uint8    // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8    // Nationality of the driver
	Platform     uint8    // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	Name         [48]byte // Name of participant in UTF-8 format – null terminated
	// Will be truncated with ... (U+2026) if too long
	CarNumber   uint8 // Car number of the player
	ReadyStatus uint8 // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketLobbyInfoData struct {
	Header       PacketHeader      // Header
	NumPlayers   uint8             // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData // Lobby information data for all players
}
