package edit

func (e *Editor) GetLimitedFileName(fileName string) string {
	fn := []rune(fileName)
	if e.Settings.Rslt.MaxLen == nil || len(fn) <= *e.Settings.Rslt.MaxLen {
		return fileName
	}

	return string(fn[0:*e.Settings.Rslt.MaxLen-3]) + "..."
}
