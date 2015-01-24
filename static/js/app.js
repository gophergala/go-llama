define(
	['marionette', 'app/views/boardView'], 
	function(Marionette, Board){
		console.log('starting app.js');


		var app = new Marionette.Application();

		app.addRegions({
			board: '#board'
		});

		app.addInitializer(function(){
			app.board.show(new Board());
		});

		app.start();


		console.log('app.js started');

		return app;
	}
);

		