package main
/*
	@author	boosh
	@desc	Implementation of "Pattern Breeder" image generator

*/

import (
	"fmt"
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

	size := 20
	var p [20][20]cell

	//prints array out, hardcoded for 20 columns
	for a := 0;a < size;a++ {
		fmt.Println(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d",p[a][0].alive,p[a][1].alive,p[a][2].alive,p[a][3].alive,p[a][4].alive,p[a][5].alive,p[a][6].alive,p[a][7].alive,p[a][8].alive,p[a][9].alive,p[a][10].alive,p[a][11].alive,p[a][12].alive,p[a][13].alive,p[a][14].alive,p[a][15].alive,p[a][16].alive,p[a][17].alive,p[a][18].alive,p[a][19].alive))
	}

	fmt.Println("")
	fmt.Println("---------------------------------------")
	fmt.Println("")

	//makes initial pattern
	//initial pattern is a single cell at the centerish
	p[size/2][size/2].alive = 1
	


	//5 iterations for testing
	for z:=0;z<5;z++{

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
			} else {
				p[ii][jj].alive=0
			}
			
			p[ii][jj].toenable = false

		}
	}

	//prints array out, hardcoded for 20 columns
	for a := 0;a < size;a++ {
		fmt.Println(fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d",p[a][0].alive,p[a][1].alive,p[a][2].alive,p[a][3].alive,p[a][4].alive,p[a][5].alive,p[a][6].alive,p[a][7].alive,p[a][8].alive,p[a][9].alive,p[a][10].alive,p[a][11].alive,p[a][12].alive,p[a][13].alive,p[a][14].alive,p[a][15].alive,p[a][16].alive,p[a][17].alive,p[a][18].alive,p[a][19].alive))
	}

	fmt.Println("")
	fmt.Println("---------------------------------------")
	fmt.Println("")

	}//end of 5 iterations for testing

}
