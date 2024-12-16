# Welcome to Go, Templ and Htmx!

This workshop project supports "warm" reload with Air. You can just run
```sh
air
```
to run the project and have it reload on every change.  
> 💡 **Note:** This still means that you will have to refresh the page in your browser, hence the "warm" reload.

**Open [`localhost:8080`](http://localhost:8080) and see the website!**

# Exercises
We're gonna keep it simple and create a small todo app in this workshop. This will hopefully teach you the basics of Go and Templ and specifically how HTMX works and saves you from having to use Js.

## Exercise 1 - Create a new templ component
First, let us get more comfortable with Templ (and Go). You can see we have an Index page already under `views/index.templ`. Take a look at that. After that, have a look at how Go serves this page in `main.go` (line 11 to 13).  
**Analogous to that, can you create a simple `/about` page?**

> 💡 **Note:** When creating a new view about.templ, you have to make sure to capitalize the templ function name (e.g. `func Index()` in `index.templ` ) for it to visible to other packages. This is because in Go, only fields that are capitalized are considered as public, while those that aren't are package private.

--------

## Exercise 2 - Getting html using htmx
Now that you have a simple about page, try to link to the about page from the index page using HTMX. You can use the `hx-get="/some_endpoint"`-html keyword to issue an HTTP GET request to the server.  

In your `index.templ` the `<a>` in line 25 links to the about page. Currently the `href` attribute is empty. Let's switch that `href` attribute in `<a>` to `hx-get="/about"`.

What is happening? Use the developer tools of your browser to inspect the request and response to and from the server under the "Network" tab. (Reload the page first)

How can we fix it?

--------

## Exercise 3 - Creating a todo component
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

--------

## Exercise 4 - Displaying a list of todo items
After creating a single todo component, we now want to **display a whole list** of them! After all, you rarely only have one task to do, right?  

**1.** Start by creating a simple array of strings in your `main.go` file. These strings will hold the names of our todo items. Be creative!

> 💡 You can create an array of strings in Go with `var todos = []string{"Buy milk", "Your mom"}`

**2.** Now create another Templ component `templ TodoList(todos []string)`, that takes an array of strings as parameter. You can just add it to the already existing `todo.templ` file.
In your component, use a for-loop to create one Todo component for each item in the string array from our parameter.

> 💡 You can just write plain Go in your templ components. A for loop can look as simple as `for _, item := range items {}`
> Also, you can insert other templ components in your component by calling the function name with an @ in front, like `@Todo()`.

**3.** Change your `/todo` handler in main.go to return the TodoList component with the array of todo items you just created as the parameter.

**4.** Like in exercise 2, use the "Add Todo" button from the index page to issue a get-request to `/todo`. Don't forget the `hx-target`!

--------

## Exercise 5 - Adding structure
**1.** Let's give our Todo Items structure by creating a `TodoItem` struct in `todo.templ`. Since right now it is just an array of strings in `main.go`.  
The struct needs an `Id`, a `Title`, and a bool checking if it is `Done`. 

> 💡 You can create a struct in Go with `type YourStructName struct {}`. Add fields to it with `fieldName type`, separated by line breaks.

**2.** Change your TodoList component to take a `[]TodoItem` array as a parameter instead of an array of strings. Careful: You will have to change how you call your singular `Todo` component also.

--------

## Exercise 6 - Handling our first post request
Let's add the functionality to **create** a Todo Item! You can use an HTML form to do so. We will keep it simple in the beginning. 

**1.** Add an HTML form (Look at the second tip below if you need help with that) to your TodoList component. It should have an input field with the attribute `name="title"`.  
The form should send an hx-post to `/todos` and target `#content`. You can use these attributes instead of the default `method=POST` attribute for forms. Also, it's always a good idea to create new templ components for items like this.

**2.** Handling this post request in your `main.go` will require you to do some things which we have not tackled yet, mainly the `http.HandleFunc` method to handle custom logic instead of simply returning a file. To make things easier, just copy and paste the following code into your `main.go`:
```go
http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
    title := r.FormValue("title")
    newTodo := views.TodoItem{Id: len(todos) + 1, Title: title, Done: false}
    todos = append(todos, newTodo)

    w.WriteHeader(http.StatusCreated)
    views.TodoList(todos).Render(r.Context(), w)
})
```
> This code will extract the "title" from the form, and create a TodoItem that will then be appended to the array of todo items. 
> Finally, it will return the TodoList component with the new todo array and the status code 201 (successfully created).  
> Note the complete lack of error handling in the above code! This is one of the things we are going to skip for now, to keep it simpler.

> 💡 2nd Tip: In case it's been a while since you created your last HTML form: They look something like this:
```html
<form method=POST>
    <input type="text" placeholder="Add a new todo"/>
    <button type="submit">Add</button>
</form>
```

--------

## Exercise 7 - Handling out-of-bands updates
There are some problems with this though: Check the response we get under the network tab as we add more and more todo items.  
As you can see, the server responds with the whole todo list every time. This is fine in our example, but you can imagine that this would possibly create unmanageably huge responses with larger lists. How can we fix this?

We have not spoken about out-of-bands updates yet. Adding The `hx-swap-oob` to an element will allow htmx to swap this element in addition to the element that issued the hx-request. With this, you can swap more than one component in multiple places in your document.

**1.** Create another templ component named `TodoItemOob`which simply wraps the TodoItem component in a div with the following tags: `id="todo-list" hx-swap-oob="afterbegin"`.  
Make sure that the div in your TodoList component that wraps the for loop has the id `todo-list`. Also make sure that your form has the attribute `hx-swap="none"` set. Do you know why?

--------

## Exercise 8 - Deleting the items
What's missing now is the ability to delete Todo Items.
**You should know what to do by now: Create a button with an hx-delete request to the server.**  
You can use this svg as your delete button:
```html
<svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
</svg>
```
Use the ID of the todo item as a parameter in the request path.  
This will require you to convert the Integer `Id` into a string. You can use the `strconv.Itoa(Id)` to do so. 
You can concatenate the strings in Templ like so: `{"/todos/" + strconv.Itoa(todo.Id)}`

> 💡 We are using plain `net/http` (The web server from the Go standard library), so you will have to extract the Id from the string manually in your Go code. Web frameworks like Echo handle this automatically. How exactly to do that is left as an exercise for the reader.
> Another thing to watch out for is that `http.handleFunc("/todos")` does not handle `/todos/` with the extra slash at the end. You will have to create a separate endpoint for that.

There are some things you need to watch out here for. One thing are the hx-target selectors: You can not only use CSS selectors, but combine them with `closest`, `next` or `previous` to select the item you want, for example `next h1`. 
You can learn more [here](https://htmx.org/attributes/hx-target/).  

--------

## Exercise 9 - Completing items
If you got this far, and there is still time, here is one more exercise for you:  
**Handle the completion of todo items.**  
Once the user checks the checkbox, make it so that the item is sent to the bottom of the list, optionally greyed out and the text striked through.  
There won't be any additional tips for this one. Good luck!