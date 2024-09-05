package dtos

type GetFileDto struct {
	OK     bool          `json:"ok"`
	Result GetFileResult `json:"result"`
}

type GetFileResult struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`
}
