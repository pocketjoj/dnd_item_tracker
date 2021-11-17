Hello! Thank you for looking at my project. My go project is the starting point for a web server that will eventually serve as the back-end for a DnD inventory web page. This web server will be used with a database to allow character inventories to be viewed and updated. 

This Go Project meets the following project requirements (and then some, I believe):

1. Create a dictionary or list, populate it with several values, retrieve at least one value, and use it in your program
2. Read data from an external file, such as text, JSON, CSV, etc and use that data in your application
3. Connect to an external/3rd party API and read data into your app

I am counting that third one because my project's main goal is actually to BE an API that can take in or send data (updating either the web page or the database depending on if it is a POST or GET request). So while the project technically does not connect to an API, it is hosted on an external platform and it will write data (or allow it to be read).

To test this, you can either:
 - Download my project folder in its entirety, navigate to its directory on your terminal/command line and use the "go run ." command. OR
 - You can navigate to https://handyhaversack.herokuapp.com/ (as I have this app hosted on a heroku server).
 Note: If you run it from your computer, all commands I reference below will need to use http://localhost:5000/ as the base url instead of https://handyhaversack.herokuapp.com/.

 From there, you can make calls either using Postman (required for POST calls) or using the url (for simple GET requests).

 I have a database of 1293 items, so you can use https://handyhaversack.herokuapp.com/items/1176 (or any number between 1 and 1293) to pull up the json data for an item by its ID. 

 You can also query by name (e.g. https://handyhaversack.herokuapp.com/items/?name=greatsword), although if the name is not 100% accurate, you will receive an error. 

 Finally, you can test the POST functionality using the sample JSON below:

{"name":"Huge Halberd","sources":{"Homebrew":1},"rarity":"rare","entries":"This over-sized weapon is almost impossible to manage but devastating when it hits. Any creature with size Large or smaller has disadvantage on all att rolls, but all damage rolls are doubled.","attunement":false,"type":"Spear","custom":true,"id":1294}

To verify it is working, you can call:

https://handyhaversack.herokuapp.com/items/?name=huge halberd (or https://handyhaversack.herokuapp.com/items/1294 if it is the only item you have added) to see you new item. Feel free to add other items if you like.

Note: This is still a work in progress, but I believe it meets the requirements! :)

Thank you again!
Joel