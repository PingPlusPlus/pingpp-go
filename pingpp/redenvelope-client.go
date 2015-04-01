package pingpp

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type RedEnvelopeClient struct {
	key     string
	backend Backend
}

func GetRedEnvelopeClient(key string) *RedEnvelopeClient {
	return &RedEnvelopeClient{
		key:     key,
		backend: getBackend(),
	}
}

func (redEnvelopeClient *RedEnvelopeClient) SetKey(key string) {
	redEnvelopeClient.key = key
}

func (redEnvelopeClient *RedEnvelopeClient) New(params *RedEnvelopeParams) (*RedEnvelope, error) {
	var redEnvelope RedEnvelope
	body := &url.Values{
		"amount": {strconv.FormatUint(params.Amount, 10)},
	}
	body.Add("order_no", params.Order_no)
	appstring, _ := json.Marshal(params.App)
	body.Add("app", string(appstring))
	fmt.Printf("App: %s\n", string(appstring))
	body.Add("body", params.Body)
	body.Add("channel", params.Channel)
	body.Add("currency", params.Currency)
	body.Add("subject", params.Subject)
	body.Add("recipient", params.Recipient)
	body.Add("description", params.Description)
	jsonstring, _ := json.Marshal(params)
	resp_byte, err := redEnvelopeClient.backend.CallJson("POST", "red_envelopes", redEnvelopeClient.key, jsonstring, &redEnvelope)
	json.Unmarshal(resp_byte, &redEnvelope)
	// fmt.Println(json.Unmarshal(resp_byte, &redEnvelope))
	return &redEnvelope, err
}

func (redEnvelopeClient *RedEnvelopeClient) Get(red_envelope_id string) (*RedEnvelope, error) {
	var redEnvelope RedEnvelope
	body := &url.Values{}
	resp_byte, err := redEnvelopeClient.backend.Call("GET", "red_envelopes/"+red_envelope_id, redEnvelopeClient.key, body, &redEnvelope)
	json.Unmarshal(resp_byte, &redEnvelope)
	return &redEnvelope, err

}

func (redEnvelopeClient *RedEnvelopeClient) List(params *RedEnvelopeListParams) (*RedEnvelopeList, error) {
	var redEnvelopes RedEnvelopeList
	body := &url.Values{}
	if params.Limit > 100 || params.Limit <= 0 {
		body.Add("limit", strconv.FormatUint(10, 10))
	} else {
		body.Add("limit", strconv.FormatUint(params.Limit, 10))
	}
	if params.Start_after != "" {
		body.Add("starting_after", params.Start_after)
	}
	if params.End_before != "" {
		body.Add("ending_before", params.End_before)
	}
	if params.Createdgt != "" {
		body.Add("created[gt]", params.Createdgt)
	}
	if params.Createdgte != "" {
		body.Add("created[gte]", params.Createdgte)
	}
	if params.Createdlt != "" {
		body.Add("created[lt]", params.Createdlt)
	}
	if params.Createdlte != "" {
		body.Add("created[lte]", params.Createdlte)
	}

	// body.Add("metadata", charge.metadata)
	resp_byte, err := redEnvelopeClient.backend.Call("GET", "red_envelopes", redEnvelopeClient.key, body, &redEnvelopes)
	json.Unmarshal(resp_byte, &redEnvelopes)
	return &redEnvelopes, err
}
