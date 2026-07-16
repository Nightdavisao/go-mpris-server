package server

import (
	"errors"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/quarckster/go-mpris-server/internal"
	"github.com/quarckster/go-mpris-server/pkg/types"
)

type Server struct {
	serviceName   string
	Conn          *dbus.Conn
	RootAdapter   types.OrgMprisMediaPlayer2Adapter
	PlayerAdapter types.OrgMprisMediaPlayer2PlayerAdapter
	stop          chan bool
	ready         chan struct{}
	readyOnce     sync.Once
}

// Create a new server with a given name and initialize needed data.
func NewServer(
	name string,
	rootAdapter types.OrgMprisMediaPlayer2Adapter,
	playerAdapter types.OrgMprisMediaPlayer2PlayerAdapter,
) *Server {
	server := Server{
		serviceName:   "org.mpris.MediaPlayer2." + name,
		RootAdapter:   rootAdapter,
		PlayerAdapter: playerAdapter,
		stop:          make(chan bool, 1),
		ready:         make(chan struct{}),
	}
	return &server
}

func (s *Server) Ready() <-chan struct{} {
	return s.ready
}

func (s *Server) exportMethods() error {
	root := internal.NewOrgMprisMediaPlayer2(s.RootAdapter)
	player := internal.NewOrgMprisMediaPlayer2Player(s.PlayerAdapter)
	properties := internal.NewOrgFreedesktopDBusProperties(root, player)
	return internal.ExportMethods(s.Conn, root, player, properties)
}

// Start the server and block.
func (s *Server) Listen() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		return err
	}
	s.Conn = conn
	err = s.exportMethods()
	if err != nil {
		s.Conn.ReleaseName(s.serviceName)
		s.Conn.Close()
		return err
	}
	reply, err := s.Conn.RequestName(s.serviceName, dbus.NameFlagReplaceExisting)
	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		s.Conn.Close()
		return errors.New("Unable to claim " + s.serviceName)
	}
	s.readyOnce.Do(func() {
		close(s.ready)
	})
	return nil
}

// Release the claimed bus name and close the connection.
func (s *Server) Stop() error {
	if s.Conn == nil {
		return errors.New("server is not started")
	}
	var err error
	err = internal.UnexportMethods(s.Conn)
	if err != nil {
		s.stop <- true
		return err
	}
	_, err = s.Conn.ReleaseName(s.serviceName)
	if err != nil {
		s.stop <- true
		return err
	}
	err = s.Conn.Close()
	if err != nil {
		s.stop <- true
		return err
	}
	s.stop <- true
	return nil
}
