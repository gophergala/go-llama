define([
	'marionette',
	'app/models/piece',
	'text!templates/piece.html'
	],
	function(Marionette, Piece, PieceTemplate){
		var PieceView = Marionette.ItemView.extend({
			className:'piece',
			template:_.template(PieceTemplate),
			onRender:function(){
				loc = this.model.get('location');
				this.$el.offset({left: loc[0] * 52, top: (9 - loc[1]) * 52});
				this.$el.draggable();
			}
		});
		return PieceView;
	}
);