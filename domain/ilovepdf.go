package domain

type File struct {
	ServerFilename string `json:"server_filename" validate:"required"`
	Filename       string `json:"filename" validate:"required"`
	Rotate         int    `json:"rotate"`
	Password       string `json:"password"`
}

type Element struct {
	Type        string `json:"type" validate:"required"`
	Pages       string `json:"pages" validate:"required"`
	Zindex      string `json:"zindex" validate:"required"`
	Dimensions  string `json:"dimensions" validate:"required"`
	Coordinates string `json:"coordinates" validate:"required"`
	Rotation    string `json:"rotation"`
	Opacity     string `json:"opacity"`

	// Type image, svg
	ServerFilename string `json:"server_filename"`

	// Type text
	Text          string `json:"text"`
	TextAlign     string `json:"text_align"`
	FontFamily    string `json:"font_family"`
	FontSize      int    `json:"font_size"`
	FontStyle     string `json:"font_style"`
	FontColor     string `json:"font_color"`
	LetterSpacing string `json:"letter_spacing"`
	UnderlineText string `json:"underline_text"`
}

type IlovepdfRequest struct {
	Task  string `json:"task"`
	Tool  string `json:"tool" validate:"required"`
	Files []File `json:"files" validate:"required"`

	// compress
	CompressionLevel string `json:"compression_level"`

	// Split
	SplitMode   string `json:"split_mode"`
	Ranges      string `json:"ranges"`
	FixedRange  int    `json:"fixed_range"`
	RemovePages string `json:"remove_pages"`
	MergeAfter  bool   `json:"merge_after"`

	// Edit
	Elements []Element `json:"elements"`
}

type IlovepdfReponse struct {
	DownloadFilename string `json:"download_filename"`
	Filesize         int    `json:"filesize"`
	OutputFilesize   int    `json:"output_filesize"`
	OutputFilenumber int    `json:"output_filenumber"`
	OutputExtensions string `json:"output_extensions"`
	Timer            string `json:"timer"`
	Status           string `json:"status"`
}
