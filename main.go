package main

func main() {
	// Setup query
	/*
		r := http.Client{}

		resp, _ := r.Get("https://api.jikan.moe/v4/manga/1")

		jb := Tags{}

		err := json.NewDecoder(resp.Body).Decode(&jb)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v", jb)
	*/

	server()

	// Search anime

}
