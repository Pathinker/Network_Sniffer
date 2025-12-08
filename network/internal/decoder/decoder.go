package decoder

import (
    "sniffer/internal/protocols"
)

type Decoder struct{}

func NewDecoder() *Decoder {
    return &Decoder{}
}

func (d *Decoder) Decode(raw []byte) (map[string]*protocols.ParsedProtocol, error) {
    result := make(map[string]*protocols.ParsedProtocol)

    // L2
    l2, err := protocols.ParseEthernet(raw)
    if err != nil {
        return nil, err
    }
    result["l2"] = l2

    // L3
    l3, err := d.decodeL3(l2)
    if err != nil {
        return result, nil
    }
    result["l3"] = l3

    // L4
    l4, err := d.decodeL4(l3)
    if err != nil {
        return result, nil
    }
    result["l4"] = l4

    // L7
    l7, err := d.decodeL7(l4)
    if err == nil {
        result["l7"] = l7
    }

    return result, nil
}
