package forms

type errors map[string][]string

//form errors to be accessible

func (e errors) Add(tagId, message string) {
	e[tagId] = append(e[tagId], message)
}

func (e errors) GetError(tagId string) string {
	es := e[tagId]
	if len(es) == 0 {
		return ""
	} else {
		return es[0]
	}

}
