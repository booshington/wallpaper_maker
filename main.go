package main
/*
	@author	boosh
	@desc	Implementation of "Pattern Breeder" image generator

*/

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

/*
	cell type for pattern breeder
	cells can be alive (int > 0) or dead (int = 0)
*/
type cell struct {
	alive		int
	toenable	bool
}

func main(){
	fmt.Println("hello")

	//makes the cell array and the output image
	size := 1000
	var p [1000][1000]cell
	imgOut := image.NewNRGBA(image.Rect(0,0,1000,1000))

	//initial pattern is a single cell at the centerish
//	p[size/2][size/2].alive = 1

	//initial pattern is all four corners are alive
	p[0][0].alive=1
	p[size-1][size-1].alive=1
	p[0][size-1].alive=1
	p[size-1][0].alive=1


	//1000 iterations, each one updates the entire image of alive/dead
	//more iterations = more complex and interesting image
	for z:=0;z<1000;z++{

	//update all cells in array
	for i:=0;i<size;i++{
		for j:=0;j<size;j++{
			//If the cell has an adjacent alive cell, increment count
			//TODO: simplify logic
			count := 0
			//check cell j-1
			if ((j-1) >= 0){
				if p[i][j-1].alive > 0{
					count++
				}
			}
			//check cell i-1
			if ((i-1) >= 0){
				if p[i-1][j].alive > 0{
					count++
				}
			}
			//check cell j+1
			if ((j+1) < size){
				if p[i][j+1].alive > 0{
					count++
				}
			}
			//check cell i+1
			if ((i+1) < size){
				if p[i+1][j].alive > 0{
					count++
				}
			}
			//if count is even, set cell to dead, else count is odd so set cell to flag as alive
			//using toenable allows for an entire cell array to get analyzed at once, then all alive/dead flags get set simultaneously
			//without this, you can get a cascading alive/dead sequence from a single cell
			if (count % 2) == 0{
				p[i][j].toenable = false

			} else {
				p[i][j].toenable = true
			}
		}
	}

	//all cells that have toenable set get set to alive, all other cells get set to dead
	//and toenable gets reset to false
	for ii:=0;ii<size;ii++{
		for jj:=0;jj<size;jj++{
			if p[ii][jj].toenable == true {
				p[ii][jj].alive=1
				imgOut.Set(ii,jj, color.RGBA{255,255,255,255})
			} else {
				p[ii][jj].alive=0
				imgOut.Set(ii,jj, color.RGBA{0,0,0,255})
			}

			p[ii][jj].toenable = false

		}
	}

	}//end of iterations for image updating

	fmt.Println("generated png image, saving...")
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f,imgOut)
}
