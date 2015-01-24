define(
	['marionette', 'text!templates/board.html',
	'app/collections/piecesCol'
	],
	function(Marionette, boardTemplate, PiecesCol)
	{
		var BoardView = Marionette.Layout.extend(
		{
			tagName:'tbody',
			template:_.template(boardTemplate),
			initialize:function(){
				_.bindAll(this, 'createPieces');

				// create pieces all over the place
				this.createPieces();

			},
			createPieces:function(){
				this.blackPieces = new PiecesCol([], {'color':'black'});
				this.whitePieces = new PiecesCol([], {'color':'white'});
			}
		});
		return BoardView;
	}
);