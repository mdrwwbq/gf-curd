package src

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
)

var (
	templateContent = map[string]string{}
)

func init() {
	getTemplateData()
}

// GetDefaultTemplateContent 获取模板内容
func GetDefaultTemplateContent(fileName string) string {
	return templateContent[filepath.Base(fileName)]
}

func getTemplateData() {
	templateData := `eyJSRUFETUUubWQiOiIjIOaooeadv+ebruW9lVxuXG4tIGFwaS50ZW1wbGF0ZVxuLSBjb25zdHMudGVtcGxhdGVcbi0gY29udHJvbGxlci50ZW1wbGF0ZVxuLSBkYi50ZW1wbGF0ZVxuLSBpMThuLnRlbXBsYXRlXG4tIGxhbmd1YWdlLnRlbXBsYXRlXG4tIGxvZ2ljLnRlbXBsYXRlXG4tIHV0aWxpdHkudGVtcGxhdGVcbiIsImFwaS50ZW1wbGF0ZSI6InBhY2thZ2Uge3tNb2R1bGVOYW1lfX1cclxuXHJcblxyXG5pbXBvcnQgKFxyXG5cdFwiZ2l0aHViLmNvbS9nb2dmL2dmL3YyL2ZyYW1lL2dcIlxyXG4pXHJcblxyXG50eXBlIHt7VXBwZXJUb29sc05hbWV9fUxpc3RSZXEgc3RydWN0IHtcclxuXHRnLk1ldGEgICBgcGF0aDpcIi97e0xvd2VyVG9vbHNOYW1lfX0vbGlzdFwiIHRhZ3M6XCJ7e01vZHVsZU5hbWV9fSB7e1Rvb2xzTmFtZX19XCIgbWV0aG9kOlwiZ2V0XCIgc3VtbWFyeTpcInt7VG9vbHNOYW1lfX3liJfooahcImBcclxuXHJcblx0dXRpbGl0eS5Vbml0UGFnZVxyXG59XHJcbnR5cGUge3tVcHBlclRvb2xzTmFtZX19Rm9ybSBzdHJ1Y3Qge1xyXG5cdE5hbWUgICAgIHN0cmluZyBganNvbjpcIm5hbWVcIiBkYzpcIuWQjeensFwiIHY6XCJyZXF1aXJlZCPlkI3np7Dlv4XloatcImBcclxufVxyXG50eXBlIHt7VXBwZXJUb29sc05hbWV9fUFkZFJlcSBzdHJ1Y3Qge1xyXG5cdGcuTWV0YSBgcGF0aDpcIi97e0xvd2VyVG9vbHNOYW1lfX0vYWRkXCIgdGFnczpcInt7TW9kdWxlTmFtZX19IHt7VG9vbHNOYW1lfX1cIiBtZXRob2Q6XCJwb3N0XCIgc3VtbWFyeTpcIuaWsOWinnt7VG9vbHNOYW1lfX1cImBcclxuXHR7e1VwcGVyVG9vbHNOYW1lfX1Gb3JtXHJcbn1cclxuXHJcbnR5cGUge3tVcHBlclRvb2xzTmFtZX19RWRpdFJlcSBzdHJ1Y3Qge1xyXG5cdGcuTWV0YSBgcGF0aDpcIi97e0xvd2VyVG9vbHNOYW1lfX0vZWRpdFwiIHRhZ3M6XCJ7e01vZHVsZU5hbWV9fSB7e1Rvb2xzTmFtZX19XCIgbWV0aG9kOlwicG9zdFwiIHN1bW1hcnk6XCLkv67mlLl7e1Rvb2xzTmFtZX19XCJgXHJcblx0SWQgICAgIGludCBganNvbjpcImlkXCIgdjpcInJlcXVpcmVkI0lE5LiN6IO95Li656m6fG1pbjoxI2lk5Y+C5pWw5LiN6IO95bCP5LqOMFwiYFxyXG5cdHt7VXBwZXJUb29sc05hbWV9fUZvcm1cclxufVxyXG50eXBlIHt7VXBwZXJUb29sc05hbWV9fURlbGV0ZVJlcSBzdHJ1Y3Qge1xyXG5cdGcuTWV0YSBgcGF0aDpcIi97e0xvd2VyVG9vbHNOYW1lfX0vZGVsZXRlXCIgdGFnczpcInt7TW9kdWxlTmFtZX19IHt7VG9vbHNOYW1lfX1cIiBtZXRob2Q6XCJwb3N0XCIgc3VtbWFyeTpcIuWIoOmZpHt7VG9vbHNOYW1lfX1cImBcclxuXHRJZCAgICAgW11pbnQgYGpzb246XCJpZFwiIHY6XCJyZXF1aXJlZCNJROS4jeiDveS4uuepulwiYFxyXG59XHJcbnR5cGUge3tVcHBlclRvb2xzTmFtZX19TG9naWNEZWxldGVSZXEgc3RydWN0IHtcclxuXHRnLk1ldGEgYHBhdGg6XCIve3tMb3dlclRvb2xzTmFtZX19L2xvZ2ljRGVsZXRlXCIgdGFnczpcInt7TW9kdWxlTmFtZX19IHt7VG9vbHNOYW1lfX1cIiBtZXRob2Q6XCJwb3N0XCIgc3VtbWFyeTpcIumAu+i+keWIoOmZpHt7VG9vbHNOYW1lfX1cImBcclxuXHRJZCAgICAgW11pbnQgYGpzb246XCJpZFwiIHY6XCJyZXF1aXJlZCNJROS4jeiDveS4uuepulwiYFxyXG59XHJcbnR5cGUge3tVcHBlclRvb2xzTmFtZX19SW5mb1JlcSBzdHJ1Y3Qge1xyXG5cdGcuTWV0YSBgcGF0aDpcIi97e0xvd2VyVG9vbHNOYW1lfX0vaW5mb1wiIHRhZ3M6XCJ7e01vZHVsZU5hbWV9fSB7e1Rvb2xzTmFtZX19XCIgbWV0aG9kOlwiZ2V0XCIgc3VtbWFyeTpcInt7VG9vbHNOYW1lfX3or6bmg4VcImBcclxuXHRJZCAgICAgaW50IGBqc29uOlwiaWRcIiB2OlwicmVxdWlyZWQjSUTkuI3og73kuLrnqbp8bWluOjEjaWTlj4LmlbDkuI3og73lsI/kuo4wXCJgXHJcbn1cclxudHlwZSB7e1VwcGVyVG9vbHNOYW1lfX1FbmFibGVSZXEgc3RydWN0IHtcclxuXHRnLk1ldGEgYHBhdGg6XCIve3tMb3dlclRvb2xzTmFtZX19L2VuYWJsZVwiIHRhZ3M6XCJ7e01vZHVsZU5hbWV9fSB7e1Rvb2xzTmFtZX19XCIgbWV0aG9kOlwicG9zdFwiIHN1bW1hcnk6XCLlkK/nlKgv56aB55Soe3tUb29sc05hbWV9fVwiYFxyXG5cdElkICAgICAgW11pbnQgYGpzb246XCJpZFwiIHY6XCJyZXF1aXJlZCNJROS4jeiDveS4uuepulwiYFxyXG5cdEVuYWJsZSBpbnQgICAgYGpzb246XCJlbmFibGVcIiBkYzpcIueKtuaAgVwiICB2OlwicmVxdWlyZWR8aW46MSwyI+eKtuaAgeW/heWhq3znirbmgIHlj4LmlbDplJnor69cImBcclxufSIsImNvbnN0cy50ZW1wbGF0ZSI6InBhY2thZ2UgdXRpbGl0eVxuXG52YXIgKFxuICAgIEVuYWJsZSAgPSAxIC8vIOWQr+eUqFxuXHREaXNhYmxlID0gMiAvLyDnpoHnlKhcblx0SXNEZWwgICA9IDIgLy8g5Yig6ZmkXG5cdE5vcm1hbCAgPSAxIC8vIOato+W4uFxuXHRObyAgICAgID0gMiAvLyDlkKZcblx0WWVzICAgICA9IDEgLy8g5pivXG5cdFNob3cgICAgPSAxIC8vIOaYvuekulxuXHRIaWRlICAgID0gMiAvLyDpmpDol49cblxuXHRNYWxlICAgPSBcIueUt1wiXG5cdEZlbWFsZSA9IFwi5aWzXCJcbikiLCJjb250cm9sbGVyLnRlbXBsYXRlIjoicGFja2FnZSB7e1BBQ0tBR0V9fVxyXG5cclxuaW1wb3J0IChcclxuXHRcImNvbnRleHRcIlxyXG5cclxuXHRcImdpdGh1Yi5jb20vZ29nZi9nZi92Mi9mcmFtZS9nXCJcclxuKVxyXG5cclxudmFyIChcclxuXHR7e1VwcGVyVG9vbHNOYW1lfX0gPSBje3tVcHBlclRvb2xzTmFtZX19e31cclxuKVxyXG5cclxudHlwZSBje3tVcHBlclRvb2xzTmFtZX19IHN0cnVjdCB7XHJcbn1cclxuXHJcbi8vIExpc3Qge3tUb29sc05hbWV9feWIl+ihqFxyXG5mdW5jIChjICpje3tVcHBlclRvb2xzTmFtZX19KSBMaXN0KGN0eCBjb250ZXh0LkNvbnRleHQsIHJlcSAqe3tNb2R1bGVOYW1lfX0ue3tVcHBlclRvb2xzTmFtZX19TGlzdFJlcSkgKHJlcyAqdXRpbGl0eS5Vbml0UmVzLCBlcnIgZXJyb3IpIHtcclxuXHRsaXN0LCBwYWdlLCBlcnIgOj0gc2VydmljZS57e1VwcGVyVG9vbHNOYW1lfX0oKS5MaXN0KGN0eClcclxuXHRpZiBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBuaWwsIGVyclxyXG5cdH1cclxuXHRyZXR1cm4gdXRpbGl0eS5SZXNwb25zZVBhZ2VUbygwLCBsaXN0LCBwYWdlLCB1dGlsaXR5LkwoY3R4LCBgU3VjY2Vzc2ApKSwgbmlsXHJcbn1cclxuXHJcbi8vIEFkZCDmlrDlop57e1Rvb2xzTmFtZX19XHJcbmZ1bmMgKGMgKmN7e1VwcGVyVG9vbHNOYW1lfX0pIEFkZChjdHggY29udGV4dC5Db250ZXh0LCByZXEgKnt7TW9kdWxlTmFtZX19Lnt7VXBwZXJUb29sc05hbWV9fUFkZFJlcSkgKHJlcyAqdXRpbGl0eS5Vbml0UmVzLCBlcnIgZXJyb3IpIHtcclxuXHRlcnIgPSBzZXJ2aWNlLnt7VXBwZXJUb29sc05hbWV9fSgpLkFkZChjdHgpXHJcblx0aWYgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsLCBlcnJcclxuXHR9XHJcblx0cmV0dXJuIHV0aWxpdHkuUmVzcG9uc2VUbygwLCBuaWwsIHV0aWxpdHkuTChjdHgsIGBTdWNjZXNzYCkpLCBuaWxcclxufVxyXG5cclxuLy8gRWRpdCDkv67mlLl7e1Rvb2xzTmFtZX19XHJcbmZ1bmMgKGMgKmN7e1VwcGVyVG9vbHNOYW1lfX0pIEVkaXQoY3R4IGNvbnRleHQuQ29udGV4dCwgcmVxICp7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1FZGl0UmVxKSAocmVzICp1dGlsaXR5LlVuaXRSZXMsIGVyciBlcnJvcikge1xyXG5cdGVyciA9IHNlcnZpY2Uue3tVcHBlclRvb2xzTmFtZX19KCkuRWRpdChjdHgpXHJcblx0aWYgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsLCBlcnJcclxuXHR9XHJcblx0cmV0dXJuIHV0aWxpdHkuUmVzcG9uc2VUbygwLCBuaWwsIHV0aWxpdHkuTChjdHgsIGBTdWNjZXNzYCkpLCBuaWxcclxufVxyXG5cclxuLy8gRGVsZXRlIOWIoOmZpHt7VG9vbHNOYW1lfX1cclxuZnVuYyAoYyAqY3t7VXBwZXJUb29sc05hbWV9fSkgRGVsZXRlKGN0eCBjb250ZXh0LkNvbnRleHQsIHJlcSAqe3tNb2R1bGVOYW1lfX0ue3tVcHBlclRvb2xzTmFtZX19RGVsZXRlUmVxKSAocmVzICp1dGlsaXR5LlVuaXRSZXMsIGVyciBlcnJvcikge1xyXG5cdGVyciA9IHNlcnZpY2Uue3tVcHBlclRvb2xzTmFtZX19KCkuRGVsZXRlKGN0eClcclxuXHRpZiBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBuaWwsIGVyclxyXG5cdH1cclxuXHRyZXR1cm4gdXRpbGl0eS5SZXNwb25zZVRvKDAsIG5pbCwgdXRpbGl0eS5MKGN0eCwgYFN1Y2Nlc3NgKSksIG5pbFxyXG59XHJcblxyXG4vLyBMb2dpY0RlbGV0ZSDpgLvovpHliKDpmaR7e1Rvb2xzTmFtZX19XHJcbmZ1bmMgKGMgKmN7e1VwcGVyVG9vbHNOYW1lfX0pIExvZ2ljRGVsZXRlKGN0eCBjb250ZXh0LkNvbnRleHQsIHJlcSAqe3tNb2R1bGVOYW1lfX0ue3tVcHBlclRvb2xzTmFtZX19TG9naWNEZWxldGVSZXEpIChyZXMgKnV0aWxpdHkuVW5pdFJlcywgZXJyIGVycm9yKSB7XHJcblx0ZXJyID0gc2VydmljZS57e1VwcGVyVG9vbHNOYW1lfX0oKS5Mb2dpY0RlbGV0ZShjdHgpXHJcblx0aWYgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsLCBlcnJcclxuXHR9XHJcblx0cmV0dXJuIHV0aWxpdHkuUmVzcG9uc2VUbygwLCBuaWwsIHV0aWxpdHkuTChjdHgsIGBTdWNjZXNzYCkpLCBuaWxcclxufVxyXG5cclxuLy8gSW5mbyB7e1Rvb2xzTmFtZX196K+m5oOFXHJcbmZ1bmMgKGMgKmN7e1VwcGVyVG9vbHNOYW1lfX0pIEluZm8oY3R4IGNvbnRleHQuQ29udGV4dCwgcmVxICp7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1JbmZvUmVxKSAocmVzICp1dGlsaXR5LlVuaXRSZXMsIGVyciBlcnJvcikge1xyXG5cdGluZm8sIGVyciA6PSBzZXJ2aWNlLnt7VXBwZXJUb29sc05hbWV9fSgpLkluZm8oY3R4KVxyXG5cdGlmIGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIG5pbCwgZXJyXHJcblx0fVxyXG5cdHJldHVybiB1dGlsaXR5LlJlc3BvbnNlVG8oMCwgaW5mbywgdXRpbGl0eS5MKGN0eCwgYFN1Y2Nlc3NgKSksIG5pbFxyXG59XHJcblxyXG4vLyBFbmFibGUg5ZCv55SoL+emgeeUqHt7VG9vbHNOYW1lfX1cclxuZnVuYyAoYyAqY3t7VXBwZXJUb29sc05hbWV9fSkgRW5hYmxlKGN0eCBjb250ZXh0LkNvbnRleHQsIHJlcSAqe3tNb2R1bGVOYW1lfX0ue3tVcHBlclRvb2xzTmFtZX19RW5hYmxlUmVxKSAocmVzICp1dGlsaXR5LlVuaXRSZXMsIGVyciBlcnJvcikge1xyXG5cdGVyciA9IHNlcnZpY2Uue3tVcHBlclRvb2xzTmFtZX19KCkuRW5hYmxlKGN0eClcclxuXHRpZiBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBuaWwsIGVyclxyXG5cdH1cclxuXHRyZXR1cm4gdXRpbGl0eS5SZXNwb25zZVRvKDAsIG5pbCwgdXRpbGl0eS5MKGN0eCwgYFN1Y2Nlc3NgKSksIG5pbFxyXG59IiwiZGIudGVtcGxhdGUiOiJwYWNrYWdlIHV0aWxpdHlcbmltcG9ydCAoXG5cdFwiY29udGV4dFwiXG5cdFwiZ2l0aHViLmNvbS9nb2dmL2dmL3YyL2ZyYW1lL2dcIlxuKVxuLy8gQ2hlY2tEYXRhRXhpc3Qg5qOA5p+l5pWw5o2u5piv5ZCm5a2Y5ZyoXG5mdW5jIENoZWNrRGF0YUV4aXN0KGN0eCBjb250ZXh0LkNvbnRleHQsIHRhYmxlIHN0cmluZywgd2hlcmUgaW50ZXJmYWNle30sIGNhbGxlciAuLi4gZnVuYyhjdHggY29udGV4dC5Db250ZXh0LCBjb3VudCBpbnQpIChib29sLCBlcnJvcikgKSAoYm9vbCwgZXJyb3IpIHtcblx0aWYgY291bnQsIGVyciA6PSBnLk1vZGVsKHRhYmxlKS5XaGVyZSh3aGVyZSkuQ291bnQoKTsgZXJyICE9IG5pbCB7XG5cdFx0cmV0dXJuIGZhbHNlLCBlcnJcblx0fSBlbHNlIHtcblx0ICAgIGlmIGxlbihjYWxsZXIpIFx1MDAzZSAwIHtcblx0ICAgICAgICByZXR1cm4gY2FsbGVyWzBdKGN0eCwgY291bnQpXG5cdCAgICB9XG5cdFx0cmV0dXJuIGNvdW50IFx1MDAzZSAwLCBuaWxcblx0fVxufSIsImkxOG4udGVtcGxhdGUiOiJ7e1VwcGVyTW9kdWxlTmFtZX19e3tVcHBlclRvb2xzTmFtZX19Q29kZUlzRXhpc3Q9XCLigJwlduKAnOW3sue7j+WtmOWcqFwiXHJcbnt7VXBwZXJNb2R1bGVOYW1lfX17e1VwcGVyVG9vbHNOYW1lfX1EYXRhSXNOb3RFeGlzdD1cIuaVsOaNruS4jeWtmOWcqGlkPSV2XCJcclxue3tVcHBlck1vZHVsZU5hbWV9fXt7VXBwZXJUb29sc05hbWV9fURlbGV0ZURhdGFOdW1Jc05vdEVxdWFsPVwi5Yig6Zmk5pWw5o2u5pWw6YePOiV25LiO5b6F5Yig6Zmk5pWw6YePOiV25LiN5LiA6Ie0XCJcclxue3tVcHBlck1vZHVsZU5hbWV9fXt7VXBwZXJUb29sc05hbWV9fUVuYWJsZWREYXRhTnVtSXNOb3RFcXVhbD1cIuemgeeUqOaVsOaNruaVsOmHjzolduS4juW+heemgeeUqOaVsOmHjzolduS4jeS4gOiHtFwiXHJcbnt7VXBwZXJNb2R1bGVOYW1lfX17e1VwcGVyVG9vbHNOYW1lfX1NdWx0aURlbGV0ZUlkSXNFbXB0eT1cIuaJuemHj+WIoOmZpO+8jElE5LiN6IO95Li656m6XCIiLCJsYW5ndWFnZS50ZW1wbGF0ZSI6InBhY2thZ2UgdXRpbGl0eVxuXG5pbXBvcnQgKFxuXHRcImNvbnRleHRcIlxuXHRcImZtdFwiXG5cblx0XCJnaXRodWIuY29tL2dvZ2YvZ2YvdjIvZnJhbWUvZ1wiXG4pXG5cbi8vIEwg6I635b6X6K+t6KiA5YyF5YC8XG5mdW5jIEwoY3R4IGNvbnRleHQuQ29udGV4dCwga2V5IHN0cmluZykgc3RyaW5nIHtcblx0bFN0ciA6PSBnLkkxOG4oKS5UKGN0eCwga2V5KVxuXHRpZiBsU3RyID09IFwiXCIge1xuXHRcdHJldHVybiBmbXQuU3ByaW50ZihcIm5vdCBmb3VuZCBsYW5ndWFnZTolc1wiLCBrZXkpXG5cdH1cblx0cmV0dXJuIGxTdHJcbn1cblxuLy8gVCDliqjmgIHljIXor63oqIDljIUgVChjdHgsIGBHRiBzYXlzOiB7I2hlbGxvfXsjd29ybGR9IWApKSAtLVx1MDAzZSBHRiBzYXlzOiDlk4jllr3kuJbnlYxcbmZ1bmMgVChjdHggY29udGV4dC5Db250ZXh0LCBjb250ZW50IHN0cmluZykgc3RyaW5nIHtcblx0bFN0ciA6PSBnLkkxOG4oKS5UcmFuc2xhdGUoY3R4LCBjb250ZW50KVxuXHRpZiBsU3RyID09IFwiXCIge1xuXHRcdHJldHVybiBmbXQuU3ByaW50ZihcIm5vdCBmb3VuZCBsYW5ndWFnZTolc1wiLCBjb250ZW50KVxuXHR9XG5cdHJldHVybiBsU3RyXG59XG5cbi8vIFRmIOWPr+i9rOaNouivreiogOWMheS4reWtmOWcqOWPmOmHjyBUZihjdHgsIGB7I2hlbGxvfWAsIOW8oOS4iSzmnY7lm5spIC0tXHUwMDNlICBoZWxsbyDlvKDkuIks5p2O5Zub44CCICB7I2hlbGxvfSAtLVx1MDAzZSBoZWxsbz1cIuWTiOWWvSMlcywjJXPjgIJcIlxuZnVuYyBUZihjdHggY29udGV4dC5Db250ZXh0LCBmb3JtYXQgc3RyaW5nLCB2YWx1ZXMgLi4uaW50ZXJmYWNle30pIHN0cmluZyB7XG5cdGxTdHIgOj0gZy5JMThuKCkuVGYoY3R4LCBmb3JtYXQsIHZhbHVlcy4uLilcblx0aWYgbFN0ciA9PSBcIlwiIHtcblx0XHRyZXR1cm4gZm10LlNwcmludGYoXCJub3QgZm91bmQgbGFuZ3VhZ2U6JXNcIiwgZm9ybWF0KVxuXHR9XG5cdHJldHVybiBsU3RyXG59IiwibG9naWMudGVtcGxhdGUiOiJwYWNrYWdlIHt7VG9vbHNOYW1lfX1cclxuXHJcbmltcG9ydCAoXHJcblx0XCJjb250ZXh0XCJcclxuXHRcImRhdGFiYXNlL3NxbFwiXHJcblx0XCJlcnJvcnNcIlxyXG5cdFwiZm10XCJcclxuXHRcImdpdGh1Yi5jb20vZ29nZi9nZi92Mi9kYXRhYmFzZS9nZGJcIlxyXG5cclxuXHRcImdpdGh1Yi5jb20vZ29nZi9nZi92Mi9mcmFtZS9nXCJcclxuKVxyXG5cclxudHlwZSAoXHJcblx0c3t7VXBwZXJUb29sc05hbWV9fSBzdHJ1Y3R7fVxyXG4pXHJcblxyXG5mdW5jIGluaXQoKSB7XHJcblx0c2VydmljZS5SZWdpc3Rlcnt7VXBwZXJUb29sc05hbWV9fShOZXcoKSlcclxufVxyXG5cclxuZnVuYyBOZXcoKSAqc3t7VXBwZXJUb29sc05hbWV9fSB7XHJcblx0cmV0dXJuIFx1MDAyNnN7e1VwcGVyVG9vbHNOYW1lfX17fVxyXG59XHJcbi8vIHRhYmxlIOiOt+W+l+W9k+WJjeWKn+iDveihqOWQjVxyXG5mdW5jIChzICpze3tVcHBlclRvb2xzTmFtZX19KSB0YWJsZSgpIHN0cmluZyB7XHJcbiAgICAvLyB0b2RvIHdyaXRlIHNwZWNpZmljIHRhYmxlIGhlcmVcclxuICAgIHJldHVybiBcIm1hbnVhbF9jaGFuZ2VzX3RhYmxlX25hbWVcIlxyXG59XHJcbi8vIG1vZGVsIOiOt+W+l+W9k+WJjeWKn+iDvU1vZGVsXHJcbmZ1bmMgKHMgKnN7e1VwcGVyVG9vbHNOYW1lfX0pIG1vZGVsKCkgKmdkYi5Nb2RlbCB7XHJcbiAgICByZXR1cm4gZy5Nb2RlbChzLnRhYmxlKCkpXHJcbn1cclxuLy8gZ2V0UHJpbWFyeUtleSDojrflvpfooajoh6rlop5JROWtl+autVxyXG5mdW5jIChzICpze3tVcHBlclRvb2xzTmFtZX19KSBnZXRQcmltYXJ5S2V5RmllbGQoKSBzdHJpbmcge1xyXG5cdC8vIHRvZG8gd3JpdGUgc3BlY2lmaWMgdGFibGUgcHJpbWFyeSBrZXkgaGVyZVxyXG5cdHJldHVybiBcImlkXCJcclxufVxyXG5cclxuLy8gZ2V0TG9naWNEZWxldGVGaWVsZCDojrflvpfooajpgLvovpHliKDpmaTlrZfmrrVcclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgZ2V0TG9naWNEZWxldGVGaWVsZCgpIHN0cmluZyB7XHJcblx0Ly8gdG9kbyB3cml0ZSBzcGVjaWZpYyB0YWJsZSBsb2dpYyBkZWxldGUgZmllbGQgaGVyZVxyXG5cdHJldHVybiBcImlzX2RlbFwiXHJcbn1cclxuXHJcbi8vIGdldExvZ2ljRGVsZXRlRmllbGQg6I635b6X6KGo5ZCv55SoL+emgeeUqOWtl+autVxyXG5mdW5jIChzICpze3tVcHBlclRvb2xzTmFtZX19KSBnZXRFbmFibGVGaWVsZCgpIHN0cmluZyB7XHJcblx0Ly8gdG9kbyB3cml0ZSBzcGVjaWZpYyB0YWJsZSBlbmFibGVkIGZpZWxkICBoZXJlXHJcblx0cmV0dXJuIFwiZW5hYmxlXCJcclxufVxyXG4vLyBMaXN0IHt7VG9vbHNOYW1lfX3liJfooahcclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgTGlzdChjdHggY29udGV4dC5Db250ZXh0KSAobGlzdCBbXSplbnRpdHkue3tVcHBlclRvb2xzTmFtZX19LCBwYWdlICp1dGlsaXR5LlVuaXRQYWdlUmVzLCBlcnIgZXJyb3IpIHtcclxuXHR2YXIgKFxyXG5cdFx0ciAgPSBnLlJlcXVlc3RGcm9tQ3R4KGN0eClcclxuXHRcdGluIHt7TW9kdWxlTmFtZX19Lnt7VXBwZXJUb29sc05hbWV9fUxpc3RSZXFcclxuXHQpXHJcblx0aWYgZXJyIDo9IHIuUGFyc2UoXHUwMDI2aW4pOyBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBuaWwsIG5pbCwgZXJyXHJcblx0fVxyXG5cclxuXHRsaXN0ID0gbWFrZShbXSplbnRpdHkue3tVcHBlclRvb2xzTmFtZX19LCAwKVxyXG5cdG1vZGVsIDo9IHMubW9kZWwoKS5DdHgoY3R4KVxyXG5cdC8vIOe7n+iuoeaVsOmHj1xyXG5cdGlmIHRvdGFsLCBlcnIgOj0gbW9kZWwuQ291bnQoKTsgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsLCBuaWwsIGVyclxyXG5cdH0gZWxzZSB7XHJcbiAgICAgICAgLy8g55Sf5oiQ5YiG6aG1XHJcbiAgICAgICAgcGFnZSA9IHV0aWxpdHkuR2V0UGFnZShjdHgsIGluLkN1cnJlbnRQYWdlLCBpbi5QZXJQYWdlLCB0b3RhbClcclxuICAgICAgICAvLyDmn6Xor6LmlbDmja5cclxuICAgICAgICBpZiBlcnIgPSBtb2RlbC5QYWdlKHBhZ2UuQ3VycmVudFBhZ2UsIHBhZ2UuUGVyUGFnZSkuT3JkZXJEZXNjKHMuZ2V0UHJpbWFyeUtleUZpZWxkKCkpLlNjYW4oXHUwMDI2bGlzdCk7IGVyciAhPSBuaWwgXHUwMDI2XHUwMDI2ICFlcnJvcnMuSXMoZXJyLCBzcWwuRXJyTm9Sb3dzKSB7XHJcbiAgICAgICAgICAgIHJldHVybiBuaWwsIG5pbCwgZXJyXHJcbiAgICAgICAgfVxyXG4gICAgICAgIHJldHVybiBsaXN0LCBwYWdlLCBuaWxcclxuXHR9XHJcbn1cclxuXHJcbi8vIEFkZCDmlrDlop57e1Rvb2xzTmFtZX19XHJcbmZ1bmMgKHMgKnN7e1VwcGVyVG9vbHNOYW1lfX0pIEFkZChjdHggY29udGV4dC5Db250ZXh0KSAoZXJyIGVycm9yKSB7XHJcblx0dmFyIChcclxuXHRcdHIgID0gZy5SZXF1ZXN0RnJvbUN0eChjdHgpXHJcblx0XHRpbiB7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1BZGRSZXFcclxuXHQpXHJcblx0aWYgZXJyIDo9IHIuUGFyc2UoXHUwMDI2aW4pOyBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBuaWxcclxuXHR9XHJcblx0Ly8gY2hlY2sgbmFtZSBpcyBleGlzdHNcclxuXHRleGlzdCwgZXJyIDo9IHV0aWxpdHkuQ2hlY2tEYXRhRXhpc3QoY3R4LCBzLnRhYmxlKCksIGcuTWFwe1xyXG5cdFx0XCJuYW1lXCI6IGluLk5hbWUsXHJcblx0fSlcclxuXHRpZiBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBlcnJcclxuXHR9XHJcblx0Ly8gZGF0YSBpcyBleGlzdHNcclxuXHRpZiBleGlzdCB7XHJcblx0XHRyZXR1cm4gZXJyb3JzLk5ldyh1dGlsaXR5LlRmKGN0eCwgYHsje3tVcHBlck1vZHVsZU5hbWV9fXt7VXBwZXJUb29sc05hbWV9fUNvZGVJc0V4aXN0fWAsIGluLk5hbWUpKVxyXG5cdH1cclxuXHRkYXRhIDo9IGcuTWFwe1xyXG5cdCAgICBcIm5hbWVcIjogaW4uTmFtZSxcclxuXHQgICAgLy8gdG9kbyB3cml0ZSBzcGVjaWZpYyBkYXRhIHN0cnVjdCBsb2dpYyBoZXJlXHJcblx0fVxyXG5cdC8vIGluZXJ0IGRhdGEgdG8gdGFibGVcclxuICAgIF8sIGVyciA9IHMubW9kZWwoKS5DdHgoY3R4KS5EYXRhKGRhdGEpLkluc2VydCgpXHJcblx0cmV0dXJuIGVyclxyXG59XHJcblxyXG4vLyBFZGl0IOS/ruaUuXt7VG9vbHNOYW1lfX1cclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgRWRpdChjdHggY29udGV4dC5Db250ZXh0KSAoZXJyIGVycm9yKSB7XHJcblx0dmFyIChcclxuXHRcdHIgID0gZy5SZXF1ZXN0RnJvbUN0eChjdHgpXHJcblx0XHRpbiB7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1FZGl0UmVxXHJcblx0KVxyXG5cdGlmIGVyciA6PSByLlBhcnNlKFx1MDAyNmluKTsgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsXHJcblx0fVxyXG5cdC8vIGNoZWNrIGRhdGEgaXMgZXhpc3RzXHJcblx0aW5mbywgZXJyIDo9IHMuR2V0KGN0eCwgZy5NYXB7XHJcblx0XHRzLmdldFByaW1hcnlLZXlGaWVsZCgpOiBpbi5JZCxcclxuXHR9KVxyXG5cdGlmIGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIGVyclxyXG5cdH1cclxuXHRpZiBpbmZvLklkIFx1MDAzYz0gMCB7XHJcblx0XHRyZXR1cm4gZXJyb3JzLk5ldyh1dGlsaXR5LlRmKGN0eCwgYHsje3tVcHBlck1vZHVsZU5hbWV9fXt7VXBwZXJUb29sc05hbWV9fURhdGFJc05vdEV4aXN0fWAsIGluLklkKSlcclxuXHR9XHJcblx0Ly8gY2hlY2sgZGF0YSBpcyBleGlzdHNcclxuXHRleGlzdCwgZXJyIDo9IHV0aWxpdHkuQ2hlY2tEYXRhRXhpc3QoY3R4LCBzLnRhYmxlKCksIGcuTWFwe1xyXG5cdFx0XCJuYW1lXCI6ICAgICAgICAgICAgICAgICAgICAgaW4uTmFtZSxcclxuXHRcdGZtdC5TcHJpbnRmKGAlcyAhPWAsIHMuZ2V0UHJpbWFyeUtleUZpZWxkKCkpOiBpbi5JZCxcclxuXHR9KVxyXG5cdGlmIGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIGVyclxyXG5cdH1cclxuXHQvLyBpcyBleGlzdHNcclxuXHRpZiBleGlzdCB7XHJcblx0XHRyZXR1cm4gZXJyb3JzLk5ldyh1dGlsaXR5LlRmKGN0eCwgYHsje3tVcHBlck1vZHVsZU5hbWV9fXt7VXBwZXJUb29sc05hbWV9fUNvZGVJc0V4aXN0fWAsIGluLk5hbWUpKVxyXG5cdH1cclxuXHJcblx0ZGF0YSA6PSBnLk1hcHtcclxuXHQgICAgXCJuYW1lXCI6IGluLk5hbWUsXHJcblx0XHQvL2Rhby5MaXNNZW51LkNvbHVtbnMoKS5OYW1lOiBpbi5OYW1lLFxyXG5cdFx0Ly8gdG9kbyB3cml0ZSBzcGVjaWZpYyBkYXRhIHN0cnVjdCBsb2dpYyBoZXJlXHJcblx0fVxyXG5cdF8sIGVyciA9IHMubW9kZWwoKS5XaGVyZShzLmdldFByaW1hcnlLZXlGaWVsZCgpLCBpbi5JZCkuRGF0YShkYXRhKS5VcGRhdGUoKVxyXG5cdHJldHVybiBlcnJcclxufVxyXG5cclxuLy8gRGVsZXRlIOWIoOmZpHt7VG9vbHNOYW1lfX1cclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgRGVsZXRlKGN0eCBjb250ZXh0LkNvbnRleHQpIChlcnIgZXJyb3IpIHtcclxuXHR2YXIgKFxyXG5cdFx0ciAgPSBnLlJlcXVlc3RGcm9tQ3R4KGN0eClcclxuXHRcdGluIHt7TW9kdWxlTmFtZX19Lnt7VXBwZXJUb29sc05hbWV9fURlbGV0ZVJlcVxyXG5cdClcclxuXHRpZiBlcnIgOj0gci5QYXJzZShcdTAwMjZpbik7IGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIG5pbFxyXG5cdH1cclxuICAgIC8vIGNoZWNrIGRhdGEgaXMgZXhpc3RzXHJcbiAgICBfLCBlcnIgPSB1dGlsaXR5LkNoZWNrRGF0YUV4aXN0KGN0eCwgcy50YWJsZSgpLCBnLk1hcHtcclxuICAgICAgICBzLmdldFByaW1hcnlLZXlGaWVsZCgpOiBpbi5JZCxcclxuICAgIH0sIGZ1bmMoY3R4IGNvbnRleHQuQ29udGV4dCwgY291bnQgaW50KSAoYm9vbCwgZXJyb3IpIHtcclxuICAgICAgICBpZiBjb3VudCAhPSBsZW4oaW4uSWQpIHtcclxuICAgICAgICAgICAgcmV0dXJuIGZhbHNlLCBlcnJvcnMuTmV3KHV0aWxpdHkuVGYoY3R4LCBgeyNCYXNpY1VzZXJEZWxldGVEYXRhTnVtSXNOb3RFcXVhbH1gLCBjb3VudCwgbGVuKGluLklkKSkpXHJcbiAgICAgICAgfVxyXG4gICAgICAgIHJldHVybiB0cnVlLCBuaWxcclxuICAgIH0pXHJcbiAgICBpZiBlcnIgIT0gbmlsIHtcclxuICAgICAgICByZXR1cm4gZXJyXHJcbiAgICB9XHJcblx0XywgZXJyID0gcy5tb2RlbCgpLldoZXJlSW4ocy5nZXRQcmltYXJ5S2V5RmllbGQoKSwgaW4uSWQpLkRlbGV0ZSgpXHJcblx0cmV0dXJuIGVyclxyXG59XHJcbi8vIExvZ2ljRGVsZXRlIOmAu+i+keWIoOmZpHt7VG9vbHNOYW1lfX1cclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgTG9naWNEZWxldGUoY3R4IGNvbnRleHQuQ29udGV4dCkgKGVyciBlcnJvcikge1xyXG5cdHZhciAoXHJcblx0XHRyICA9IGcuUmVxdWVzdEZyb21DdHgoY3R4KVxyXG5cdFx0aW4ge3tNb2R1bGVOYW1lfX0ue3tVcHBlclRvb2xzTmFtZX19RGVsZXRlUmVxXHJcblx0KVxyXG5cdGlmIGVyciA6PSByLlBhcnNlKFx1MDAyNmluKTsgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsXHJcblx0fVxyXG5cdC8vIGNoZWNrIGRhdGEgaXMgZXhpc3RzXHJcbiAgICBfLCBlcnIgPSB1dGlsaXR5LkNoZWNrRGF0YUV4aXN0KGN0eCwgcy50YWJsZSgpLCBnLk1hcHtcclxuICAgICAgICBzLmdldFByaW1hcnlLZXlGaWVsZCgpOiBpbi5JZCxcclxuICAgICAgICBzLmdldExvZ2ljRGVsZXRlRmllbGQoKTogdXRpbGl0eS5Ob3JtYWwsXHJcbiAgICB9LCBmdW5jKGN0eCBjb250ZXh0LkNvbnRleHQsIGNvdW50IGludCkgKGJvb2wsIGVycm9yKSB7XHJcbiAgICAgICAgaWYgY291bnQgIT0gbGVuKGluLklkKSB7XHJcbiAgICAgICAgICAgIHJldHVybiBmYWxzZSwgZXJyb3JzLk5ldyh1dGlsaXR5LlRmKGN0eCwgYHsjQmFzaWNVc2VyRGVsZXRlRGF0YU51bUlzTm90RXF1YWx9YCwgY291bnQsIGxlbihpbi5JZCkpKVxyXG4gICAgICAgIH1cclxuICAgICAgICByZXR1cm4gdHJ1ZSwgbmlsXHJcbiAgICB9KVxyXG4gICAgaWYgZXJyICE9IG5pbCB7XHJcbiAgICAgICAgcmV0dXJuIGVyclxyXG4gICAgfVxyXG5cdF8sIGVyciA9IHMubW9kZWwoKS5XaGVyZUluKHMuZ2V0UHJpbWFyeUtleUZpZWxkKCksIGluLklkKS5EYXRhKGcuTWFwe1xyXG5cdCAgICBzLmdldExvZ2ljRGVsZXRlRmllbGQoKTogdXRpbGl0eS5Jc0RlbCxcclxuXHR9KS5VcGRhdGUoKVxyXG5cdHJldHVybiBlcnJcclxufVxyXG5cclxuLy8gRW5hYmxlIOWQr+eUqC/npoHnlKh7e1Rvb2xzTmFtZX19XHJcbmZ1bmMgKHMgKnN7e1VwcGVyVG9vbHNOYW1lfX0pIEVuYWJsZShjdHggY29udGV4dC5Db250ZXh0KSBlcnJvciB7XHJcblx0dmFyIChcclxuXHRcdHIgID0gZy5SZXF1ZXN0RnJvbUN0eChjdHgpXHJcblx0XHRpbiB7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1FbmFibGVSZXFcclxuXHQpXHJcblx0aWYgZXJyIDo9IHIuUGFyc2UoXHUwMDI2aW4pOyBlcnIgIT0gbmlsIHtcclxuXHRcdHJldHVybiBlcnJcclxuXHR9XHJcblx0Ly8gY2hlY2sgZGF0YSBpcyBleGlzdHNcclxuXHRfLCBlcnIgOj0gdXRpbGl0eS5DaGVja0RhdGFFeGlzdChjdHgsIHMudGFibGUoKSwgZy5NYXB7XHJcblx0XHRzLmdldFByaW1hcnlLZXlGaWVsZCgpOiBpbi5JZCxcclxuXHRcdHMuZ2V0RW5hYmxlRmllbGQoKTogICAgdXRpbGl0eS5FbmFibGUsXHJcblx0fSwgZnVuYyhjdHggY29udGV4dC5Db250ZXh0LCBjb3VudCBpbnQpIChib29sLCBlcnJvcikge1xyXG5cdFx0aWYgY291bnQgIT0gbGVuKGluLklkKSB7XHJcblx0XHRcdHJldHVybiBmYWxzZSwgZXJyb3JzLk5ldyh1dGlsaXR5LlRmKGN0eCwgYHsjQmFzaWNVc2VyRW5hYmxlZERhdGFOdW1Jc05vdEVxdWFsfWAsIGNvdW50LCBsZW4oaW4uSWQpKSlcclxuXHRcdH1cclxuXHRcdHJldHVybiB0cnVlLCBuaWxcclxuXHR9KVxyXG5cdGlmIGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIGVyclxyXG5cdH1cclxuXHRfLCBlcnIgPSBzLm1vZGVsKCkuV2hlcmVJbihzLmdldFByaW1hcnlLZXlGaWVsZCgpLCBpbi5JZCkuRGF0YShnLk1hcHtcclxuXHRcdHMuZ2V0RW5hYmxlRmllbGQoKTogdXRpbGl0eS5Jc0RlbCxcclxuXHR9KS5VcGRhdGUoKVxyXG5cdHJldHVybiBlcnJcclxufVxyXG5cclxuLy8gSW5mbyB7e1Rvb2xzTmFtZX196K+m5oOFXHJcbmZ1bmMgKHMgKnN7e1VwcGVyVG9vbHNOYW1lfX0pIEluZm8oY3R4IGNvbnRleHQuQ29udGV4dCkgKHJlcyAqZW50aXR5Lnt7VXBwZXJUb29sc05hbWV9fSwgZXJyIGVycm9yKSB7XHJcblx0dmFyIChcclxuXHRcdHIgID0gZy5SZXF1ZXN0RnJvbUN0eChjdHgpXHJcblx0XHRpbiB7e01vZHVsZU5hbWV9fS57e1VwcGVyVG9vbHNOYW1lfX1JbmZvUmVxXHJcblx0KVxyXG5cdGlmIGVyciA6PSByLlBhcnNlKFx1MDAyNmluKTsgZXJyICE9IG5pbCB7XHJcblx0XHRyZXR1cm4gbmlsLCBlcnJcclxuXHR9XHJcblx0aW5mbywgZXJyIDo9IHMuR2V0KGN0eCwgZy5NYXB7XHJcblx0XHRzLmdldFByaW1hcnlLZXlGaWVsZCgpOiBpbi5JZCxcclxuXHR9KVxyXG5cdGlmIGVyciAhPSBuaWwge1xyXG5cdFx0cmV0dXJuIG5pbCwgZXJyXHJcblx0fVxyXG5cdGlmIGluZm8uSWQgXHUwMDNjPSAwIHtcclxuXHRcdHJldHVybiBuaWwsIGVycm9ycy5OZXcodXRpbGl0eS5UZihjdHgsIGB7I0Jhc2ljVXNlckRhdGFJc05vdEV4aXN0fWAsIGluLklkKSlcclxuXHR9XHJcblx0cmV0dXJuIGluZm8sIG5pbFxyXG59XHJcbi8vIEdldCDmn6Xor6LkuIDmnaHmlbDmja5cclxuZnVuYyAocyAqc3t7VXBwZXJUb29sc05hbWV9fSkgR2V0KGN0eCBjb250ZXh0LkNvbnRleHQsIHdoZXJlIGludGVyZmFjZXt9KSAocmVzICplbnRpdHkue3tVcHBlclRvb2xzTmFtZX19LCBlcnIgZXJyb3IpIHtcclxuXHRyZXMgPSBcdTAwMjZlbnRpdHkue3tVcHBlclRvb2xzTmFtZX19e31cclxuXHRlcnIgPSBzLm1vZGVsKCkuQ3R4KGN0eCkuV2hlcmUod2hlcmUpLkxpbWl0KDEpLlNjYW4oXHUwMDI2cmVzKVxyXG5cdGlmIGVyciAhPSBuaWwgXHUwMDI2XHUwMDI2ICFlcnJvcnMuSXMoZXJyLCBzcWwuRXJyTm9Sb3dzKSB7XHJcblx0XHRyZXR1cm4gbmlsLCBlcnJcclxuXHR9XHJcblx0cmV0dXJuIHJlcywgbmlsXHJcbn1cclxuXHJcbi8vIEdldExpc3Qg6I635b6X5oyH5a6a5p2h5Lu255qE5aSa5p2h5pWw5o2uXHJcbmZ1bmMgKHMgKnN7e1VwcGVyVG9vbHNOYW1lfX0pIEdldExpc3QoY3R4IGNvbnRleHQuQ29udGV4dCwgd2hlcmUgaW50ZXJmYWNle30pIChyZXMgW10qZW50aXR5Lnt7VXBwZXJUb29sc05hbWV9fSwgZXJyIGVycm9yKSB7XHJcbiAgICAvLyB0b2RvIHdyaXRlIGxvZ2ljIGhlcmVcclxuXHRyZXMgPSBtYWtlKFtdKmVudGl0eS57e1VwcGVyVG9vbHNOYW1lfX0sIDApXHJcblx0ZXJyID0gcy5tb2RlbCgpLkN0eChjdHgpLldoZXJlKHdoZXJlKS5TY2FuKFx1MDAyNnJlcylcclxuXHRpZiBlcnIgIT0gbmlsIFx1MDAyNlx1MDAyNiAhZXJyb3JzLklzKGVyciwgc3FsLkVyck5vUm93cykge1xyXG5cdFx0cmV0dXJuIG5pbCwgZXJyXHJcblx0fVxyXG5cdHJldHVybiByZXMsIG5pbFxyXG59XHJcbiIsInV0aWxpdHkudGVtcGxhdGUiOiJwYWNrYWdlIHV0aWxpdHlcblxuaW1wb3J0IChcblx0XCJjb250ZXh0XCJcblx0XCJtYXRoXCJcblx0XCJyZWZsZWN0XCJcblx0XCJnaXRodWIuY29tL2dvZ2YvZ2YvdjIvZnJhbWUvZ1wiXG4pXG5cbi8vIFVuaXRQYWdlIOWIhumhtVxudHlwZSBVbml0UGFnZSBzdHJ1Y3Qge1xuXHRDdXJyZW50UGFnZSBpbnQgYGpzb246XCJjdXJyZW50UGFnZVwiIGRjOlwi6aG156CBXCIgZDpcIjFcImBcblx0UGVyUGFnZSAgICAgaW50IGBqc29uOlwicGVyUGFnZVwiIGRjOlwi5q+P6aG15pWw5o2u5p2h5pWwXCIgZDpcIjIwXCJgXG59XG5cbi8vIFVuaXRQYWdlUmVzIOWIhumhteW3peWFt+adoVxudHlwZSBVbml0UGFnZVJlcyBzdHJ1Y3Qge1xuXHRVbml0UGFnZVxuXHRUb3RhbCAgICBpbnQgICAgYGpzb246XCJ0b3RhbFwiIGRjOlwi5oC75p2h5pWwXCJgXG5cdExhc3RQYWdlIGludCAgICBganNvbjpcImxhc3RQYWdlXCIgZGM6XCLmgLvpobXmlbBcImBcblx0UGFnZUJhciAgc3RyaW5nIGBqc29uOlwicGFnZUJhclwiYFxufVxuXG4vLyBVbml0UmVzIOmAmueUqOeahGh0dHDlk43lupRcbnR5cGUgVW5pdFJlcyBzdHJ1Y3Qge1xuXHRDb2RlICAgICAgaW50ICAgICAgICAgYGpzb246XCJjb2RlXCIgZDpcIjIwMFwiYFxuXHREYXRhICAgICAgaW50ZXJmYWNle30gYGpzb246XCJkYXRhXCJgXG5cdE1lc3NhZ2UgICBzdHJpbmcgICAgICBganNvbjpcIm1lc3NhZ2VcImBcblx0UmVxdWVzdElkIHN0cmluZyAgICAgIGBqc29uOlwicmVxdWVzdElkXCJgXG59XG5cbi8vIFVuaXREYXRhUGFnZVJlcyDliIbpobXmlbDmja7pgJrnlKjov5Tlm55cbnR5cGUgVW5pdERhdGFQYWdlUmVzIHN0cnVjdCB7XG5cdFVuaXRQYWdlUmVzXG5cdERhdGEgaW50ZXJmYWNle30gYGpzb246XCJkYXRhXCJgXG59XG5cbi8vIEdldFBhZ2Ug6I635b6X5pWw5o2u5YiG6aG1XG4vLyBwYWdlIOmhteeggVxuLy8gbGltaXQg5q+P6aG15p2h5pWwXG4vLyB0b3RhbCDmlbDmja7mgLvmlbBcbmZ1bmMgR2V0UGFnZShjdHggY29udGV4dC5Db250ZXh0LCBwYWdlLCBzaXplLCB0b3RhbCBpbnQpICpVbml0UGFnZVJlcyB7XG5cdHAgOj0gXHUwMDI2VW5pdFBhZ2VSZXN7fVxuXHRpZiBwYWdlIFx1MDAzZT0gMSB7XG5cdFx0cC5DdXJyZW50UGFnZSA9IHBhZ2Vcblx0fSBlbHNlIHtcblx0XHRwLkN1cnJlbnRQYWdlID0gMVxuXHR9XG5cdGlmIHNpemUgXHUwMDNlPSAxIHtcblx0XHRwLlBlclBhZ2UgPSBzaXplXG5cdH0gZWxzZSB7XG5cdFx0cC5QZXJQYWdlID0gMjBcblx0fVxuXHRwLlRvdGFsID0gMVxuXHRwLkxhc3RQYWdlID0gMVxuXG5cdGlmIHRvdGFsIFx1MDAzZSBwLlBlclBhZ2Uge1xuXHRcdHAuTGFzdFBhZ2UgPSBpbnQobWF0aC5DZWlsKGZsb2F0NjQodG90YWwpIC8gZmxvYXQ2NChwLlBlclBhZ2UpKSlcblx0fVxuXHRwLlRvdGFsID0gdG90YWxcblx0cC5QYWdlQmFyID0gZy5SZXF1ZXN0RnJvbUN0eChjdHgpLkdldFBhZ2UocC5Ub3RhbCwgcC5QZXJQYWdlKS5HZXRDb250ZW50KDEpXG5cdHJldHVybiBwXG59XG4vLyBSZXNwb25zZVRvIOmAmueUqOaVsOaNrui/lOWbnlxuZnVuYyBSZXNwb25zZVRvKGNvZGUgaW50LCBkYXRhIGludGVyZmFjZXt9LCBtZXNzYWdlIC4uLnN0cmluZykgKlVuaXRSZXMge1xuXHRkZWZhdWx0TXNnIDo9IFwi5pON5L2c5oiQ5YqfXCJcblx0Ly8g5aSE55CG5pWw5o2u77yM5aaC5p6caW50ZXJmYWNle30gPT0gbmlsXG5cdGlmIGRhdGEgIT0gbmlsIHtcblx0XHRpZiByZWZsZWN0LlR5cGVPZihkYXRhKS5LaW5kKCkgPT0gcmVmbGVjdC5TbGljZSB7XG5cdFx0XHRzIDo9IHJlZmxlY3QuVmFsdWVPZihkYXRhKVxuXHRcdFx0aWYgcy5MZW4oKSA9PSAwIHtcblx0XHRcdFx0ZGF0YSA9IG1ha2UoW11zdHJpbmcsIDApXG5cdFx0XHR9XG5cdFx0fVxuXHR9IGVsc2Uge1xuXHRcdGRhdGEgPSBtYWtlKG1hcFtzdHJpbmddaW50ZXJmYWNle30pXG5cdH1cblx0aWYgbGVuKG1lc3NhZ2UpID09IDEge1xuXHRcdGRlZmF1bHRNc2cgPSBtZXNzYWdlWzBdXG5cdH1cblx0cmV0dXJuIFx1MDAyNlVuaXRSZXN7XG5cdFx0Q29kZTogICAgICBjb2RlLFxuXHRcdERhdGE6ICAgICAgZGF0YSxcblx0XHRNZXNzYWdlOiAgIGRlZmF1bHRNc2csXG5cdFx0UmVxdWVzdElkOiBcIlwiLCAvLyDkuLrnqbrlsLHooYzvvIzlm6DkuLrvvIzmoLnmja7or7fmsYLpk77ot6/vvIzmnIDnu4jkvJrmiZPliLByZXNwb25zZV9taWRkbGV3YXJl55qEUmVzcG9uc2XkuIpcblx0fVxufVxuXG4vLyBSZXNwb25zZVBhZ2VUbyDpgJrnlKjliIbpobXmlbDmja7lpITnkIZcbmZ1bmMgUmVzcG9uc2VQYWdlVG8oY29kZSBpbnQsIGRhdGEgaW50ZXJmYWNle30sIHBhZ2UgKlVuaXRQYWdlUmVzLCBtZXNzYWdlIC4uLnN0cmluZykgKlVuaXRSZXMge1xuXHRkZWZhdWx0TXNnIDo9IFwi5pON5L2c5oiQ5YqfXCJcblx0aWYgbGVuKG1lc3NhZ2UpID09IDEge1xuXHRcdGRlZmF1bHRNc2cgPSBtZXNzYWdlWzBdXG5cdH1cblx0ZGF0YU1hcCA6PSBcdTAwMjZVbml0RGF0YVBhZ2VSZXN7fVxuXHQvLyDlpITnkIbmlbDmja5cblx0ZGF0YU1hcC5EYXRhID0gZGF0YVxuXHRkYXRhTWFwLlBlclBhZ2UgPSBwYWdlLlBlclBhZ2Vcblx0ZGF0YU1hcC5DdXJyZW50UGFnZSA9IHBhZ2UuQ3VycmVudFBhZ2Vcblx0ZGF0YU1hcC5Ub3RhbCA9IHBhZ2UuVG90YWxcblx0ZGF0YU1hcC5MYXN0UGFnZSA9IHBhZ2UuTGFzdFBhZ2Vcblx0ZGF0YU1hcC5QYWdlQmFyID0gcGFnZS5QYWdlQmFyXG5cdC8vIOWmguaenOaYr2cuTWFw6ZyA6KaB6YGN5Y6GTWFw5Lit55qE5YC85Y6L5YWlZGF0YU1hcOS4re+8jOWwhuWFtktleei9rOaNouS4uuWwj+WGmVxuXHRpZiByZWZsZWN0LlR5cGVPZihkYXRhKS5LaW5kKCkgPT0gcmVmbGVjdC5TbGljZSB7XG5cdFx0cyA6PSByZWZsZWN0LlZhbHVlT2YoZGF0YSlcblx0XHRpZiBzLkxlbigpID09IDAge1xuXHRcdFx0ZGF0YU1hcC5EYXRhID0gbWFrZShbXXN0cmluZywgMClcblx0XHR9XG5cdH1cblx0cmV0dXJuIFx1MDAyNlVuaXRSZXN7XG5cdFx0Q29kZTogICAgICBjb2RlLFxuXHRcdERhdGE6ICAgICAgZGF0YU1hcCxcblx0XHRNZXNzYWdlOiAgIGRlZmF1bHRNc2csXG5cdFx0UmVxdWVzdElkOiBcIlwiLCAvLyDkuLrnqbrlsLHooYzvvIzlm6DkuLrvvIzmoLnmja7or7fmsYLpk77ot6/vvIzmnIDnu4jkvJrmiZPliLByZXNwb25zZV9taWRkbGV3YXJl55qEUmVzcG9uc2XkuIpcblx0fVxufSJ9`
	if baseDeContent, err := base64.StdEncoding.DecodeString(templateData); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal([]byte(baseDeContent), &templateContent); err != nil {
			panic(err)
		}
	}
}

// listPath 开发环境使用，用于遍历模板文件的内容处理后配合getTemplateData()使用
func listPath() {
	// 开发环境下遍历模板文件夹，读取模板内容
	filepath.Walk("./template", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			if con, err := os.ReadFile(path); err != nil {
				panic(err)
			} else {
				templateContent[filepath.Base(path)] = string(con)
			}
		}
		return nil
	})
	if con, err := json.Marshal(templateContent); err != nil {
		panic(err)
	} else {
		//fmt.Println(con)
		file, _ := os.OpenFile("./template/out.txt", os.O_CREATE|os.O_WRONLY, 0666)
		_, err := file.WriteString(base64.StdEncoding.EncodeToString(con))
		if err != nil {
			panic(err)
		}
	}
}
