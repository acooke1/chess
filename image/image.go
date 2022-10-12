package image

import (
	"fmt"
	"image/color"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/acooke1/chess"
	"github.com/acooke1/chess/image/internal"
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
	orderOfRanks = []chess.Rank{chess.Rank8, chess.Rank7, chess.Rank6, chess.Rank5, chess.Rank4, chess.Rank3, chess.Rank2, chess.Rank1}
	orderOfFiles = []chess.File{chess.File1, chess.File2, chess.File3, chess.File4, chess.File5, chess.File6, chess.File7, chess.File8}
)

// TODO: wrap w param into a struct
func GenerateBoardSVG(b *chess.Board, w http.ResponseWriter) {
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
			sq := chess.NewSquare(rank, file)
			c := sq.Color()
			piece := b.GetPiece(sq)
			canvas.Rect(x, y, sqWidth, sqHeight, "fill: "+colorToHex(c))
			if piece != chess.NoPiece {
				file, err := os.ReadFile("./resources/BB.svg")
				if err != nil {
					panic(err)
				}
				println(file)
				io.WriteString(canvas.Writer, pieceXML(x, y, piece))
			}
		}
	}

	canvas.End()
}

func pieceXML(x, y int, p chess.Piece) string {
	fileName := fmt.Sprintf("pieces/%s%s.svg", p.Color().String(), pieceTypeMap[p.Type()])
	svgStr := string(internal.MustAsset(fileName))
	old := `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="45" height="45">`
	new := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" version="1.1" width="360" height="360" viewBox="%d %d 360 360">`, (-1 * x), (-1 * y))
	return strings.Replace(svgStr, old, new, 1)
}

func colorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", uint8(float64(r)+0.5), uint8(float64(g)*1.0+0.5), uint8(float64(b)*1.0+0.5))
}

var (
	pieceTypeMap = map[chess.PieceType]string{
		chess.King:   "K",
		chess.Queen:  "Q",
		chess.Rook:   "R",
		chess.Bishop: "B",
		chess.Knight: "N",
		chess.Pawn:   "P",
	}
)
