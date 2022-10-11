package image

import (
	"fmt"
	"image/color"
	"io"
	"net/http"
	"os"

	svg "github.com/ajstarks/svgo"
)

const (
	sqWidth     = 45
	sqHeight    = 45
	ranks       = 8
	files       = 8
	boardWidth  = ranks * sqWidth
	boardHeight = files * sqHeight
)

var (
	orderOfRanks = []Rank{Rank8, Rank7, Rank6, Rank5, Rank4, Rank3, Rank2, Rank1}
	orderOfFiles = []File{File1, File2, File3, File4, File5, File6, File7, File8}
)

// TODO: wrap w param into a struct
func generateBoardSVG(b *Board, w http.ResponseWriter) {
	//pieceMap := b.generateMapping()

	canvas := svg.New(w)
	canvas.Start(boardWidth, boardHeight)
	canvas.Rect(0, 0, boardWidth, boardHeight)

	ranks := orderOfRanks
	files := orderOfFiles

	for i, rank := range ranks {
		for j, file := range files {
			x := j * sqHeight
			y := i * sqWidth
			sq := newSquare(rank, file)
			c := sq.Color()
			piece := b.getPiece(sq)
			canvas.Rect(x, y, sqWidth, sqHeight, "fill: "+colorToHex(c))
			if piece != NoPiece {
				file, err := os.ReadFile("./resources/BB.svg")
				if err != nil {
					panic(err)
				}
				println(file)
				io.WriteString(canvas.Writer, string(file))
			}
		}
	}

	canvas.End()
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(float64(r)+0.5), uint8(float64(g)*1.0+0.5), uint8(float64(b)*1.0+0.5))
}
