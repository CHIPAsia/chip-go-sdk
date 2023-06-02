package chip

type RetainLevelDetail struct {
	Description     string  `json:"description"`
	CampaignID      string  `json:"campaign_id"`
	DiscountAmount  int     `json:"discount_amount"`
	DiscountPercent float32 `json:"discount_percent"`
}
