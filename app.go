package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

type pokemon struct {
	Name    string `json:"name"`
	Sprites struct {
		Front_default string `json:"front_default"`
	} `json:"sprites"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowSetDarkTheme(a.ctx)
}

// Greet returns a greeting for the given name
func (a *App) GetPokemon() []string {

	// random number between 1 and 898

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	number := randGen.Intn(897) + 1

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + fmt.Sprintf("%d", number))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var p pokemon

	json.Unmarshal(body, &p)

	if strings.Contains(p.Name, "-") {
		return a.GetPokemon()
	}

	return []string{p.Name, p.Sprites.Front_default}

}
