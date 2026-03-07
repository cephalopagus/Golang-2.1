package payments

type PaymentModule struct {
	paymentinfo map[int]PaymentInfo

	paymentMethod PaymentMethod
}

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
		paymentinfo:   make(map[int]PaymentInfo),
	}
}

/*
Принимает - описание проводимой оплаты, сумма оплаты
Возвращает - id операции
*/
func (p *PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}
	p.paymentinfo[id] = info

	return id
}

/*
Принимает - id операции
Возвращает - ничего
*/
func (p *PaymentModule) Cancel(id int) {

	info, ok := p.paymentinfo[id]
	if !ok {
		return
	}
	p.paymentMethod.Cancel(id)
	info.Cancelled = true
	p.paymentinfo[id] = info
}

/*
Принимает - id операции
Возвращает - информация о проведенной операции
*/
func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentinfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

/*
Принимает - ничего
Возвращает - информация о всех проведенных операциях
*/
func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentinfo))
	for k, v := range p.paymentinfo {
		tempMap[k] = v
	}
	return tempMap

}
