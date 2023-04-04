package models

type FraudData struct {
	CreditorFraudScore          int32  `json:"creditorFraudScore,omitempty"`
	EloBrandFraudScore          int32  `json:"eloBrandFraudScore,omitempty"`
	FraudScorePrimaryReason     int32  `json:"fraudScorePrimaryReason,omitempty"`
	FraudScoreSecondaryReason   int32  `json:"fraudScoreSecondaryReason,omitempty"`
	FraudScoreTertiaryReason    int32  `json:"fraudScoreTertiaryReason,omitempty"`
	FraudDecisionRecommendation string `json:"fraudDecisionRecommendation,omitempty"`
	ScoreOriginIndicator        int32  `json:"scoreOriginIndicator,omitempty"`
}
