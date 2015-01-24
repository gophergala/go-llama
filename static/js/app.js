define(
	['marionette', 'app/views/boardView'], 
	function(Marionette, Board){
		var app = new Marionette.Application();

		app.addRegions({
			board: '#board-div'
		});

		app.addInitializer(function(){
			app.board.show(new Board());
		});

		app.start();

		return app;
	}
);

		