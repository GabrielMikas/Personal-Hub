package config_test

import (
	"strings"
	"testing"

	"github.com/personalhub/personalhub-back/src/config"
	"github.com/personalhub/personalhub-back/src/internal/testdb"
)

func TestInitFails(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		wantErr string
	}{
		{name: "missing DSN", dsn: "", wantErr: "DB_DSN not set"},
		{
			name:    "unreachable database",
			dsn:     "host=127.0.0.1 port=1 user=test dbname=test sslmode=disable connect_timeout=1",
			wantErr: "opening database",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DB_DSN", tt.dsn)

			err := config.Init()
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Init() error = %v, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestSetDB(t *testing.T) {
	conn := testdb.New(t)
	config.SetDB(nil)
	if config.GetDB() != nil {
		t.Fatal("GetDB() should return the connection given to SetDB")
	}
	config.SetDB(conn)
	if config.GetDB() != conn {
		t.Fatal("GetDB() should return the connection given to SetDB")
	}
}

// The success path is exercised by testdb.New, which calls Migrate on a
// fresh database. Re-running it is not tested because the SQLite driver
// cannot detect existing tables in attached schemas.
func TestMigrateFailsOnClosedDatabase(t *testing.T) {
	conn := testdb.New(t)

	sqlDB, _ := conn.DB()
	sqlDB.Close()
	if err := config.Migrate(conn); err == nil {
		t.Fatal("Migrate() on a closed database should fail")
	}
}
