package main

import (
	"github.com/riccardotornesello/f1-23-parser-go/packets"
)

func NewPacketById(packetId packets.PacketId) interface{} {
	switch packetId {
	case packets.PacketMotionId:
		return new(packets.PacketMotionData)
	case packets.PacketSessionId:
		return new(packets.PacketSessionData)
	case packets.PacketLapDataId:
		return new(packets.PacketLapData)
	case packets.PacketEventDataId:
		return new(packets.PacketEventData)
	case packets.PacketParticipantsDataId:
		return new(packets.PacketParticipantsData)
	case packets.PacketCarSetupId:
		return new(packets.PacketCarSetupData)
	case packets.PacketCarTelemetryId:
		return new(packets.PacketCarTelemetryData)
	case packets.PacketCarStatusId:
		return new(packets.PacketCarStatusData)
	case packets.PacketFinalClassificationId:
		return new(packets.PacketFinalClassificationData)
	case packets.PacketLobbyInfoId:
		return new(packets.PacketLobbyInfoData)
	case packets.PacketCarDamageId:
		return new(packets.PacketCarDamageData)
	case packets.PacketSessionHistoryId:
		return new(packets.PacketSessionHistoryData)
	case packets.PacketTyreSetsId:
		return new(packets.PacketTyreSetsData)
	case packets.PacketMotionExId:
		return new(packets.PacketMotionExData)
	}

	return nil
}

func NewEventById(eventId packets.EventStringCode) interface{} {
	switch eventId {
	case packets.SessionStartedCode:
		return new(interface{})
	case packets.SessionEndedCode:
		return new(interface{})
	case packets.FastestLapCode:
		return new(packets.FastestLap)
	case packets.RetirementCode:
		return new(packets.Retirement)
	case packets.DrsEnabledCode:
		return new(interface{})
	case packets.DrsDisabledCode:
		return new(interface{})
	case packets.TeamMateInPitsCode:
		return new(packets.TeamMateInPits)
	case packets.ChequeredFlagCode:
		return new(interface{})
	case packets.RaceWinnerCode:
		return new(packets.RaceWinner)
	case packets.PenaltyIssuedCode:
		return new(packets.Penalty)
	case packets.SpeedTrapTriggeredCode:
		return new(packets.SpeedTrap)
	case packets.StartLightsCode:
		return new(packets.StartLIghts)
	case packets.LightsOutCode:
		return new(interface{})
	case packets.DriveThroughServedCode:
		return new(packets.DriveThroughPenaltyServed)
	case packets.StopGoServedCode:
		return new(packets.StopGoPenaltyServed)
	case packets.FlashbackCode:
		return new(packets.Flashback)
	case packets.ButtonStatusCode:
		return new(packets.Buttons)
	case packets.RedFlagCode:
		return new(interface{})
	case packets.OvertakeCode:
		return new(packets.Overtake)
	}

	return nil
}
