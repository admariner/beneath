package migrations

import (
	"github.com/beneath-core/beneath-go/control/entity"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) (err error) {
		// Extensions
		_, err = db.Exec(`
			CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		`)
		if err != nil {
			return err
		}

		// Project
		err = db.Model(&entity.Project{}).CreateTable(defaultCreateOptions)
		if err != nil {
			return err
		}

		// Project name index
		_, err = db.Exec(`
			ALTER TABLE projects DROP CONSTRAINT projects_name_key;
			CREATE UNIQUE INDEX projects_name_key ON public.projects USING btree ((lower(name)));
		`)
		if err != nil {
			return err
		}

		// Done
		return nil
	}, func(db migrations.DB) (err error) {
		// Project
		err = db.Model(&entity.Project{}).DropTable(defaultDropOptions)
		if err != nil {
			return err
		}

		// Extensions
		_, err = db.Exec(`
			DROP EXTENSION "uuid-ossp";
		`)
		if err != nil {
			return err
		}

		// Done
		return nil
	})
}
