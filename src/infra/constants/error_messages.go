package constants

const (
	// global errors
	ErrCommonError     = "httpStatus:400;message:Bad request"
	ErrSortValidation  = "httpStatus:400;message:Value of sort param must be 'asc' or 'desc'"
	ErrBoolValidation  = "httpStatus:400;message:Failed boolean type"
	ErrParsingFormData = "httpStatus:400;message:Failed to parsing form data"
	ErrTypeOfData      = "Request type data number not correct, please check again your request"
	ErrTypeOfDataBool  = "Request type data boolean not correct, please check again your request"
	ErrEmptyRequest    = "Can not process empty request"
	ErrNotFoundRequest = "No response founded"

	// otp errors
	OtpTimeWaitingErr = "httpStatus:400;message:Silahkan tunggu 30 detik untuk mengirim ulang otp kembali"
)
