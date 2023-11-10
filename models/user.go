package models

type (
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	TransferRequest struct {
		ReferenceNumber string  `json:"referenceNumber"`
		CreditAccount   string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		AdminFee        float64 `json:"adminFee"`
	}
	TransferData struct {
		ReferenceNumber string  `json:"referenceNumber"`
		CreditAccount   string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		Description     string  `json:"description"`
		Amount          float64 `json:"amount"`
		AdminFee        float64 `json:"adminFee"`
	}
	InquiryTransferCredit struct {
		ResponseCode    string  `json:"responsecode"`
		Description     string  `json:"description"`
		CreditAccoun    string  `json:"creditAccount"`
		DebitAccount    string  `json:"debitAccount"`
		ReferenceNumber string  `json:"referenceNumber"`
		CreditName      string  `json:"creditname"`
		Amount          float64 `json:"amount"`
	}
)
