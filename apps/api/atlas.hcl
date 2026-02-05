// Copyright (c) 2026 dilocash
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

variable "db_url" {
  type    = string
  default = getenv("DB_URL")
}

env "local" {
  // Source of the schema
  src = "file://../../packages/database/schema.sql"

  // URL of the database to migrate. 
  // This will be read from the DB_URL env var.
  url = var.db_url

  // Dev database for Atlas to use as a sandbox to calculate diffs
  dev = "docker://postgres/16/dev"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
