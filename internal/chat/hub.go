package chat

// Hub contain clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the client.
	broadcast chan []byte

	// Register requests from the client.
	Register chan *Client

	// Unregister request from the client.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[*Client]bool),
		broadcast: make(chan []byte),
		Register: make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h Hub) GetClients() map[*Client]bool {
	return h.Clients
}

func (h Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}