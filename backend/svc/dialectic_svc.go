package svc

import (
	ai "epistemic-me-backend/ai"
	db "epistemic-me-backend/db"
	"epistemic-me-backend/svc/models"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
)

type DialecticService struct {
	kv   *db.KeyValueStore
	bsvc *BeliefService
	aih  *ai.AIHelper
}

// NewDialecticService initializes and returns a new DialecticService.
func NewDialecticService(kv *db.KeyValueStore, bsvc *BeliefService, aih *ai.AIHelper) *DialecticService {
	return &DialecticService{
		kv:   kv,
		bsvc: bsvc,
		aih:  aih,
	}
}

func (dsvc *DialecticService) CreateDialectic(input *models.CreateDialecticInput) (*models.CreateDialecticOutput, error) {
	newDialecticId := "di_" + uuid.New().String()

	// Create a new Dialectic struct
	dialectic := models.Dialectic{
		ID:     newDialecticId,
		UserID: input.UserID,
		Agent: models.Agent{
			AgentType:     models.AgentTypeGPTLatest,
			DialecticType: input.DialecticType,
		},
		UserInteractions: []models.DialecticalInteraction{},
	}

	// Generate the first interaction
	newInteraction, err := dsvc.generatePendingDialecticalInteraction(input.UserID, dialectic.UserInteractions)
	if err != nil {
		return nil, err
	}

	// Add the first interaction to the dialectic
	dialectic.UserInteractions = append(dialectic.UserInteractions, *newInteraction)

	// Store the dialectic
	err = dsvc.kv.Store(input.UserID, dialectic.ID, dialectic)
	if err != nil {
		return nil, err
	}

	return &models.CreateDialecticOutput{
		DialecticID: dialectic.ID,
		Dialectic:   dialectic,
	}, nil
}

func (dsvc *DialecticService) ListDialectics(input *models.ListDialecticsInput) (*models.ListDialecticsOutput, error) {
	// Retrieve all dialectics for the user
	dialectics, err := dsvc.kv.ListByType(input.UserID, reflect.TypeOf(models.Dialectic{}))
	if err != nil {
		return nil, err
	}

	// Convert the retrieved dialectics to the required type
	var dialecticModels []models.Dialectic
	for _, dialectic := range dialectics {
		storedDialectic := dialectic.(*models.Dialectic)
		dialecticModels = append(dialecticModels, *storedDialectic)
	}

	return &models.ListDialecticsOutput{
		Dialectics: dialecticModels,
	}, nil
}

func (dsvc *DialecticService) UpdateDialectic(input *models.UpdateDialecticInput) (*models.UpdateDialecticOutput, error) {
	log.Printf("UpdateDialectic called with input: %+v", input)

	kvResponse, err := dsvc.kv.Retrieve(input.UserID, input.DialecticID)
	if err != nil {
		log.Printf("Error retrieving dialectic: %v", err)
		return nil, err
	}

	// Type assertion to convert kvResponse to *Dialectic
	dialectic, ok := kvResponse.(*models.Dialectic)
	if !ok {
		return nil, fmt.Errorf("failed to assert kvResponse to *Dialectic")
	}

	interaction, err := getPendingInteraction(*dialectic)
	if err != nil {
		log.Printf("Error getting pending interaction: %v", err)
		return nil, err
	}
	log.Printf("Pending interaction before update: %+v", interaction)

	// Update pending interaction with answer and mark it as answered
	interaction.UserAnswer = input.Answer
	interaction.Status = models.StatusAnswered
	log.Printf("Interaction after update: %+v", interaction)

	// extract beliefs from the completed interaction
	log.Printf("Calling getDialecticalInteractionAsEvent with interaction: %+v", interaction)
	interactionEvent, err := getDialecticalInteractionAsEvent(*interaction)
	if err != nil {
		log.Printf("Error in getDialecticalInteractionAsEvent: %v", err)
		return nil, err
	}
	log.Printf("Interaction event created: %+v", interactionEvent)

	// given the interaction event update the users existing belief system
	// by updating old beleifs or creating new ones
	err = dsvc.updateBeliefSystemForInteraction(*interactionEvent, input.UserID)
	if err != nil {
		return nil, err
	}

	// generate a new interaction given updated state of dialectic and user belief system
	newInteraction, err := dsvc.generatePendingDialecticalInteraction(input.UserID, dialectic.UserInteractions)
	if err != nil {
		return nil, err
	}

	dialectic.UserInteractions = append(dialectic.UserInteractions, *newInteraction)

	err = dsvc.kv.Store(input.UserID, dialectic.ID, *dialectic)
	if err != nil {
		return nil, err
	}

	return &models.UpdateDialecticOutput{
		Dialectic: *dialectic,
	}, nil

}

