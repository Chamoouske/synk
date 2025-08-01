package domain

type FileInfo struct {
	ModTime string      `json:"mod_time"`
	Size    int64       `json:"size"`
	Blocks  []BlockInfo `json:"blocks"`
}

type BlockInfo struct {
	Hash string `json:"hash"`
}
