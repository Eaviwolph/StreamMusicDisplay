package structs

type FileSaveStruct struct {
	Path    string
	Format  string
	Default string
}

type FileSaveConfStruct struct {
	Frequency int
	SaveImg   bool
	ImgPath   string
	FileSaves []FileSaveStruct
}
