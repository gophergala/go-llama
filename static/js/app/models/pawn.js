define([
	'marionette', 'app/models/piece',
	'app/views/pawnView'
	],
	function(Marionette, Piece, PawnView){
		var Pawn = Piece.extend({
			initialize:function(){
				this.view = new PawnView();
			}
		});
		return Pawn;
	}
);