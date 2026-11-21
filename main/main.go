package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/tijchy/vremenska/vreme"
)

func main() {
	app := &cli.Command{
		Name:  "vremenska",
		Usage: "Izpis vremenskih meritev in poročila",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
				Usage: "najnižja dovoljena temperatura",
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
				Usage: "najvišja dovoljena temperatura",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "najmanjše število meritev za statistiko",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float64("minTemp")
			maxTemp := cmd.Float64("maxTemp")
			stMeritev := cmd.Int("stMeritev")
			postaje := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(postaje, "LJ01", "Ljubljana Tivoli", "Ljubljana")
			vreme.DodajPostajo(postaje, "TR01", "Tržič AP", "Tržič")
			vreme.DodajPostajo(postaje, "MB01", "Maribor center", "Maribor")
			vreme.DodajPostajo(postaje, "KR01", "Kranj Planina", "Kranj")

			vreme.DodajMeritev(postaje, "LJ01", 12.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 13.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 14.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "TR01", 30.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "TR01", 30.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "TR01", 28.3, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", 15.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", 16.2, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "MB01", 17.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KR01", 11.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KR01", 12.5, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "KR01", 13.3, minTemp, maxTemp)

			vreme.DodajMeritev(postaje, "LJ01", 15.0, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "TR01", -200, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ02", 18.5, minTemp, maxTemp)

			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, stMeritev)
			return nil
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}
