data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/gen-sql/",
  ]
}

env "gorm" {
    url = "postgres://root:secret@localhost:5432/training-system?sslmode=disable"
    src = data.external_schema.gorm.url
    dev = "docker://postgres/15/dev"
    migration {
        dir = "file://internal/migrations"
    }
    format {
        migrate {
        diff = "{{ sql . \"  \" }}"
        }
    }
}