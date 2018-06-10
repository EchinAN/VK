package main

type PresponseJson struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int    `json:"id"`
			FromID      int    `json:"from_id"`
			OwnerID     int    `json:"owner_id"`
			Date        int    `json:"date"`
			MarkedASADS int    `json:"marked_as_ads"`
			PostType    string `json:"post_type"`
			Text        string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}
type CresponseJson struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID     int    `json:"id"`
			FromID int    `json:"from_id"`
			Date   int    `json:"date"`
			Text   string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}
type PostStructJson struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int    `json:"id"`
			FromID      int    `json:"from_id"`
			OwnerID     int    `json:"owner_id"`
			Date        int    `json:"date"`
			MarkedASADS int    `json:"marked_as_ads"`
			PostType    string `json:"post_type"`
			Text        string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}
type CommentsStructJson struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID     int    `json:"id"`
			FromID int    `json:"from_id"`
			Date   int    `json:"date"`
			Text   string `json:"text"`
		} `json:"items"`
	} `json:"response"`
}