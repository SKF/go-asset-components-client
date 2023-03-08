package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SKF/go-rest-utility/problems"
)

type componentsProblemDecoder struct{}

func (d *componentsProblemDecoder) DecodeProblem(_ context.Context, resp *http.Response) (problems.Problem, error) {
	defer resp.Body.Close()

	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	basicProblem := problems.BasicProblem{}
	if err := json.Unmarshal(bb, &basicProblem); err != nil {
		return nil, fmt.Errorf("BasicProblem json decoder: %w", err)
	}

	switch basicProblem.Type {
	case "/problems/invalid-request":
		validationProblem := problems.ValidationProblem{}
		if err := json.Unmarshal(bb, &validationProblem); err != nil {
			return nil, fmt.Errorf("BasicProblem json decoder: %w", err)
		}

		return validationProblem, nil
	default:
		return basicProblem, nil
	}
}
