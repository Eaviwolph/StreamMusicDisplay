package structs

import "time"

type FileSaveStruct struct {
	Path    string
	Format  string
	Default string
}

type FileSaveConfStruct struct {
	Frequency time.Duration
	SaveTxt   bool
	SaveImg   bool
	ImgPath   string
	FileSaves []FileSaveStruct
}
