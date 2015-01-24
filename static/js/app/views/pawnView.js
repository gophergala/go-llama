define([
	'marionette',
	'app/models/pawn',
	'text!templates/pawn.html'
	],
	function(Marionette, Pawn, PawnTemplate){
		var PawnView = Marionette.ItemView.extend({
			className:'pawn',
			template:_.template(PawnTemplate)
		});
		return PawnView;
	}
);