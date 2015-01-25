define([
	'marionette',
	'wsHandler',
	'jquery',
	'text!templates/interactionWindow/loading.html',
	'text!templates/interactionWindow/loginForm.html',
	'text!templates/interactionWindow/registerForm.html',
	'text!templates/interactionWindow/gameRequest.html',
	'text!templates/interactionWindow/generic.html'
	],
	function(Marionette, wsHandler, $, loadingTpl, loginTpl, registerTpl, gameRequest, genericTpl){


		var tplData = {};

		var changeView = function(tpl, newData){
			tplData = newData || {};

			iw.template = _.template(tpl);
			iw.render();
		}


		var InteractionWindow = Marionette.Layout.extend({
			template: _.template(loadingTpl),

			render: function() {
				this.$el.html(this.template(tplData));
				return this;
			},

			events: {
				'click #login-button': 'doLogin',
				'click #i-want-to-register': 'showRegisterView',
				'click #i-want-to-login': 'showLoginView',
				'click #register-button': 'doRegister',
				'click #game-accept': 'acceptGame',
				'click #game-deny': 'denyGame'
			},
			

			doLogin: function(e){
				e.preventDefault();
				
				var username = $('#login-form input[name="username"]').val();
				var password = $('#login-form input[name="password"]').val();
				
				changeView(loadingTpl);

				return wsHandler.authenticate(username, password);
			},

			showRegisterView: function(){
				return changeView(registerTpl);
			},

			showLoginView: function(){
				return changeView(loginTpl);
			},

			doRegister: function(e){
				e.preventDefault();
				
				var username = $('#register-form input[name="username"]').val();
				var password = $('#register-form input[name="password"]').val();
				var ai = $('#register-form input[name="verseai"]').is(':checked');
				
				changeView(loadingTpl);

				return wsHandler.register(username, password, ai);
			},

			acceptGame: function(e){
				wsHandler.gameResponse(true);
				return changeView(loadingTpl, {msg: 'Waiting for other player to accept or deny'});
			},

			denyGame: function(e){
				wsHandler.gameResponse(false);
				return changeView(loadingTpl, {msg: 'Waiting for game'});
			}




		});

		var iw = new InteractionWindow();

		var showLogin = function(){ changeView(loginTpl); };


		//If we haven't connected yet, we need to wait for the opened event
		//If we have, just show the login.
		if(wsHandler.connected){
			showLogin();
		}else{
			wsHandler.on('opened', showLogin);
		}


		wsHandler.on('authentication_response', function(code, user){
			if(code != 'ok'){
				return changeView(loginTpl, {error: 'Login incorrect'});
			}

			return changeView(loadingTpl, {msg: 'Waiting for game'});

		});

		wsHandler.on('signup_response', function(code, user){
			if(code != 'ok'){
				return changeView(registerTpl, {error: 'Registration incorrect'});
			}

			return changeView(loadingTpl, {msg: 'Waiting for game'});

		});

		console.log(iw);


		wsHandler.on('game_request', function(opponent){


			//Only do this once if the game starts
			wsHandler.once('game_move_update', function(opponent){
				return changeView(genericTpl, {msg: 'Game started!'});
			});

			return changeView(gameRequest, {opponent: opponent});
		});

		wsHandler.on('game_response_rejection', function(opponent){
			return changeView(loadingTpl, {msg: 'Last game did not start.<br>Waiting for game'});
		});


		return iw;
	}
);