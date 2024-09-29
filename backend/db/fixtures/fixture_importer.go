package fixture_models

import (
	"fmt"
	"os"

	"epistemic-me-backend/db"
	"epistemic-me-backend/svc/models"

	"gopkg.in/yaml.v3"
)

// ImportFixtures imports the belief system fixtures into the given KeyValueStore
func ImportFixtures(kvStore *db.KeyValueStore) error {
	// Read the YAML file
	yamlFile, err := os.ReadFile("db/fixtures/belief_system_fixture.yaml")
	if err != nil {
		return fmt.Errorf("error reading YAML file: %v", err)
	}

	var fixture BeliefSystemFixture
	err = yaml.Unmarshal(yamlFile, &fixture)
	if err != nil {
		return fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	fixtureUserID := "fixture-user-id"

	beliefSystem := models.BeliefSystem{
		Beliefs:             []*models.Belief{},
		ObservationContexts: []*models.ObservationContext{},
	}

	for _, example := range fixture.BeliefSystem.Examples {
		// Create observation contexts
		contextMap := make(map[string]*models.ObservationContext)
		for _, oc := range example.ObservationContext {
			context := &models.ObservationContext{
				ID:             fmt.Sprintf("%s_%s", example.Name, oc.ContextName),
				Name:           oc.ContextName,
				ParentID:       oc.NestedWithin,
				PossibleValues: []string{},
			}
			contextMap[oc.ContextName] = context
			beliefSystem.ObservationContexts = append(beliefSystem.ObservationContexts, context)
		}

		// Create beliefs
		for _, b := range example.Beliefs {
			belief := &models.Belief{
				ID:                    fmt.Sprintf("%s_%s", example.Name, b.BeliefName),
				UserID:                fixtureUserID,
				Content:               []models.Content{{RawStr: b.Description}},
				ObservationContextIDs: []string{contextMap[b.Context].ID},
				Probabilities:         map[string]float32{b.PredictedOutcome: 0.8, b.CounterfactualOutcome: 0.2},
				Action:                "",
				Result:                b.PredictedOutcome,
			}
			beliefSystem.Beliefs = append(beliefSystem.Beliefs, belief)
		}
	}

	// Store the belief system
	err = kvStore.Store(fixtureUserID, "BeliefSystemId", beliefSystem, 1)
	if err != nil {
		return fmt.Errorf("error storing fixture belief system: %v", err)
	}

	fmt.Printf("Fixture belief system has been successfully imported with %d beliefs and %d observation contexts\n",
		len(beliefSystem.Beliefs), len(beliefSystem.ObservationContexts))

	return nil
}
