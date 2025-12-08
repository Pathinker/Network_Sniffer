package protocols

type ParsedProtocol struct {
    Name    string                 `json:"name"`
    Header  map[string]interface{} `json:"header"`
    Payload []byte                 `json:"payload"`
}

type Protocol interface {
    Parse(data []byte) (*ParsedProtocol, error)
}
