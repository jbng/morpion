console.log("Noughts & Crosses starting...");

const render = () => {
	const appElement = new DOMParser().parseFromString(`
			<div class="game">
				<div class="me">Ready Player One</div>
				<div class="game-table">
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
					<div class="game-table-tile"> </div>
				</div>
			</div>
				`, "text/html").body.firstChild
	document.body.innerHTML = ''
	document.body.appendChild(appElement)
}

document.addEventListener('DOMContentLoaded', (e) => {
	const ws = new WebSocket("ws://localhost:3000/ws")
	ws.onerror = () => console.log('WebSocket error')
	ws.onopen = () => {
		setInterval(
			function () { ws.send("ping") }
			, 1000)
	}
	ws.onclose = () => console.log('WebSocket connection closed')
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
	window.onbeforeunload = function () {
		ws.onclose = function () { }
		ws.send('close')
		ws.close()
	}
	render()
})

