data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./cmd/gen-sql/",
  ]
}

env "gorm" {
    url = getenv("DATABASE_URL")
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