package documents

func (m *DocumentMgr) GetThumb(id int) ([]byte, error) {
	return m.paperless.GetDocumentThumb(id)
}
