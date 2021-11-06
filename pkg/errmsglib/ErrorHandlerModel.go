package errormsg

type ErrorHandlerModel struct {
	Error            error
	ErrorApplication string
	ErrorModule      string
	ErrorFile        string
	ErrorFunction    string
}
