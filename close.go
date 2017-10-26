package sock

func Close(key string) {
	wbooll.Lock()
	wbools := wboolm[key]
	for _, wb := range wbools {
		close(wb.w)
	}
	delete(wboolm, key)
	wbooll.Unlock()

	rbooll.Lock()
	rbools := rboolm[key]
	for _, rb := range rbools {
		close(rb.r)
	}
	delete(rboolm, key)
	rbooll.Unlock()

	wstringl.Lock()
	wstrings := wstringm[key]
	for _, ws := range wstrings {
		close(ws.w)
	}
	delete(wstringm, key)
	wstringl.Unlock()

	rstringl.Lock()
	rstrings := rstringm[key]
	for _, rs := range rstrings {
		close(rs.r)
	}
	delete(rstringm, key)
	rstringl.Unlock()

	wintl.Lock()
	wints := wintm[key]
	for _, wi := range wints {
		close(wi.w)
	}
	delete(wintm, key)
	wintl.Unlock()

	rintl.Lock()
	rints := rintm[key]
	for _, ri := range rints {
		close(ri.r)
	}
	delete(rintm, key)
	rintl.Unlock()

	wint8l.Lock()
	wint8s := wint8m[key]
	for _, wi := range wint8s {
		close(wi.w)
	}
	delete(wint8m, key)
	wint8l.Unlock()

	rint8l.Lock()
	rint8s := rint8m[key]
	for _, ri := range rint8s {
		close(ri.r)
	}
	delete(rint8m, key)
	rint8l.Unlock()

	wint16l.Lock()
	wint16s := wint16m[key]
	for _, wi := range wint16s {
		close(wi.w)
	}
	delete(wint16m, key)
	wint16l.Unlock()

	rint16l.Lock()
	rint16s := rint16m[key]
	for _, ri := range rint16s {
		close(ri.r)
	}
	delete(rint16m, key)
	rint16l.Unlock()

	wint32l.Lock()
	wint32s := wint32m[key]
	for _, wi := range wint32s {
		close(wi.w)
	}
	delete(wint32m, key)
	wint32l.Unlock()

	rint32l.Lock()
	rint32s := rint32m[key]
	for _, ri := range rint32s {
		close(ri.r)
	}
	delete(rint32m, key)
	rint32l.Unlock()

	wint64l.Lock()
	wint64s := wint64m[key]
	for _, wi := range wint64s {
		close(wi.w)
	}
	delete(wint64m, key)
	wint64l.Unlock()

	rint64l.Lock()
	rint64s := rint64m[key]
	for _, ri := range rint64s {
		close(ri.r)
	}
	delete(rint64m, key)
	rint64l.Unlock()

	wuintl.Lock()
	wuints := wuintm[key]
	for _, wu := range wuints {
		close(wu.w)
	}
	delete(wuintm, key)
	wuintl.Unlock()

	ruintl.Lock()
	ruints := ruintm[key]
	for _, ru := range ruints {
		close(ru.r)
	}
	delete(ruintm, key)
	ruintl.Unlock()

	wuint8l.Lock()
	wuint8s := wuint8m[key]
	for _, wu := range wuint8s {
		close(wu.w)
	}
	delete(wuint8m, key)
	wuint8l.Unlock()

	ruint8l.Lock()
	ruint8s := ruint8m[key]
	for _, ru := range ruint8s {
		close(ru.r)
	}
	delete(ruint8m, key)
	ruint8l.Unlock()

	wuint16l.Lock()
	wuint16s := wuint16m[key]
	for _, wu := range wuint16s {
		close(wu.w)
	}
	delete(wuint16m, key)
	wuint16l.Unlock()

	ruint16l.Lock()
	ruint16s := ruint16m[key]
	for _, ru := range ruint16s {
		close(ru.r)
	}
	delete(ruint16m, key)
	ruint16l.Unlock()

	wuint32l.Lock()
	wuint32s := wuint32m[key]
	for _, wu := range wuint32s {
		close(wu.w)
	}
	delete(wuint32m, key)
	wuint32l.Unlock()

	ruint32l.Lock()
	ruint32s := ruint32m[key]
	for _, ru := range ruint32s {
		close(ru.r)
	}
	delete(ruint32m, key)
	ruint32l.Unlock()

	wuint64l.Lock()
	wuint64s := wuint64m[key]
	for _, wu := range wuint64s {
		close(wu.w)
	}
	delete(wuint64m, key)
	wuint64l.Unlock()

	ruint64l.Lock()
	ruint64s := ruint64m[key]
	for _, ru := range ruint64s {
		close(ru.r)
	}
	delete(ruint64m, key)
	ruint64l.Unlock()

	wbytel.Lock()
	wbytes := wbytem[key]
	for _, wb := range wbytes {
		close(wb.w)
	}
	delete(wbytem, key)
	wbytel.Unlock()

	rbytel.Lock()
	rbytes := rbytem[key]
	for _, rb := range rbytes {
		close(rb.r)
	}
	delete(rbytem, key)
	rbytel.Unlock()

	wbytesl.Lock()
	wbytess := wbytesm[key]
	for _, wb := range wbytess {
		close(wb.w)
	}
	delete(wbytesm, key)
	wbytesl.Unlock()

	rbytesl.Lock()
	rbytess := rbytesm[key]
	for _, rb := range rbytess {
		close(rb.r)
	}
	delete(rbytesm, key)
	rbytesl.Unlock()

	wrunel.Lock()
	wrunes := wrunem[key]
	for _, wr := range wrunes {
		close(wr.w)
	}
	delete(wrunem, key)
	wrunel.Unlock()

	rrunel.Lock()
	rrunes := rrunem[key]
	for _, rr := range rrunes {
		close(rr.r)
	}
	delete(rrunem, key)
	rrunel.Unlock()

	wfloat32l.Lock()
	wfloat32s := wfloat32m[key]
	for _, wf := range wfloat32s {
		close(wf.w)
	}
	delete(wfloat32m, key)
	wfloat32l.Unlock()

	rfloat32l.Lock()
	rfloat32s := rfloat32m[key]
	for _, rf := range rfloat32s {
		close(rf.r)
	}
	delete(rfloat32m, key)
	rfloat32l.Unlock()

	wfloat64l.Lock()
	wfloat64s := wfloat64m[key]
	for _, wf := range wfloat64s {
		close(wf.w)
	}
	delete(wfloat64m, key)
	wfloat64l.Unlock()

	rfloat64l.Lock()
	rfloat64s := rfloat64m[key]
	for _, rf := range rfloat64s {
		close(rf.r)
	}
	delete(rfloat64m, key)
	rfloat64l.Unlock()

	werrorl.Lock()
	werrors := werrorm[key]
	for _, we := range werrors {
		close(we.w)
	}
	delete(werrorm, key)
	werrorl.Unlock()

	rerrorl.Lock()
	rerrors := rerrorm[key]
	for _, re := range rerrors {
		close(re.r)
	}
	delete(rerrorm, key)
	rerrorl.Unlock()
}
