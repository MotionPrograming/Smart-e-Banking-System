package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
)

// Migration হলো database-এর জন্য “version control system” যা database schema-র পরিবর্তনগুলি পরিচালনা করে।
//
//	এটি database schema-র পরিবর্তনগুলি সংরক্ষণ করে এবং প্রয়োজনে সেই পরিবর্তনগুলি প্রয়োগ করতে সাহায্য করে। Migration ব্যবহারের মাধ্যমে,
//
// আপনি database schema-র পরিবর্তনগুলি সহজে ট্র্যাক করতে পারেন এবং বিভিন্ন পরিবেশে একই schema বজায় রাখতে পারেন।

func MigrateDB(db *sqlx.DB, dir string) error {
	migration := &migrate.FileMigrationSource{
		Dir: dir,
	}

	n, err := migrate.Exec(db.DB, "mysql", migration, migrate.Up)
	if err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	fmt.Printf("Successfully applied %d migrations!\n", n)
	return nil
}
