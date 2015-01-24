require.config({
	paths: {
		jquery: 'lib/jquery/jquery',
		jqueryui: 'lib/jquery/jquery-ui.min',
		underscore: 'lib/underscore/underscore',
		backbone: 'lib/backbone/backbone',
		marionette: 'lib/backbone/backbone.marionette',
		templates: '../templates',
		app:'app'
	},
	shim: {
		jquery : {
			exports: '$'
		},
		jqueryui : {
			deps: ['jquery']
		},
		underscore : {
			exports : '_'
		},
		backbone : {
			exports : 'Backbone',
			deps : ['jquery', 'underscore']
		},
		marionette : {
			exports : 'Marionette',
			deps : ['backbone']
		}
	}	
});

require([
		'app', 'jquery', 'jqueryui', 'underscore', 'backbone', 'marionette'
	],
	function(App, $, _, jqueryui, Backbone, Marionette){
		App.start();
	}
);