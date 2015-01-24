define(
	['app/views/boardView', 'app/views/interactionWindowView'], 
	function(Board, InteractionWindow){
		var app = new Marionette.Application();

		app.addRegions({
			board: '#board-div',
			interactionWindow: '#interaction-window'
		});

		app.addInitializer(function(){
			app.board.show(new Board());
			app.interactionWindow.show(new InteractionWindow());
			$('td').droppable({
				drop: function(){
					alert('dropped');
				},
				accept: '.piece',
				activeClass:'highlight'
			});
		});

		app.start();

		return app;
	}
);

		