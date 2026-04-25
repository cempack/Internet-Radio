package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getFavoritesPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".local", "share", "radiodrift")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "favorites.json"), nil
}

func LoadFavorites() (map[string]bool, error) {
	path, err := getFavoritesPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]bool), nil
		}
		return nil, err
	}

	var favs []string
	if err := json.Unmarshal(data, &favs); err != nil {
		return nil, err
	}

	favMap := make(map[string]bool)
	for _, id := range favs {
		favMap[id] = true
	}

	return favMap, nil
}

func SaveFavorites(favMap map[string]bool) error {
	path, err := getFavoritesPath()
	if err != nil {
		return err
	}

	var favs []string
	for id, isFav := range favMap {
		if isFav {
			favs = append(favs, id)
		}
	}

	data, err := json.MarshalIndent(favs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
