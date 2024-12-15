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

### Exercise 1 - Create a new templ component
First, let us get more comfortable with Templ (and Go). You can see we have an Index page already under `views/index.templ`. Take a look at that. After that, have a look at how Go serves this page in `main.go` (line 11 to 13).  
**Analogous to that, can you create a simple `/about` page?**

> 💡 **Note:** When creating a new view about.templ, you have to make sure to capitalize the templ function name (e.g. `func Index()` in `index.templ` ) for it to visible to other packages. This is because in Go, only fields that are capitalized are considered as public, while those that aren't are package private.

### Exercise 2 - Getting html using htmx
Now that you have a simple about page, try to link to the about page from the index page using HTMX. You can use the `hx-get="/some_endpoint"`-html keyword to issue an HTTP GET request to the server.  

In your `index.templ` the `<a>` in line 25 links to the about page. Currently the `href` attribute is empty. Let's switch that `href` attribute in `<a>` to `hx-get="/about"`.

What is happening? Use the developer tools of your browser to inspect the request and response to and from the server under the "Network" tab. (Reload the page first)

How can we fix it?

### Exercise 3 - Creating a todo component
Now that we fixed our bug, let's start actually implementing our todo app!  

**1.** Create a new templ component `todo.templ` with the function `templ Todo(todo string)`. Let it display a singular todo Item.

**2.** Include a checkbox in your `todo.templ`to check whether the item has been completed. 

> 💡 You can use `<input type="checkbox" />` to define a checkbox in html

**3.** Add another http route `/todo` to your `main.go`that returns your new `todo.templ` component. It will require a string as parameter, e.g. "Buy milk". Then check if your component works by visiting `localhost:8080/todo`

The code of your final Todo component should look something like this:

```go
templ Todo(todo string) {
    <div class="todo"> 
        <input type="checkbox"/>
        {todo} 
    </div>
}
```

### Exercise 4 - Displaying a list of todo items
After creating a single todo component, we now want to **display a whole list** of them! After all, you rarely only have one task to do, right?  

**1.** Start by creating a simple array of strings in your `main.go` file. These strings will hold the names of our todo items. Be creative!

> 💡 You can create an array of strings in Go with `var todos = []string{"item1", "item2"}`

**2.** Now create another Templ component `templ TodoList(todos []string)`, that takes an array of strings as parameter. You can just add it to the already existing `todo.templ` file.
In your component, use a for-loop to create one Todo component for each item in the string array from our parameter.

> 💡 You can just write plain Go in your templ components. A for loop can look as simple as `for _, item := range items {}`
> Also, you can insert other templ components in your component by calling the function name with an @ in front, like `@Todo()`.

**3.** Change your "/todo" handler in main.go to return the TodoList component with the array of todo items you just created as the parameter.

**4.** Like in exercise 2, use the "Add Todo" button from the index page to issue a get-request to `/todo`. Don't forget the `hx-target`!

### Exercise 5
Let's do two more things before we really get into htmx and adding functionality to our Todo App.  

**First off, let's give our Todo Items structure by adding a TodoItem struct.** Right now it is just an array ow strings.  
The struct needs an `Id`, a `Title`, and a bool checking if it is `Done`.

> 💡 You can create a struct in Go with `type YourStructName struct {}`. You can just put the struct definition into todo.templ.

### Exercise 6
Let's add the functionality to create a Todo Item! You can use an HTML form to do so. We will keep it simple in the beginning.  
**Add a form to your TodoList component with an input field named "title" which will send an hx-post to `/contacts` and target `#content`. You will also need to handle the related Server code.**  
There are some problems with this though: Check the response we get under the network tab as we add more and more todo items.  
**To fix this, create another templ component named `TodoItemOob`which simply wraps the TodoItem component in a div with the following tags: `id="todo-list" hx-swap-oob="afterbegin"`.**  
Make sure that the div in your TodoList component that wraps the for loop has the id `todo-list`. Also make sure that your form has the attribute `hx-swap="none"` set. Why?

### Exercise 7
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

### Exercise 8
If you got this far, and there is still time, here is one more exercise for you:  
**Handle the completion of todo items.**  
Once the user checks the checkbox, make it so that the item is sent to the bottom of the list, optionally greyed out and the text striked through.  
There won't be any additional tips for this one. Good luck!