package types

import (
	"github.com/godbus/dbus/v5"
)

type OrgMprisMediaPlayer2Adapter interface {
	Raise() error
	Quit() error
	CanQuit() (bool, error)
	CanRaise() (bool, error)
	HasTrackList() (bool, error)
	Identity() (string, error)
	SupportedUriSchemes() ([]string, error)
	SupportedMimeTypes() ([]string, error)
}

type OrgMprisMediaPlayer2AdapterFullscreen interface {
	Fullscreen() (bool, error)
	SetFullscreen(bool) error
}

type OrgMprisMediaPlayer2AdapterCanSetFullscreen interface {
	CanSetFullscreen() (bool, error)
}

type OrgMprisMediaPlayer2AdapterDesktopEntry interface {
	DesktopEntry() (string, error)
}

type OrgMprisMediaPlayer2PlayerAdapter interface {
	Next() error
	Previous() error
	Pause() error
	PlayPause() error
	Stop() error
	Play() error
	Seek(offset Microseconds) error
	SetPosition(trackId dbus.ObjectPath, position Microseconds) error
	OpenUri(uri string) error
	PlaybackStatus() (PlaybackStatus, error)
	Rate() (float64, error)
	SetRate(float64) error
	Metadata() (Metadata, error)
	Volume() (float64, error)
	SetVolume(float64) error
	Position() (int64, error)
	MinimumRate() (float64, error)
	MaximumRate() (float64, error)
	CanGoNext() (bool, error)
	CanGoPrevious() (bool, error)
	CanPlay() (bool, error)
	CanPause() (bool, error)
	CanSeek() (bool, error)
	CanControl() (bool, error)
}

type OrgMprisMediaPlayer2PlayerAdapterLoopStatus interface {
	LoopStatus() (LoopStatus, error)
	SetLoopStatus(LoopStatus) error
}

type OrgMprisMediaPlayer2PlayerAdapterShuffle interface {
	Shuffle() (bool, error)
	SetShuffle(bool) error
}

type Microseconds int64

type Metadata struct {
	TrackId        dbus.ObjectPath
	Length         Microseconds
	ArtUrl         string
	Album          string
	AlbumArtist    []string
	Artist         []string
	AsText         string
	AudioBPM       int
	AutoRating     float64
	Comment        []string
	Composer       []string
	ContentCreated string
	DiscNumber     int
	FirstUsed      string
	Genre          []string
	LastUsed       string
	Lyricist       []string
	Title          string
	TrackNumber    int
	Url            string
	UseCount       int
	UserRating     float64
}

func (m *Metadata) MakeMap() map[string]dbus.Variant {
	metadataMap := make(map[string]dbus.Variant)

	if m.TrackId.IsValid() {
		metadataMap["mpris:trackid"] = dbus.MakeVariant(m.TrackId)
	}
	if m.Length >= 0 {
		metadataMap["mpris:length"] = dbus.MakeVariant(m.Length)
	}
	if m.ArtUrl != "" {
		metadataMap["mpris:artUrl"] = dbus.MakeVariant(m.ArtUrl)
	}
	if m.Album != "" {
		metadataMap["xesam:album"] = dbus.MakeVariant(m.Album)
	}
	if len(m.AlbumArtist) > 0 {
		metadataMap["xesam:albumArtist"] = dbus.MakeVariant(m.AlbumArtist)
	}
	if len(m.Artist) > 0 {
		metadataMap["xesam:artist"] = dbus.MakeVariant(m.Artist)
	}
	if m.AsText != "" {
		metadataMap["xesam:asText"] = dbus.MakeVariant(m.AsText)
	}
	if m.AudioBPM > 0 {
		metadataMap["xesam:audioBPM"] = dbus.MakeVariant(m.AudioBPM)
	}
	if m.AutoRating > 0 && m.AutoRating <= 1 {
		metadataMap["xesam:autoRating"] = dbus.MakeVariant(m.AutoRating)
	}
	if len(m.Comment) > 0 {
		metadataMap["xesam:comment"] = dbus.MakeVariant(m.Comment)
	}
	if len(m.Composer) > 0 {
		metadataMap["xesam:composer"] = dbus.MakeVariant(m.Composer)
	}
	if m.ContentCreated != "" {
		metadataMap["xesam:contentCreated"] = dbus.MakeVariant(m.ContentCreated)
	}
	if m.DiscNumber > 0 {
		metadataMap["xesam:discNumber"] = dbus.MakeVariant(m.DiscNumber)
	}
	if m.FirstUsed != "" {
		metadataMap["xesam:firstUsed"] = dbus.MakeVariant(m.FirstUsed)
	}
	if len(m.Genre) > 0 {
		metadataMap["xesam:genre"] = dbus.MakeVariant(m.Genre)
	}
	if m.LastUsed != "" {
		metadataMap["xesam:lastUsed"] = dbus.MakeVariant(m.LastUsed)
	}
	if len(m.Lyricist) > 0 {
		metadataMap["xesam:lyricist"] = dbus.MakeVariant(m.Lyricist)
	}
	if m.Title != "" {
		metadataMap["xesam:title"] = dbus.MakeVariant(m.Title)
	}
	if m.TrackNumber > 0 {
		metadataMap["xesam:trackNumber"] = dbus.MakeVariant(m.TrackNumber)
	}
	if m.Url != "" {
		metadataMap["xesam:url"] = dbus.MakeVariant(m.Url)
	}
	if m.UseCount > 0 {
		metadataMap["xesam:useCount"] = dbus.MakeVariant(m.UseCount)
	}
	if m.UserRating > 0 && m.UserRating <= 1 {
		metadataMap["xesam:userRating"] = dbus.MakeVariant(m.UserRating)
	}

	return metadataMap
}

type PlaybackStatus string

const (
	PlaybackStatusPlaying PlaybackStatus = "Playing"
	PlaybackStatusPaused  PlaybackStatus = "Paused"
	PlaybackStatusStopped PlaybackStatus = "Stopped"
)

type LoopStatus string

const (
	LoopStatusNone     LoopStatus = "None"
	LoopStatusTrack    LoopStatus = "Track"
	LoopStatusPlaylist LoopStatus = "Playlist"
)

type OrgMprisMediaPlayer2EventHandler interface {
	OnAll() error
}

type OrgMprisMediaPlayer2PlayerEventHandler interface {
	OnEnded() error
	OnVolume() error
	OnPlayback() error
	OnPosition() error
	OnPlayPause() error
	OnTitle() error
	OnSeek(position Microseconds) error
	OnOptions() error
	OnAll() error
}
