package main

type id = string

type Hub map[id]*Client

func NewHub() Hub {
	h := make(Hub)
	return h
}

func (h Hub) Append(key id, client *Client) {
	if _, ok := h[key]; !ok {
		h[key] = client
	}
}

func (h Hub) Pop(key id) {
	if _, ok := h[key]; ok {
		delete(h, key)
	}
}
