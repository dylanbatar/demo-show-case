package main

type Person struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	Bio      string  `json:"bio"`
	Version  float32 `json:"version"`
}

type Persons struct {
	Persons []Person `json:"persons"`
}
