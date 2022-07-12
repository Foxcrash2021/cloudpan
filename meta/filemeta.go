package meta

type FileMeta struct {
	FileSha1 string
	FileName string
	Filesize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func GetFileMeta(FileSha1 string) FileMeta {
	return fileMetas[FileSha1]
}

//删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
