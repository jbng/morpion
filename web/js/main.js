console.log("Noughts & Crosses starting...");

document.addEventListener('DOMContentLoaded', (e) => {
	const ws = new WebSocket("ws://localhost:3000/ws")
	ws.onerror = () => console.log('WebSocket error')
	ws.onopen = () => console.log('WebSocket connection established')
	ws.onclose = () => console.log('WebSocket connection closed')
    window.onbeforeunload = function () {
        ws.onclose = function () { }
        ws.send('close')
        ws.close()
    }
    let state, render

    ws.onmessage = function (event) {
		const message = event.data
		console.log(`Message received: ${message}`)
		let data
		try {
			data = JSON.parse(message)
		} catch (err) {
			return
		}

		if (data.event == 'update_game_state') {
			state = data.game_state
			render()
		}
	}


    state = {
        game_ended: false,
        you_win: false,
        your_color: "red",
        table: [
            'red', null, null,
            null, 'red', null,
            null, 'red', 'green'
        ],
        its_your_turn: true
    }

    render = () => {
        const appElement = new DOMParser().parseFromString(`
        <div class="game">
            <div class="me">Ready Player, You play as <span style="color: ${state.your_color}">${state.your_color}</span></div>
            <div class="game-table">
                ${state.table.map(
                (tile, k) => {
                    return `
						    <div data-index="${k}"
						        class="game-table-tile ${tile !== null ? `game-table-tile--checked game-table-tile--${tile}` : ''}">
                            </div>`
                }).join('')}
            </div>
                ${state.its_your_turn ? `<div class="your-turn">It's your turn</div>`
                : `<div class="your-turn">Waiting for other player move</div>`}
                ${state.game_ended ? `<div class="winner">
                                        ${state.you_win ? `You win !` : `You loose !`}
                                      </div>
                                      <button class="play-again">Play again</button>
                                   ` : ``}
            </div>`, "text/html").body.firstChild
        const winnerElement = appElement.querySelectorAll('.winner')
        const meElement = appElement.querySelectorAll('.me')
        const tileElements = appElement.querySelectorAll('.game-table-tile')
        const playAgainButton = appElement.querySelectorAll('.play-again')

        tileElements.forEach((tileElement) => {

            const tileIndex = tileElement.getAttribute('data-index')

            tileElement.addEventListener('click', (e) => {
                ws.send(JSON.stringify({
                    event: 'play_at_coordinate',
                    tile_index: tileIndex
                }))
            })
        })

        appElement.addEventListener('click', (e) => {
            if (e.target.classList.contains('play-again')) {
                ws.send(JSON.stringify({event: 'play_again'}))
            }
        })
        document.body.innerHTML = ''
        document.body.appendChild(appElement)
    }
    render()
})

