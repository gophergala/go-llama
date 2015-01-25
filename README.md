##Go-Llama Internet Chess Environment (and AI proving ground)
___


###Gopher Gala 2015 entry (48 hour programming challenge)

This project is an internet chess server for players to be matched in real time with other players and AIs from around the world.

Its purpose is twofold: Firstly, to provide players with the ability to play chess through the internet through their browsers; and secondly, to provide a fully-featured API that developers can connect their own front-ends and (more excitingly) their own AIs to.

Thus, the project is first and foremost an *API server* written entirely in the Go programming language. To showcase its abilities, we provide a *demo* javascript frontend that players can use to play the game on. The API and server is represented by the **intchess** package.
In addition, we provide the **chessai** package, which is a basic framework for players to define their own AIs on without needing to worry about boilerplate API development.

In order for these two packages to function, a third package called **chessverifier** was designed and built. This package's sole function is to validate chess moves, which proved to be a more complex task than we initially thought.

Our demo Javascript files can be found in /static. It was constructed using Bootstrap, Backbone.js, Underscore.js, Marionette.js, Require.js and jQuery.

The Internet Chess server depends on a number of publically available open-source Go libraries, including:

- [Gorm](http://github.com/jinzhu/gorm)
- [Websocket](http://code.google.com/p/websocket)
- [Bcrypt](http://code.google.com/p/go.crypto/bcrypt)
- [Mysql](http://github.com/go-sql-driver/mysql)


The details of the API can be found in /static/chessapi.html