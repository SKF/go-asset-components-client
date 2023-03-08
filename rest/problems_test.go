package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/SKF/go-rest-utility/problems"
	"github.com/stretchr/testify/require"
)

func TestProblemDecoder_DecodeProblem(t *testing.T) {
	tests := []struct {
		Name     string
		body     interface{}
		wantType string
	}{
		{
			Name:     "internal server error",
			body:     problems.Internal(fmt.Errorf("some error")),
			wantType: "/problems/internal-server-error",
		},
		{
			Name: "validation error",
			body: problems.Validation(problems.ValidationReason{
				Name:   "nope",
				Reason: "don't wanna",
				Cause:  nil,
			}),
			wantType: "/problems/invalid-request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			cd := componentsProblemDecoder{}

			bs, err := json.Marshal(tt.body)
			require.NoError(t, err)

			resp := http.Response{
				Body: io.NopCloser(bytes.NewReader(bs)),
			}

			ctx := context.Background()

			res, err := cd.DecodeProblem(ctx, &resp)
			require.NoError(t, err)

			require.Equal(t, tt.wantType, res.ProblemType())
		})
	}
}

func TestValidation(t *testing.T) {
	cd := componentsProblemDecoder{}

	resBody := problems.Validation(problems.ValidationReason{
		Name:   "req",
		Reason: "to test",
		Cause:  nil,
	})

	bs, err := json.Marshal(resBody)
	require.NoError(t, err)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(bs)),
	}

	ctx := context.Background()

	res, err := cd.DecodeProblem(ctx, &resp)
	require.NoError(t, err)

	vp, ok := res.(problems.ValidationProblem)
	require.True(t, ok)
	require.Equal(t, "to test", vp.Reasons[0].Reason)
}
