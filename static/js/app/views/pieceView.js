define([
	'marionette',
	'app/models/piece',
	'text!templates/piece.html'
	],
	function(Marionette, Piece, PieceTemplate){
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
					this.model.set('location', loc);
					$('#movelog').append('[' + oldloc[0] + ',' + oldloc[1] + '],[' + loc[0] + ',' + loc[1] + ']<br>');
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