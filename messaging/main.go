package messaging

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ws-chart/core"
	"github.com/ws-chart/datastore"
	"github.com/ws-chart/messaging/client"
	"github.com/ws-chart/messaging/rpc"
	"github.com/ws-chart/protocol"
)

const (
	writeWait  = 30 * time.Second
	pingPeriod = 10 * time.Second
)

func Start(r *mux.Router) {
	go publishListener()
	r.Methods("GET").Path("/{client_id}").HandlerFunc(handler)
}

func publishListener() {
	datastore.Redis.Subscribe(func(channel string, data []byte) {
		chunks := strings.Split(channel, ":")
		clientID := chunks[len(chunks)-1]
		conn, err := client.ConnectionByID(clientID)
		if err != nil {
			core.Logger.Errorf("MESSAGING: Subscribe error %s %v %v", clientID, string(data), err)
			return
		}
		core.Logger.Infof("MESSAGING: Subscribe %s %v", clientID, string(data))
		conn.SetWriteDeadline(time.Now().Add(writeWait))
		conn.WriteMessage(websocket.TextMessage, data)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := client.NewFromRequest(params["client_id"], w, r)
	if client == nil {
		core.Logger.Error("Unauthorized")
		return
	}

	ch := make(chan *protocol.RPC)
	go dispatcher(client, ch)
	reader(client, ch)
	close(ch)
}

func dispatcher(c *client.Client, ch chan *protocol.RPC) {
	core.Logger.Infof("MESSAGING: Dispatcher started for client ID %s\n", c.ID)
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		core.Logger.Infof("MESSAGING: Disconnect client ID %s\n", c.ID)
		ticker.Stop()
	}()

	for {
		select {
		case r, ok := <-ch:
			if !ok {
				core.Logger.Infof("MESSAGING: Could not receive event from client ID %s", c.ID)
				c.SendCloseConnection()
				return
			}

			rpc.CallMethod(c, r)
		case <-ticker.C:
			err := c.SendPing()
			if err != nil {
				core.Logger.Errorf("MESSAGING: Could not send ping to client ID %s %s\n", c.ID, err)
				return
			}
		}
	}
}

func reader(c *client.Client, ch chan *protocol.RPC) {
	defer func() {
		c.Close()
		core.Logger.Errorf("MESSAGING: Disconnect reader for client ID %s\n", c.ID)
	}()

	c.Setup()

	for {
		rpc, err := c.ReadRPC()
		if err != nil {
			core.Logger.Errorf("MESSAGING: Error in event from client ID %s: %v", c.ID, err)
			return
		}

		core.Logger.Infof("MESSAGING: Received event from client ID %s: %v", c.ID, rpc)

		ch <- rpc
	}
}
