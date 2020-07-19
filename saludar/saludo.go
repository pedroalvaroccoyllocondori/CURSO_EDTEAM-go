package saludar

// no se puede utilizar el acoratdor de asignacion ":="

var emoji = "🙋‍♂"

// se tienen u poner en mayusculas las primeras letras
// para que sepa si hay que exportarlas o no
func Saludoingles() string {
	return "hi" + emoji
}
func Saludoitaliano() string {
	return "ciao" + emoji

}
