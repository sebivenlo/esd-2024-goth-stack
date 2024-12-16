package main

import (
	"cookbook/views"
	"fmt"
	"github.com/a-h/templ"
	"math/rand"
	"net/http"
)

func main() {

	http.Handle("/", templ.Handler(views.Index()))
	http.Handle("/recipes", templ.Handler(views.RecipeList()))

	http.HandleFunc("/load-recipes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, err := w.Write([]byte(`
        <ul class="list-disc list-inside space-y-2">
            <li>Spaghetti Carbonara</li>
            <li>Chicken Curry</li>
            <li>Vegetarian Lasagna</li>
        </ul>
    `))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	})

	http.HandleFunc("/load-featured", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
			<div class="rounded-lg shadow-lg overflow-hidden bg-white border border-[#e0d6d1] hover-card">
				<img src="https://www.tinaskuechenzauber.de/wp-content/uploads/2019/10/IMG_7003-819x1024.jpg" class="w-full h-48 object-cover"/>
				<div class="p-6">
					<h3 class="text-2xl font-semibold text-[#8b5e3c]">Spiced Apple Pie</h3>
					<p class="mt-2 text-[#6b4f3c] text-sm">
						A traditional dessert with hints of cinnamon and nutmeg.
					</p>
				</div>
			</div>
<div class="rounded-lg shadow-lg overflow-hidden bg-white border border-[#e0d6d1] hover-card">
							<img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmYhiMOdF90dKiHYzQG5zx8aU9S3IF1d6FIw&s" alt="Rustic Sourdough Bread" class="w-full h-48 object-cover"/>
							<div class="p-6">
								<h3 class="text-2xl font-semibold text-[#8b5e3c]">Rustic Sourdough Bread</h3>
								<p class="mt-2 text-[#6b4f3c] text-sm">
									A crusty sourdough loaf, perfect with butter or jam.
								</p>
							</div>
						</div>
						<div class="rounded-lg shadow-lg overflow-hidden bg-white border border-[#e0d6d1] hover-card">
							<img src="https://www.leckerschmecker.me/wp-content/uploads/sites/6/2024/08/crumbl-cookies.jpg " alt="Chocolate Chip Cookies" class="w-full h-48 object-cover"/>
							<div class="p-6">
								<h3 class="text-2xl font-semibold text-[#8b5e3c]">Chocolate Chip Cookies</h3>
								<p class="mt-2 text-[#6b4f3c] text-sm">
									Crispy edges, gooey centers – a timeless classic.
								</p>
							</div>
						</div>
        </div>
    `))
	})

	http.HandleFunc("/random-recipe", func(w http.ResponseWriter, r *http.Request) {
		recipes := []string{
			"Spiced Apple Pie",
			"Rustic Sourdough Bread",
			"Chocolate Chip Cookies",
		}
		randomIndex := rand.Intn(len(recipes))
		recipe := recipes[randomIndex]

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<p class="text-lg font-bold text-[#8b5e3c]">%s</p>`, recipe)
	})

	// Start the server

	fmt.Println("Server is running on http://localhost:8080")

	http.ListenAndServe("localhost:8080", nil)
}
