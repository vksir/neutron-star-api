package model

type CalculatorIpv4CalcParams struct {
	IpAddr string `form:"ip_addr" binding:"required"`
}

type CalculatorIpv4 struct {
}
