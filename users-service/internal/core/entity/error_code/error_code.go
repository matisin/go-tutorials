package error_code

type ErrorCode string

// error code
const (
	Success               ErrorCode = "SUCCESS"
	InvalidRequest        ErrorCode = "INVALID_REQUEST"
	DuplicateUser         ErrorCode = "DUPLICATE_USER"
	InternalError         ErrorCode = "INTERNAL_ERROR"
	MailAlreadyExists     ErrorCode = "MAIL_ALREADY_EXISTS"
	WrongFormatMail       ErrorCode = "WRONG_FORMAT_MAIL"
	MissingSessionID      ErrorCode = "MISSING_SESSIONS_ID"
	WrongValidationCode   ErrorCode = "WRONG_VALIDATION_CODE"
	ExpiredValidationCode ErrorCode = "EXPIRED_VALIDATION_CODE"
	SignInFailed          ErrorCode = "SING_IN_FAILED"
	RefreshFailed         ErrorCode = "SING_IN_FAILED"
)

// error message
const (
	SuccessErrMsg        = "success"
	InternalErrMsg       = "internal error"
	InvalidRequestErrMsg = "invalid request"
)
