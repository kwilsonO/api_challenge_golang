package zk

import (
	cfg "dogbreed/config"
	"errors"
)

/**

	Note just for PoC, would need to hook into ZK here.

**/

func ListDogs() ([]string, error) {

	return nil, nil
}

func ListDogs(breed string) ([]string, error) {

	//either call ListDogs() above
	//then filter on breed or implement on zk side

	return nil, nil
}

func GetDog(url string) (cfg.Dog, error) {
	if url == "" {
		return cfg.Dog{}, errors.New("Please specify a url")
	}

	return nil, nil
}

func SetDog(dog cfg.Dog) error {

	if dog.URL == "" {
		return errors.New("Please specify a URL")
	}

	return nil
}

func DeleteDog(url string) error {
	if url == "" {
		return errors.New("Please specify a url")
	}

	return nil
}

func FavoriteDog(url string) error {

	if url == "" {
		return errors.New("Please specify a url")
	}

	return nil
}

func UpvoteDog(url string) error {

	if url == "" {
		return errors.New("Please specify a url")
	}

	return nil

}

func DownvoteDog(url string) error {

	if url == "" {
		return errors.New("Please specify a url")
	}

	return nil
}
