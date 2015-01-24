define('wsHandler', ['jquery', 'underscore', 'backbone'], function($, _, Backbone) {

	var wsUrl = 'ws://192.168.1.25:800/ws';

	var wsHandler = {};

	_.extend(wsHandler, Backbone.Events);

	wsHandler.on('message', function(msg){
		handleIncomingMessage(msg);
	});

	wsHandler.on('error', function(err){
		console.log(err);
	});

	wsHandler.on('opened', function(msg){

		//console.log('starting auth');
		wsHandler.authenticate('test', 'test');
		
	});


	var handleIncomingMessage = function(msg){
		var data = JSON.parse(msg.data);

		switch(data.type){
			case 'authentication_response':
					wsHandler.trigger('authentication_response', data.response, data.user);
				break;

			default:
				//console.log('Unknown data type', data.type, data);
				break;
		}
	};


	wsHandler.authenticate = function(username, password){
		wsHandler.socket.sendJSON({type: 'authentication_request', username: username, user_token: password});
	};




	wsHandler.on('all', function(){
		console.log(arguments);
	});


	//Start the socket
	wsHandler.socket = new WebSocket(wsUrl);
	wsHandler.socket.sendJSON = function(message){ wsHandler.socket.send(JSON.stringify(message)) };
 
	wsHandler.socket.onopen = function(evt) { wsHandler.trigger('opened', evt); }; 
	wsHandler.socket.onclose = function(evt) { wsHandler.trigger('closed', evt); }; 
	wsHandler.socket.onmessage = function(evt) { wsHandler.trigger('message', evt); }; 
	wsHandler.socket.onerror = function(evt) { wsHandler.trigger('error', evt); };

	return wsHandler;
});

