define(
	['marionette', 'text!templates/board.html',
	'app/collections/piecesCol',
	'app/views/piecesColView',
	'app/models/pawn',
	'app/models/rook',
	'app/models/knight',
	'app/models/bishop',
	'app/models/queen',
	'app/models/king'
	],
	function(Marionette, boardTemplate, PiecesCol, PiecesColView, Pawn, Rook, Knight, Bishop, Queen, King)
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
				_.bindAll(this, 'createPieces', 'createSet');

				// create pieces all over the place
				this.createPieces();

			},
			createPieces:function(){
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

				this.createSet('black', this.blackPieces);
				this.createSet('white', this.whitePieces);
				

			},
			createSet: function(color, collection){
				var pawnrow = 7;
				var mainrow = 8;
				if(color == 'white'){
					pawnrow = 2;
					mainrow = 1;
				}

				for (var i = 1; i <= 8; i++) {
					collection.add(new Pawn({'location':[i, pawnrow], 'color':color}));
				}

				collection.add(new Rook({'location':[1, mainrow], 'color':color}));
				collection.add(new Rook({'location':[8, mainrow], 'color':color}));

				collection.add(new Knight({'location':[2, mainrow], 'color':color}));
				collection.add(new Knight({'location':[7, mainrow], 'color':color}));

				collection.add(new Bishop({'location':[3, mainrow], 'color':color}));
				collection.add(new Bishop({'location':[6, mainrow], 'color':color}));

				collection.add(new Queen({'location':[4, mainrow], 'color':color}));
				collection.add(new King({'location':[5, mainrow], 'color':color}));
			},
			onRender:function(){
				this.blackPiecesRegion.show(this.blackPiecesView);
				this.whitePiecesRegion.show(this.whitePiecesView);
				$('td').droppable({
					drop: function(){
						alert('dropped');
					},
					accept: '.piece'
				});
			}
		});
		return BoardView;
	}
);