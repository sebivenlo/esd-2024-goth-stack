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

> 💡 You can create a struct in Go with `type YourStructName struct {}`. You can just put the struct definition into todo.templ.

**Secondly, create another Templ component. This one will be called `TodoList` and take a TodoItem array as a parameter.**  
It should wrap around and create one TodoItem component for each item in the TodoItem array from our parameter.

> 💡 You can just write plain Go in your templ components. A for loop can look as simple as `for _, item := range items {}`

### Exercise 5
Let's add the functionality to create a Todo Item! You can use an HTML form to do so. We will keep it simple in the beginning.  
**Add a form to your TodoList component with an input field named "title" which will send an hx-post to `/contacts` and target `#content`. You will also need to handle the related Server code.**  
There are some problems with this though: Check the response we get under the network tab as we add more and more todo items.  
**To fix this, create another templ component named `TodoItemOob`which simply wraps the TodoItem component in a div with the following tags: `id="todo-list" hx-swap-oob="afterbegin"`.**  
Make sure that the div in your TodoList component that wraps the for loop has the id `todo-list`. Also make sure that your form has the attribute `hx-swap="none"` set. Why?

### Exercise 6
What's missing now is the ability to delete Todo Items now.
You can use this svg as your delete button:
```
<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
</svg>
```
**You should know what to do by now: Create a button with an hx-delete request to the server.**  
Use the ID of the todo item as a parameter in the request path. 
> 💡 We are using plain `net/http` (The web server from the Go standard library), so you will have to extract the Id from the string manually in your Go code. Web frameworks like Echo handle this automatically.

There are some things you need to watch out here for. One thing are the hx-target selectors: You can not only use CSS selectors, but combine them with `closest`, `next` or `previous` to select the item you want. 
You can learn more [here](https://htmx.org/attributes/hx-target/).  
Another thing is that http.handleFunc("/todos") does not handle `/todos/` with the extra slash at the end. You will have to create a separate endpoint for that.

### Exercise 7
If you got this far, and there is still time, here is one more exercise for you:  
**Handle the completion of todo items.**  
Once the user checks the checkbox, make it so that the item is sent to the bottom of the list, optionally greyed out and the text striked through.  
There won't be any additional tips for this one. Good luck!