func (dsvc *DialecticService) updateBeliefSystemForInteraction(interactionEvent ai.InteractionEvent, userID string) error {
	listBeliefsOutput, err := dsvc.bsvc.ListBeliefs(&models.ListBeliefsInput{
		UserID: userID,
	})

	if err != nil {
		return err
	}

	updatedBeleifs := 0

	// for each of the user beliefs, check to see if event has relevance and update accordingly
	for _, existingBelief := range listBeliefsOutput.Beliefs {

		shouldUpdate, interpretedBeliefStr, err := dsvc.aih.UpdateBeliefWithInteractionEvent(interactionEvent, existingBelief.GetContentAsString())
		if err != nil {
			return err
		}

		if shouldUpdate {
			updatedBeleifs += 1
			// store the interpeted belief as a user belief so it will be included in the belief system
			_, err = dsvc.bsvc.UpdateBelief(&models.UpdateBeliefInput{
				UserID:               userID,
				BeliefID:             existingBelief.ID,
				CurrentVersion:       existingBelief.Version,
				UpdatedBeliefContent: interpretedBeliefStr,
			})

			if err != nil {
				return err
			}
		}
	}

	// if we've updated no existing beleifs, create a new one
	// todo: @deen this may need to become more sophisticated in the future
	if updatedBeleifs == 0 {
		interpretedBeliefStr, err := dsvc.aih.GetInteractionEventAsBelief(interactionEvent)
		if err != nil {
			return err
		}

		// store the interpeted belief as a user belief so it will be included in the belief system
		_, err = dsvc.bsvc.CreateBelief(&models.CreateBeliefInput{
			UserID:        userID,
			BeliefContent: interpretedBeliefStr,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (dsvc *DialecticService) generatePendingDialecticalInteraction(userID string, previousInteractions []models.DialecticalInteraction) (*models.DialecticalInteraction, error) {
	// Get the Latest Belief System in orer to update the dialectic
	listBeliefsOutput, err := dsvc.bsvc.ListBeliefs(&models.ListBeliefsInput{
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	user_belief_system := listBeliefsOutput.BeliefSystem

	var events []ai.InteractionEvent
	for _, interaction := range previousInteractions {
		if interaction.Status == models.StatusAnswered {
			interactionEvent, err := getDialecticalInteractionAsEvent(interaction)
			if err != nil {
				return nil, err
			}
			events = append(events, *interactionEvent)
		}
	}

	question, err := dsvc.aih.GenerateQuestion(user_belief_system.RawStr, events)

	if err != nil {
		return nil, err
	}

	return &models.DialecticalInteraction{
		Question: models.Question{
			Question: question,
		},
		Status: models.StatusPendingAnswer,
	}, nil
}

func getPendingInteraction(dialectic models.Dialectic) (*models.DialecticalInteraction, error) {
	// Get the latest interaction
	if len(dialectic.UserInteractions) == 0 {
		return nil, fmt.Errorf("no interactions found in the dialectic")
	}
	latestInteraction := dialectic.UserInteractions[len(dialectic.UserInteractions)-1]
	// Check if the latest interaction is pending
	if latestInteraction.Status != models.StatusPendingAnswer {
		return nil, fmt.Errorf("latest interaction is not pending")
	}
	return &latestInteraction, nil
}

func getDialecticalInteractionAsEvent(interaction models.DialecticalInteraction) (*ai.InteractionEvent, error) {
	log.Printf("getDialecticalInteractionAsEvent called with interaction status: %v", interaction.Status)
	if interaction.Status != models.StatusAnswered {
		return nil, fmt.Errorf("interaction is not answered yet")
	}
	return &ai.InteractionEvent{
		Question: interaction.Question.Question,
		Answer:   interaction.UserAnswer.UserAnswer,
	}, nil
}
