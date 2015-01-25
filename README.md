<h1>Go-Llama Internet Chess Environment<h1>
<h2>(and AI proving ground)<h2>
<h3>Gopher Gala 2015 entry (48 hour programming challenge)</h3>
<p>This project is an internet chess server for players to be matched in real time with other players and AIs from around the world. </p>
<p>Its purpose is twofold: Firstly, to provide players with the ability to play chess through the internet through their browsers; and secondly, to provide a fully-featured API that developers can connect their own front-ends and (more excitingly) their own AIs to.</p>
<p>Thus, the project is first and foremost an _API server_ written entirely in the Go programming language. To showcase its abilities, we provide a _demo_ javascript frontend that players can use to play the game on. The API and server is represented by the *intchess* package.</p>
<p>In addition, we provide the *chessai* package, which is a basic framework for players to define their own AIs on without needing to worry about boilerplate API development.</p>
<p>In order for these two packages to function, a third package called *chessverifier* was designed and built. This package's sole function is to validate chess moves, which proved to be a more complex task than we initially thought.</p>
<p>Our demo Javascript files can be found in /static. It was constructed using Bootstrap, Backbone.js, Underscore.js, Marionette.js, Require.js and jQuery.</p>
<p>The Internet Chess server depends on a number of publically available open-source Go libraries, including:
<ul>
<li>Gorm - github.com/jinzhu/gorm</li>
<li>Websocket - code.google.com/p/websocket</li>
<li>Bcrypt - code.google.com/p/go.crypto/bcrypt</li>
<li>Mysql - github.com/go-sql-driver/mysql</li>
</ul>