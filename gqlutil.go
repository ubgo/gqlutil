package gqlutil

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// get the fields requested in a graphql query
// Useful for constructing SQL queries from the GraphQL query
func GetFieldsRequested(ctx context.Context) []string {
	reqCtx := graphql.GetOperationContext(ctx)
	fieldSelections := graphql.GetFieldContext(ctx).Field.Selections
	return RecurseSelectionSets(reqCtx, []string{}, fieldSelections)
}

// this get the same as above, but for use within the field resolver. So in the field resolver
// if we didn't get a field we can grab what was requested above and load it.
// Useful for lazy resolving data
// NB this will panic if there is no parent.
func GetParentFieldsRequested(ctx context.Context) []string {
	reqCtx := graphql.GetOperationContext(ctx)
	fieldSelections := graphql.GetFieldContext(ctx).Parent.Field.Selections
	return RecurseSelectionSets(reqCtx, []string{}, fieldSelections)
}

// loop through selection sets flattening out the field names from fragments,
// ignoring private fields
func RecurseSelectionSets(reqCtx *graphql.OperationContext, fields []string, selection ast.SelectionSet) []string {
	for _, sel := range selection {
		switch sel := sel.(type) {
		case *ast.Field:
			// ignore private field names
			if !strings.HasPrefix(sel.Name, "__") {
				fields = append(fields, sel.Name)
			}
		case *ast.InlineFragment:
			fields = RecurseSelectionSets(reqCtx, fields, sel.SelectionSet)
		case *ast.FragmentSpread:
			fragment := reqCtx.Doc.Fragments.ForName(sel.Name)
			fields = RecurseSelectionSets(reqCtx, fields, fragment.SelectionSet)
		}
	}
	return fields
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		// fmt.Println(prefixColumn)
		preloads = append(preloads, prefixColumn)

		// preloads = append(preloads, GetPreloadString(prefix, column.ObjectDefinition.Name	))

		// We need this for fragments identification ... on UserLink {}
		preloads = append(preloads, column.ObjectDefinition.Name+"."+column.Name)

		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + name
	}
	return name
}
