package handlers

import (
	"fmt"

	"github.com/Jitesh117/brainrot-lsp/mappers"
	_ "github.com/tliron/commonlog/simple"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TextDocumentCompletion(
	context *glsp.Context,
	params *protocol.CompletionParams,
) (interface{}, error) {
	var completionItems []protocol.CompletionItem
	for word, entry := range mappers.BrainrotMapper {
		term := entry.Term
		description := entry.Description
		detail := fmt.Sprintf("%s\n%s", term, description)
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word,
			Detail:     &detail,
			InsertText: &term,
		})
	}
	return completionItems, nil
}
