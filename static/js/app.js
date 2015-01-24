console.log('starting app.js');


var app = new Marionette.Application();

app.addRegions({
	board: '#board'
});

app.addInitializer(function(){

});

app.start();


console.log('app.js started');