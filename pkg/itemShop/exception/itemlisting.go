package exception 

type  ItemListing struct {}


func ( i *ItemListing) Error() string{	
	return "itemlisting error.."
}