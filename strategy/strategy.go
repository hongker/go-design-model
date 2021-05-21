package strategy

// CashSuper
type CashSuper interface {
	AcceptMoney(money float64) float64
}
// CashNormal 普通
type CashNormal struct {

}

func (cash *CashNormal) AcceptMoney(money float64) float64  {
	return money
}

func NewCashNormal() *CashNormal {
	return &CashNormal{}
}
// CashRebate 折扣
type CashRebate struct {
	rebate float64
}

func NewCashRebate(rebate float64) *CashRebate {
	return &CashRebate{rebate: rebate}
}

func (cash CashRebate) AcceptMoney(money float64) float64 {
	return money * cash.rebate
}

// CashReturn 返现
type CashReturn struct {
	standardAmount float64
	returnAmount float64
}

func (cash CashReturn) AcceptMoney(money float64) float64 {
	if money >= cash.standardAmount {
		moneyMinus := int(money / cash.standardAmount)
		return money - float64(moneyMinus)*cash.returnAmount
	}

	return money
}
func NewCashReturn(standardAmount, returnAmount float64) *CashReturn {
	return &CashReturn{standardAmount: standardAmount,returnAmount: returnAmount}
}

type CashContext struct {
	strategy CashSuper
}

func (context CashContext) GetMoney(money float64) float64{
	return context.strategy.AcceptMoney(money)
}
type CashType int
const(
	HalfOff CashType = iota+1 // 打五折
	HundredDiscountTwenty // 满100减20
)

// NewCashContext 根据不同的策略设定收款方式
func NewCashContext(cashType CashType) CashContext {
	context := CashContext{}
	switch cashType {
	case HalfOff:
		context.strategy = NewCashRebate(0.5)
	case HundredDiscountTwenty:
		context.strategy = NewCashReturn(100, 20)
	default:
		context.strategy = NewCashNormal()
	}

	return context
}




