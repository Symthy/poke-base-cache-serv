package client

import (
	"github.com/mtslzr/pokeapi-go"
)

type PokeApi interface {
	Abilities() error
}

type PokeApiWrapper struct {
	limit int
}

func NewPokeApi() PokeApi {
	pokeapi.CacheSettings.CustomExpire = 30
	pokeapi.CacheSettings.UseCache = true
	return &PokeApiWrapper{
		limit: 100,
	}
}

func (p *PokeApiWrapper) Items() error {
	return p.getResources("item")
}

func (p *PokeApiWrapper) Moves() error {
	return p.getResources("move")
}

func (p *PokeApiWrapper) Abilities() error {
	return p.getResources("ability")
}

func (p *PokeApiWrapper) Nature() error {
	return p.getResources("nature")
}

func (p *PokeApiWrapper) Pokemons() error {
	return p.getResources("pokemon")
}

func (p *PokeApiWrapper) Pokemonforms() error {
	return p.getResources("pokemon-form")
}

func (p *PokeApiWrapper) PokemonSpecies() error {
	return p.getResources("pokemon-species")
}

func (p *PokeApiWrapper) getResources(reourceName string) error {
	result, err := pokeapi.Resource(reourceName, 0, p.limit)
	if err != nil {
		return err
	}

	for {
		if result.Next == "" {
			return nil
		}
		offset, err := extractOffset(result.Next)
		if err != nil {
			return nil
		}

		result, err = pokeapi.Resource(reourceName, offset, p.limit)
		if err != nil {
			return err
		}
	}
}
