package rsp

type Errno struct {
	Code int
	Message string
}

var (
	OK				= &Errno{0,"OK"}
	DatabaseErr		= &Errno{10201,"Database Error"}
)



