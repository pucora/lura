package graphql

import "testing"

func TestEmptyVariablePlaceholderDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("unexpected panic: %v", r)
		}
	}()
	New(Options{
		Type: OperationQuery,
		GraphQLRequest: GraphQLRequest{
			Query: "{x}",
			Variables: map[string]interface{}{
				"foo": "",
			},
		},
	})
}
