package ws

import (
	"net/http"

	"github.com/dopedao/dope-monorepo/packages/api/game"
	"github.com/dopedao/dope-monorepo/packages/api/internal/dbprovider"
	"github.com/rs/zerolog/log"

	"github.com/dopedao/dope-monorepo/packages/api/internal/middleware"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

func HandleConnection(game *game.Game) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		client, err := middleware.FirebaseInit(r.Context())
		if err != nil {
			log.Err(err).Msgf("firebase init error")
			http.Error(w, "firebase init error", http.StatusInternalServerError)
			return
		}
		// check if authenticated
		if !middleware.IsAuthenticated(r.Context(), client) {
			http.Error(w, "not authenticated", http.StatusUnauthorized)
			return
		}

		wsConn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the connection when the function returns
		defer wsConn.Close()

		// handle messages
		game.HandleGameMessages(r.Context(), dbprovider.Ent(), wsConn)
	}
}
