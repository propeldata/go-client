package models

type Operator string

const (
	OperatorEquals               Operator = "EQUALS"
	OperatorNotEquals            Operator = "NOT_EQUALS"
	OperatorGreaterThan          Operator = "GREATER_THAN"
	OperatorGreaterThanOrEqualTo Operator = "GREATER_THAN_OR_EQUAL_TO"
	OperatorLessThan             Operator = "LESS_THAN"
	OperatorLessThanOrEqualTo    Operator = "LESS_THAN_OR_EQUAL_TO"
	OperatorIsNull               Operator = "IS_NULL"
	OperatorIsNotNull            Operator = "IS_NOT_NULL"
	OperatorLike                 Operator = "LIKE"
	OperatorNotLike              Operator = "NOT_LIKE"
)

type FilterInput struct {
	Column   string        `json:"column"`
	Operator Operator      `json:"operator"`
	Value    *string       `json:"value,omitempty"`
	And      []FilterInput `json:"and,omitempty"`
	Or       []FilterInput `json:"or,omitempty"`
}
