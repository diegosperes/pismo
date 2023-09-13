package model

type OperationType int

const (
	OperationTypeLumpSum      OperationType = 1
	OperationTypeInstallments OperationType = 2
	OperationTypeWithdraw     OperationType = 3
	OperationTypePayment      OperationType = 4
)

func (i OperationType) GetName() string {
	switch i {
	case OperationTypeLumpSum:
		return "LumpSum"
	case OperationTypeInstallments:
		return "Installments"
	case OperationTypeWithdraw:
		return "Withdraw"
	case OperationTypePayment:
		return "Payment"
	default:
		return "UNKNOW"
	}
}
