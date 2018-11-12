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
  render()
})

