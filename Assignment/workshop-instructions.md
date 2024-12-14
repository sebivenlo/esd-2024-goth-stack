## Welcome to Go, Templ and Htmx!

This workshop project supports "warm" reload with Air. You can just run
```sh
air
```
to run the project and have it reload on every change.  
> 💡 **Note:** This still means that you will have to refresh the page in your browser, hence the "warm" reload.

**Open [`localhost:8080`](http://localhost:8080) and see the website!**

## Exercises
We're gonna keep it simple and create a small todo app in this workshop. This will hopefully teach you the basics of Go and Templ and specifically how HTMX works and saves you from having to use Js.

### Exercise 1
First, let us get more comfortable with Templ (and Go). You can see we have an Index page already under `views/index.templ`. Take a look at that. After that, have a look at how Go serves this page in `main.go`.  
**Analogous to that, can you create a simple `/about` page?**

### Exercise 2
Now that you have a simple about page, try to link to the about page from the index page using HTMX. You can use the `hx-get="/some_endpoint"`-html keyword to issue an HTTP GET request to the server.  

What is happening? Use the developer tools to expect the request and response to and from the server under the "Network" tab. (Reload the page first)

### Exercise 3
Now that we fixed our bug, let's start actually implementing our todo app!  
**Start by creating a simple array of strings in our main.go file.**  
These strings will hold the names of our todo items.

> 💡 You can create an array of strings in Go with `var todos = []string{"item1", "item2"}`

**After that, create a new templ component that will display a singular Todo Item.**  
Also give it a checkbox to check whether the item has been completed.  
Check if it works by adding another http route `/todo` that returns this templ component.

### Exercise 4
Let's do two more things before we really get into htmx and adding  functionality to our Todo App.  
**First off, let's give our Todo Items structure by adding a TodoItem struct.**  
The struct needs an `Id`, a `Title`, and a bool checking if it is `Done`.
