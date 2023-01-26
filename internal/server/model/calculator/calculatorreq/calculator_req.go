package calculatorreq

type IpCalcParams struct {
	IpAddr string `form:"ip_addr" binding:"required"`
}
