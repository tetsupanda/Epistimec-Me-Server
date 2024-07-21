package models

import pbmodels "epistemic-me-backend/pb/models"

// Question represents a request for information from a user.
type Question struct {
	Question           string `json:"question"`
	CreatedAtMillisUTC int64  `json:"created_at_millis_utc"`
}

func (q Question) ToProto() *pbmodels.Question {
	return &pbmodels.Question{
		Question:           q.Question,
		CreatedAtMillisUtc: q.CreatedAtMillisUTC,
	}
}

// UserAnswer represents a user's answer to a question.
type UserAnswer struct {
	UserAnswer         string `json:"user_answer"`
	CreatedAtMillisUTC int64  `json:"created_at_millis_utc"`
}

func (ua UserAnswer) ToProto() *pbmodels.UserAnswer {
	return &pbmodels.UserAnswer{
		UserAnswer:         ua.UserAnswer,
		CreatedAtMillisUtc: ua.CreatedAtMillisUTC,
	}
}

// DialecticalInteractionStatus represents the status of the interaction.
type DialecticalInteractionStatus int32

const (
	StatusInvalid       DialecticalInteractionStatus = 0
	StatusPendingAnswer DialecticalInteractionStatus = 1
	StatusAnswered      DialecticalInteractionStatus = 2
)

// DialecticType represents the type of dialectic strategy.
type DialecticType int32

const (
	DialecticTypeInvalid DialecticType = 0
	Default              DialecticType = 1
	Hegelian             DialecticType = 2
)

// AgentType represents the type of agent.
type AgentType int32

const (
	AgentTypeInvalid   AgentType = 0
	AgentTypeGPTLatest AgentType = 1
)

// Agent represents the system or user interacting with the user.
type Agent struct {
	AgentType     AgentType     `json:"agent_type"`
	DialecticType DialecticType `json:"dialectic_type"`
}

func (a Agent) ToProto() *pbmodels.Agent {
	var protoAgentType pbmodels.Agent_AgentType
	switch a.AgentType {
	case AgentTypeGPTLatest:
		protoAgentType = pbmodels.Agent_AGENT_TYPE_GPT_LATEST
	default:
		protoAgentType = pbmodels.Agent_AGENT_TYPE_INVALID
	}

	var protoDialecticType pbmodels.DialecticType
	switch a.DialecticType {
	case Default:
		protoDialecticType = pbmodels.DialecticType_DEFAULT
	case Hegelian:
		protoDialecticType = pbmodels.DialecticType_HEGELIAN
	default:
		protoDialecticType = pbmodels.DialecticType_DIALECTIC_TYPE_INVALID
	}

	return &pbmodels.Agent{
		AgentType:     protoAgentType,
		DialecticType: protoDialecticType,
	}
}

// DialecticalInteraction represents a question and answer interaction.
type DialecticalInteraction struct {
	Status             DialecticalInteractionStatus `json:"status"`
	Question           Question                     `json:"question"`
	UserAnswer         UserAnswer                   `json:"user_answer"`
	Beliefs            []Belief                     `json:"beliefs"`
	UpdatedAtMillisUTC int64                        `json:"updated_at_millis_utc"`
}

func (di DialecticalInteraction) ToProto() *pbmodels.DialecticalInteraction {
	protoBeliefs := make([]*pbmodels.Belief, len(di.Beliefs))
	for i, belief := range di.Beliefs {
		protoBeliefs[i] = belief.ToProto()
	}

	return &pbmodels.DialecticalInteraction{
		Question:           di.Question.ToProto(),
		UserAnswer:         di.UserAnswer.ToProto(),
		Beliefs:            protoBeliefs,
		UpdatedAtMillisUtc: di.UpdatedAtMillisUTC,
	}
}

// Dialectic represents a session to determine and clarify a user's beliefs.
type Dialectic struct {
	ID               string                   `json:"id"`
	UserID           string                   `json:"user_id"`
	Agent            Agent                    `json:"agent"`
	UserInteractions []DialecticalInteraction `json:"user_interactions"`
}

func (d Dialectic) ToProto() *pbmodels.Dialectic {
	protoInteractions := make([]*pbmodels.DialecticalInteraction, len(d.UserInteractions))
	for i, interaction := range d.UserInteractions {
		protoInteractions[i] = interaction.ToProto()
	}

	return &pbmodels.Dialectic{
		Id:               d.ID,
		UserId:           d.UserID,
		Agent:            d.Agent.ToProto(),
		UserInteractions: protoInteractions,
	}
}
