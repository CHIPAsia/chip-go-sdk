package chip

type RewardType string

const (
	GIFT     RewardType = "gift"
	DISCOUNT RewardType = "discount"
	MONEY    RewardType = "money"
)

type ReferralCodeDetail struct {
	Description     string     `json:"description"`
	CampaignID      string     `json:"campaign_id"`
	RewardType      RewardType `json:"reward_type"`
	DiscountAmount  int        `json:"discount_amount"`
	DiscountPercent float32    `json:"discount_percent"`
}
