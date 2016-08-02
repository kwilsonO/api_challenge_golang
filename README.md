#Dogbreed API


//dogbreed api endpoints
gmux.HandleFunc("/dogs", ListDogs).Methods("GET")

gmux.HandleFunc("/dogs/breed/{Breed}", GetDogsByBreed).Methods("GET")

gmux.HandleFunc("/dogs/{URL}", GetDog).Methods("GET")
gmux.HandleFunc("/dogs/{URL}", SetDog).Methods("PUT")
gmux.HandleFunc("/dogs/{URL}", AddDog).Methods("POST")
gmux.HandleFunc("/dogs/{URL}", DeleteDog).Methods("DELETE")

//maybe instead:
//dogs/{URL}/?upvote=true
//dogs/{URL}/?downvote=true
//dogs/{URL}/?favorite=true
gmux.HandleFunc("/dogs/favorite/{URL}", FavoriteDog).Methods("PUT")
gmux.HandleFunc("/dogs/upvote/{URL}", UpvoteDog).Methods("PUT")
gmux.HandleFunc("/dogs/downvote/{URL}", DownvoteDog).Methods("PUT")

