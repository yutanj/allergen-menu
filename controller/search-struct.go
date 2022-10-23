package controller

type Sushiro struct{

}

//アレルゲン-特定原材料７品目
type SevenAllergenicIngredients struct {
	Shrimp int
	Crab int 
	Wheat int
	Buckwheat int
	Egg int 
	Milk int
	Peanuts int
}

//アレルゲン-特定原材料以外の２１品目
type TwentyOneAllergenicIngredients struct{
	Almond int 
	Abolone int 
	Squid int 
	SalmonRoe int
	Orange int 
	Cashewnut int
	KiwiFruit int
	Beef int 
	Walnut int
	Sesame int 
	Mackrel int
	Soybean int 
	Chicken int
	Banana int
	Pork int
	Matsutake int 
	Peach int
	Wildyam int
	Apple int
	Gelatin int
}