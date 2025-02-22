package sys

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FuturFusion/migration-manager/internal/util"
)

// OS is a high-level facade for accessing operating-system level functionalities.
type OS struct {
	// Directories
	CacheDir string // Cache directory (e.g., /var/cache/migration-manager/)
	LogDir   string // Log directory (e.g. /var/log/).
	RunDir   string // Runtime directory (e.g. /run/migration-manager/).
	VarDir   string // Data directory (e.g. /var/lib/migration-manager/).
}

// DefaultOS returns a fresh uninitialized OS instance with default values.
func DefaultOS() *OS {
	newOS := &OS{
		CacheDir: util.CachePath(),
		LogDir:   util.LogPath(),
		RunDir:   util.RunPath(),
		VarDir:   util.VarPath(),
	}

	return newOS
}

// GetUnixSocket returns the full path to the unix.socket file that this daemon is listening on.
func (s *OS) GetUnixSocket() string {
	path := os.Getenv("MIGRATION_MANAGER_SOCKET")
	if path != "" {
		return path
	}

	return filepath.Join(s.RunDir, "unix.socket")
}

// LocalDatabaseDir returns the path of the local database directory.
func (s *OS) LocalDatabaseDir() string {
	return filepath.Join(s.VarDir, "database")
}

// Returns the name of the migration manger worker ISO image.
func (s *OS) GetMigrationManagerISOName() (string, error) {
	files, err := filepath.Glob(fmt.Sprintf("%s/migration-manager-minimal-boot*.iso", s.CacheDir))
	if err != nil {
		return "", err
	}

	if len(files) != 1 {
		return "", fmt.Errorf("Unable to determine migration manager ISO name")
	}

	return filepath.Base(files[0]), nil
}

// Returns the name of the virtio drivers ISO image.
func (s *OS) GetVirtioDriversISOName() (string, error) {
	files, err := filepath.Glob(fmt.Sprintf("%s/virtio-win-*.iso", s.CacheDir))
	if err != nil {
		return "", err
	}

	if len(files) != 1 {
		return "", fmt.Errorf("Unable to determine virtio drivers ISO name")
	}

	return filepath.Base(files[0]), nil
}
