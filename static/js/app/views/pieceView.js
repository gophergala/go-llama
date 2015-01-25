define([
	'marionette',
	'app/models/piece',
	'text!templates/piece.html',
	'wsHandler'
	],
	function(Marionette, Piece, PieceTemplate, wsHandler){
		var PieceView = Marionette.ItemView.extend({
			className:'piece',
			template:_.template(PieceTemplate),
			initialize:function(){
				_.bindAll(this, 'revert');
			},
			onRender:function(){
				this.$el.draggable({
					snap: 'td',
					snapModel:'inner',
					opacity: 0.8,
					distance: 10,
					revert:this.revert
				});

				this.$el.css('position','absolute');
				loc = this.model.get('location');
				this.$el.offset({left: loc[0] * 52, top: (9 - loc[1]) * 52});
			},
			revert:function(socketObj){
				if(socketObj){
					// Drag success - this would be where we trigger submitting the move
					var dataset = socketObj[0].dataset;
					var newrow = dataset.row;
					var newcol = dataset.col;
					var loc = [newcol, newrow];
					var oldloc = this.model.get('location');
					var oldcol = oldloc[0];

					var oldchr = String.fromCharCode(96 + parseInt(oldcol));
					var newchr = String.fromCharCode(96 + parseInt(newcol));

					this.model.set('location', loc);

					var moveString = oldchr + oldloc[1] + '-' + newchr + loc[1];
					$('#movelog').append(moveString + '<br>');

					wsHandler.moveRequest(moveString);

					console.log('success!');
					return false;
				}
				else {
					// drag fail
					console.log('fail!');
					return true;
				}
			}
		});
		return PieceView;
	}
);