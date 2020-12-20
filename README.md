# ShakeSearch

Welcome to the Pulley Shakesearch Take-home Challenge! In this repository,
you'll find a simple web app that allows a user to search for a text string in
the complete works of Shakespeare.

You can see a live version of the app at
https://pulley-shakesearch.herokuapp.com/. Try searching for "Hamlet" to display
a set of results.

In it's current state, however, the app is just a rough prototype. The search is
case sensitive, the results are difficult to read, and the search is limited to
exact matches.

## Your Mission

Improve the search backend. Think about the problem from the user's perspective
and prioritize your changes according to what you think is most useful.

To submit your solution, fork this repository and send us a link to your fork
after pushing your changes. The project includes a Heroku Procfile and, in its
current state, can be deployed easily on Heroku's free tier.

If you are stronger on the front-end, complete the react-prompt.md in this
folder.

## Local Deploy
- > $ cp .env.example .env
- > $ sudo ./start
- > $ 2
- Then access: ```http://localhost:8200```

## Possible Improvements
- Back-end
    - Improve text parsing
    - Implement Elasticsearch to improve search performance
- Front-end
    - Fix some timing issues with components Animation
    - Implement Pagination

## Observations
- I'm not stronger on the front-end, I just was curious about VueJs and spent 2 days on the front-end for funsies
- I really prefer back-end dev
