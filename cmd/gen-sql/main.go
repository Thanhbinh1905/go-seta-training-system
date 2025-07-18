package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	assetsmodel "github.com/Thanhbinh1905/seta-training-system/internal/assets/model"
	teammodel "github.com/Thanhbinh1905/seta-training-system/internal/team/model"
	usermodel "github.com/Thanhbinh1905/seta-training-system/internal/user/model"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&usermodel.User{},
		&teammodel.Team{},
		&teammodel.TeamMember{},
		&teammodel.TeamManager{},
		&assetsmodel.Folder{},
		&assetsmodel.Note{},
		&assetsmodel.FolderShare{},
		&assetsmodel.NoteShare{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
