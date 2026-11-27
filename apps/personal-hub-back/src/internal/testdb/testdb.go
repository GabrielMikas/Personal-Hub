// Package testdb provides an isolated in-memory database for tests.
package testdb

import (
	"errors"
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/personalhub/personalhub-back/src/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// New opens a fresh in-memory SQLite database, runs the application
// migrations and installs it as the global connection (config.GetDB) for
// the duration of the test. The previous connection is restored on cleanup.
func New(t testing.TB) *gorm.DB {
	t.Helper()

	conn, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("opening sqlite: %v", err)
	}

	sqlDB, err := conn.DB()
	if err != nil {
		t.Fatalf("getting sql.DB: %v", err)
	}
	// A single connection keeps the in-memory database (and the attached
	// schemas below) alive and shared across queries.
	sqlDB.SetMaxOpenConns(1)

	// Tables live in Postgres schemas (e.g. "hobbies.cards"); SQLite maps a
	// schema to an attached database.
	if err := conn.Exec("ATTACH DATABASE ':memory:' AS hobbies").Error; err != nil {
		t.Fatalf("attaching hobbies schema: %v", err)
	}

	if err := config.Migrate(conn); err != nil {
		t.Fatalf("migrating: %v", err)
	}

	previous := config.GetDB()
	config.SetDB(conn)
	t.Cleanup(func() {
		config.SetDB(previous)
		sqlDB.Close()
	})
	return conn
}

// ErrInjected is returned by every write on a connection passed to FailWrites.
var ErrInjected = errors.New("testdb: injected write failure")

// FailWrites makes every create, update and delete on conn fail with
// ErrInjected, while reads keep working. It is used to exercise error paths
// that happen after a successful lookup.
func FailWrites(t testing.TB, conn *gorm.DB) {
	t.Helper()

	fail := func(tx *gorm.DB) { tx.AddError(ErrInjected) }
	callbacks := conn.Callback()
	for _, err := range []error{
		callbacks.Create().Before("gorm:create").Register("testdb:fail_create", fail),
		callbacks.Update().Before("gorm:update").Register("testdb:fail_update", fail),
		callbacks.Delete().Before("gorm:delete").Register("testdb:fail_delete", fail),
	} {
		if err != nil {
			t.Fatalf("registering failure callback: %v", err)
		}
	}
}
