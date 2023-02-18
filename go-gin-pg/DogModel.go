package main

type Dog struct {
	tableName struct{} `pg:"dogs"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsSpotted bool     `json:"is_spotted" pg:"is_spotted"`
	Color     string   `json:"color"  pg:"color"`
	Breed     string   `json:"breed"  pg:"breed"`
}

// FindAllDogs Получить список собак.
func FindAllDogs() []Dog {
	var dogs []Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dogs).Select()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dogs
}

// CreateDog Создать собаку.
func CreateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).Insert()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// FindCatById Получить собаку по id.
func FindDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&dog).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

// DeleteCatById Удалить собаку по id.
func DeleteDogById(id string) Dog {
	var dog Dog
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&dog).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return dog
}

func UpdateDog(dog Dog) Dog {
	pgConnect := PostgresConnect()

	oldDog := FindDogById(dog.ID)

	oldDog.Name = dog.Name
	oldDog.IsSpotted = dog.IsSpotted
	oldDog.Color = dog.Color

	_, err := pgConnect.Model(&oldDog).
		Set("name = ?", oldDog.Name).
		Set("is_spotted = ?", oldDog.IsSpotted).
		Set("color = ?", oldDog.Color).
		Where("id = ?", oldDog.ID).
		Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldDog
}
