package constants

type Message string

const (
	OkMessage Message = "OK"

	BadRequestMessage              Message = "Bad Request"
	BadRequestAlreadyExistsMessage Message = "Already Exists"
	UnAuthorizedMessage            Message = "UnAuthorized"
	ForbiddenMessage               Message = "Forbidden"
	NotFoundMessage                Message = "Not Found"
	NotFoundSubscriptionMessage    Message = "Not Found Subscription"
	NotFoundUser                   Message = "Not Found User"
	NotFoundUserId                 Message = "Not Found UserId"
	NotFountLlm                    Message = "Not Found LLM"
	InternalServerErrorMessage     Message = "Internal Server Error"

	DatabaseError Message = "Database Error"
	StripeError   Message = "Stripe Error"
)
