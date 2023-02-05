package commands

import (
	"example-bot/api"
	googleSearch "github.com/rocketlaunchr/google-search"
)

type Searching struct {
	MessageID    string
	SearchResult []googleSearch.Result
}

var SearchingList = []Searching{}

const SearchingStamp = "7bc14aa3-d930-4ce3-9c2b-ebb557d4e6b2"

func Search(args ArgsV2) {
	//if len(args.Slice) <= 1 {
	//	return
	//}
	message, err := api.PostMessageWithErr(args.ChannelID, ":loading: Searching...")
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on PostMessageWithErr")
		return
	}
	searchResult, err := api.SearchByText(args.MessageText)
	if err != nil {
		api.EditMessage(message.Id, "Error happen on SearchByText")
		return
	}
	if len(searchResult) == 0 {
		api.EditMessage(message.Id, "No result")
		return
	}
	err = api.EditMessageWithErr(message.Id, searchResult[0].URL)
	if err != nil {
		api.PostMessage(args.ChannelID, "Error happen on EditMessageWithErr")
		return
	}
	//err = api.AddStampByStampID(message.Id, SearchingStamp)
	//if args.Slice[2] == "all" || args.Slice[2] == "-a" {
	//	for _, result := range searchResult {
	//		api.AddMessageWithNewLine(message.Id, result.URL)
	//	}
	//	return
	//}
	//if err != nil {
	//	api.PostMessage(args.ChannelID, "Error happen on AddStampByStampID")
	//	return
	//}
	//SearchingList = append(SearchingList, Searching{MessageID: message.Id, SearchResult: searchResult})
}
