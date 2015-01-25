define(
	['text!templates/board.html',
	'app/collections/piecesCol',
	'app/views/piecesColView',
	'app/models/pawn',
	'app/models/rook',
	'app/models/knight',
	'app/models/bishop',
	'app/models/queen',
	'app/models/king',
	'wsHandler'
	],
	function(boardTemplate, PiecesCol, PiecesColView, Pawn, Rook, Knight, Bishop, Queen, King, wsHandler)
	{
		var BoardView = Marionette.Layout.extend(
		{
			// tagName:'tbody',
			regions:{
				tbody: '#board-body',
				blackPiecesRegion: '#black-pieces',
				whitePiecesRegion: '#white-pieces'
			},
			template:_.template(boardTemplate),
			initialize:function(){
				_.bindAll(this, 'updateBoard', 'columnLoop', 'rowLoop', 'showValidMoves');
				wsHandler.on('game_move_update', this.updateBoard);
				wsHandler.on('game_over', this.gameOver);
				wsHandler.on('game_get_valid_moves_response', this.showValidMoves);

				this.blackPieces = new PiecesCol({'color':'black'});
				this.whitePieces = new PiecesCol({'color':'white'});
				this.blackPiecesView = new PiecesColView({
					collection: this.blackPieces
				});
				this.whitePiecesView = new PiecesColView({
					collection: this.whitePieces
				});

				this.blackPieces.reset();
				this.whitePieces.reset();
			},
			updateBoard:function(game){
				this.blackPieces.reset();
				this.whitePieces.reset();

				console.log(game);
				console.log(wsHandler.user);
				if(wsHandler.user.username === game.white.username){
					window.YourColor = 'white';
				}
				else {
					window.YourColor = 'black';
				}
				if(game.game_moves){
					if(game.game_moves.length % 2 == 0){
						window.whichColor = 'white';
						window.WhosMove = game.white.username;
						$('#currentColor').html(game.white.username + '`s turn! (white)'); 
					}
					else {
						window.whichColor = 'black';
						window.WhosMove = game.black.username;
						$('#currentColor').html(game.black.username + '`s turn! (black)'); 
					}
				}
				else {
					window.whichColor = 'white';
					window.WhosMove = game.white.username;
					$('#currentColor').html(game.white.username + '`s turn! (white)'); 
				}

				boardStatus = game.board_status;
				_.each(boardStatus, this.columnLoop);
			},
			gameOver: function(game){

				var text = '';

				switch(game.game_status){
					case 'victory':
						if(wsHandler.user.user_id == game.winner.user_id){
							text += 'You win. Your rank changed by '+game.winner.rank_change;
						}else{
							text += 'You loose. Your rank changed by '+game.winner.rank_change;
						}
						break;

					case 'stalemate':
						text += 'Stalemate!';
						break;

					case 'disconnection':
						text += 'The other player disconnected.';
						break;

				}

				$('#currentColor').html('<p>' + text + '</p>');
			},
			columnLoop:function(column, colnum){
				this.columnNumber = colnum + 1;
				_.each(column, this.rowLoop);
			},
			rowLoop:function(cell, rownum){
				this.rowNumber = rownum + 1;
				if(window.YourColor === 'black'){
					this.rowNumber = 9 - this.rowNumber;	
				}
				// console.log(colnum);
				// console.log(rownum);
				// console.log(this.columnNumber);
				// console.log(this.rowNumber);
				// console.log(cell);
				var decodedCell = window.atob(cell);
				// console.log(decodedCell);

				var colorChar = decodedCell.charAt(0);
				var pieceChar = decodedCell.charAt(1);

				var color = 'white';
				if(colorChar == 'B'){
					color = 'black';
				}

				switch(pieceChar){
					case 'P':
						if(colorChar == 'B'){
							this.blackPieces.add(new Pawn({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new Pawn({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;
					case 'R':
						if(colorChar == 'B'){
							this.blackPieces.add(new Rook({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new Rook({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;
					case 'N':
						if(colorChar == 'B'){
							this.blackPieces.add(new Knight({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new Knight({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;
					case 'B':
						if(colorChar == 'B'){
							this.blackPieces.add(new Bishop({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new Bishop({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;
					case 'Q':
						if(colorChar == 'B'){
							this.blackPieces.add(new Queen({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new Queen({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;
					case 'K':
						if(colorChar == 'B'){
							this.blackPieces.add(new King({'location':[this.columnNumber, this.rowNumber], 'color':'black'}));
						}
						else if (colorChar == 'W'){
							this.whitePieces.add(new King({'location':[this.columnNumber, this.rowNumber], 'color':'white'}));
						}
						break;


				}
			},
			showValidMoves:function(moves){
				console.log(moves);
			},
			onRender:function(){
				this.blackPiecesRegion.show(this.blackPiecesView);
				this.whitePiecesRegion.show(this.whitePiecesView);
			}
		});
		return BoardView;
	}
);