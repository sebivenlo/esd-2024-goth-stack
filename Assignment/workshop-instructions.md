## Welcome to Go, Templ and Htmx!

This workshop project supports "warm" reload with Air. You can just run
```sh
air
```
to run the project and have it reload on every change.  
> **Note:** This still means that you will have to refresh the page in your browser, hence the "warm" reload.

**Open [`localhost:8080`](http://localhost:8080) and see the website!**

## Exercises
### Exercise 1
First, let us get more comfortable with Templ (and Go). You can see we have an Index page already under `views/index.templ`. Take a look at that. After that, have a look at how Go serves this page in `main.go`.  
**Analogous to that, can you create a simple `/about` page?**

### Exercise 2
Now that you have a simple about page, try to link to the about page from the index page using HTMX. You can use the "hx-get"-html keyword to GET a resource from the server.  

What is happening? Use the developer tools to expect the request and response to and from the server under the "Network" tab. (Reload the page first)