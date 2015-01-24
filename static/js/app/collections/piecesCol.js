define([
	'marionette',
	'app/models/piece',
	'app/models/pawn',
	'app/models/rook',
	'app/models/knight',
	'app/models/bishop',
	'app/models/queen',
	'app/models/king',
	],
	function(Marionette, Piece, Pawn, Rook, Knight, Bishop, Queen, King){
		var PiecesCol = Backbone.Collection.extend({
			model:Piece,
			initialize:function(models, options){
				var pawnrow = 2;
				var mainrow = 1;
				if(options.color == 'black'){
					pawnrow = 7;
					mainrow = 8;
				}
				for (var i = 1; i <= 8; i++) {
					this.add(new Pawn({'location':[i, pawnrow], 'color':options.color}));
				};

				this.add(new Rook({'location':[1, mainrow], 'color':options.color}));
				this.add(new Rook({'location':[8, mainrow], 'color':options.color}));

				this.add(new Knight({'location':[2, mainrow], 'color':options.color}));
				this.add(new Knight({'location':[7, mainrow], 'color':options.color}));

				this.add(new Bishop({'location':[3, mainrow], 'color':options.color}));
				this.add(new Bishop({'location':[6, mainrow], 'color':options.color}));

				this.add(new Queen({'location':[4, mainrow], 'color':options.color}));
				this.add(new King({'location':[5, mainrow], 'color':options.color}));

				// this.add(new)
			}
		});
		return PiecesCol;
	}
);