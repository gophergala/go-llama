define([
	'marionette',
	'text!templates/interactionWindow.html'
	],
	function(Marionette, InteractionWindowTemplate){
		var InteractionWindow = Marionette.Layout.extend({
			template:_.template(InteractionWindowTemplate)
		});
		return InteractionWindow;
	}
);