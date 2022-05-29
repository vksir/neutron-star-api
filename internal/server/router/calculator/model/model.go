package model

type CalculatorErr struct {
	Detail string `json:"detail"`
}

type IPv4 struct {
	IP           string `json:"ip"`
	CIDR         string `json:"cidr"`
	Start        string `json:"start"`
	End          string `json:"end"`
	MaskBits     int    `json:"mask_bits"`
	SubnetMask   string `json:"subnet_mask"`
	Count        int64  `json:"count"`
	WildcardMask string `json:"wildcard_mask"`
}

type Ipv4CalcParams struct {
	IpAddr string `form:"ip_addr" binding:"required"`
}

type IPv6 struct {
	Short SubIPv6 `json:"short"`
	Long  SubIPv6 `json:"long"`
}

type SubIPv6 struct {
	IP       string `json:"ip"`
	CIDR     string `json:"cidr"`
	Start    string `json:"start"`
	End      string `json:"end"`
	MaskBits int    `json:"mask_bits"`
	Count    int64  `json:"count"`
}

type Ipv6CalcParams struct {
	IpAddr string `form:"ip_addr" binding:"required"`
}
