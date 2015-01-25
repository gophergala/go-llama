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
				_.bindAll(this, 'createPieces', 'createSet', 'updateBoard', 'columnLoop', 'rowLoop');
				wsHandler.on('game_move_update', this.updateBoard);

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
				boardStatus = game.board_status;
				_.each(boardStatus, this.columnLoop);
			},
			columnLoop:function(column, colnum){
				this.columnNumber = colnum + 1;
				_.each(column, this.rowLoop);
			},
			rowLoop:function(cell, rownum){
				this.rowNumber = rownum + 1;
				// console.log(colnum);
				// console.log(rownum);
				console.log(this.columnNumber);
				console.log(this.rowNumber);
				console.log(cell);
				var decodedCell = window.atob(cell);
				console.log(decodedCell);

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
			onRender:function(){
				this.blackPiecesRegion.show(this.blackPiecesView);
				this.whitePiecesRegion.show(this.whitePiecesView);
			}
		});
		return BoardView;
	}
);